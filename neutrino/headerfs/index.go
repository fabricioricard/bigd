package headerfs

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"sort"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/btcutil/gcs/builder"
	"github.com/pkt-cash/pktd/chaincfg"
	"github.com/pkt-cash/pktd/pktlog/log"
	"github.com/pkt-cash/pktd/wire"

	"github.com/pkt-cash/pktd/chaincfg/chainhash"
	"github.com/pkt-cash/pktd/chaincfg/genesis"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
)

var (
	//header_by_height contains both the header and filter
	//heights_by_hash_pfx contains the start of the hash and a list of heights
	// oldestBucket is very old code which used FS + db
	oldestBucket = []byte("header-index")
	// oldRootBucket was not as efficient storing data
	oldRootBucket = []byte("headers")

	// Current version of root bucket, inside of this is multiple buckets for named databases
	// We don't have a use for more than one, but if we were tracking multiple blockchains we might
	//
	// Under the dbBucketName bucket(s) there are all of the below buckets...
	newRootBucket = []byte("headers2")

	// This is a serialized headerEntry keyd by big endian height, it may by 80 bytes or 112 bytes
	// depending on whether the entry contains also a filter header
	bucketNameHeaderByHeight = []byte("header_by_height")

	// These are arrays of uint32 stored as packed byte arrays, they are keyed by 4 first bytes of the
	// block hash, there are multiple heights per block hash prefix because collisions are expected.
	bucketNameHeightsByHashPfx = []byte("heights_by_hash_pfx")

	// Height of the tip entry in the db
	bucketNameBlockTip = []byte("block_tip")

	// Height of the highest entry to have a filter header added
	bucketNameFilterTip = []byte("filter_tip")
)

var Err er.ErrorType = er.NewErrorType("headerfs.Err")

var (
	// ErrHeightNotFound is returned when a specified height isn't found in
	// a target index.
	ErrHeightNotFound = Err.CodeWithDetail("ErrHeightNotFound",
		"target height not found in index")

	// ErrHashNotFound is returned when a specified block hash isn't found
	// in a target index.
	ErrHashNotFound = Err.CodeWithDetail("ErrHashNotFound",
		"target hash not found in index")
)

// HeaderType is an enum-like type which defines the various header types that
// are stored within the index.
type HeaderType uint8

const (
	// Block is the header type that represents regular Bitcoin block
	// headers.
	Block HeaderType = iota

	// RegularFilter is a header type that represents the basic filter
	// header type for the filter header chain.
	RegularFilter
)

const (
	// BlockHeaderSize is the size in bytes of the Block header type.
	BlockHeaderSize = 80

	// FilterHeaderSize is the size in bytes of the RegularFilter
	// header type.
	FilterHeaderSize = 32

	TotalSize = BlockHeaderSize + FilterHeaderSize
)

type NeutrinoDBStore struct {
	Db           walletdb.DB
	dbBucketName []byte
}

type headerEntry struct {
	blockHeader  wire.BlockHeader
	filterHeader *chainhash.Hash
}

type headerEntryWithHeight struct {
	Header *headerEntry
	Height uint32
}

type headerWithHeightBatch []headerEntryWithHeight

func (h headerWithHeightBatch) Len() int {
	return len(h)
}

func (h headerWithHeightBatch) Less(i, j int) bool {
	return h[i].Height-h[j].Height < 0
}

func (h headerWithHeightBatch) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

type filterHeaderBatch []FilterHeader

func (h filterHeaderBatch) Len() int {
	return len(h)
}

func (h filterHeaderBatch) Less(i, j int) bool {
	return h[i].Height-h[j].Height < 0
}

func (h filterHeaderBatch) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

type heightByHashPfx struct {
	hashPrefix uint32
	heights    []uint32
}

// NewNeutrinoDBStore creates a new instance of the NeutrinoDBStore based on
// a target file path, an open database instance, and finally a set of
// parameters for the target chain. These parameters are required as if this is
// the initial start up of the NeutrinoDBStore, then the initial genesis
// header will need to be inserted.
func NewNeutrinoDBStore(db walletdb.DB, netParams *chaincfg.Params, verify bool) (*NeutrinoDBStore, er.R) {
	retry := false
	var hi *NeutrinoDBStore
	if err := walletdb.Update(db, func(tx walletdb.ReadWriteTx) er.R {
		// Drop the old buckets if it happens to exist
		if err := tx.DeleteTopLevelBucket(oldestBucket); err != nil && !walletdb.ErrBucketNotFound.Is(err) {
			return err
		}
		if err := tx.DeleteTopLevelBucket(oldRootBucket); err != nil && !walletdb.ErrBucketNotFound.Is(err) {
			return err
		}
		hi = &NeutrinoDBStore{
			Db:           db,
			dbBucketName: []byte(netParams.Name),
		}
		if err := hi.createBuckets(tx); err != nil {
			return err
		}
		if he, err := hi.headerHeightsByHash(tx, netParams.GenesisHash); err != nil || he == nil {
			gen := genesis.Block(netParams.GenesisHash)
			gh := []headerEntryWithHeight{}
			he := headerEntry{
				gen.Header,
				nil,
			}
			gh = append(gh, headerEntryWithHeight{&he, 0})
			log.Debug("Inserting genesis block header")
			if err := hi.addBlockHeaders(tx, gh, true); err != nil {
				return err
			}

			if basicFilter, err := builder.BuildBasicFilter(gen, nil); err != nil {
				return err
			} else if genesisFilterHash, err := builder.MakeHeaderForFilter(
				basicFilter,
				gen.Header.PrevBlock,
			); err != nil {
				return err
			} else {
				fh := filterHeaderBatch{
					FilterHeader{
						HeaderHash: *netParams.GenesisHash,
						FilterHash: genesisFilterHash,
						Height:     0,
					},
				}
				log.Debug("Inserting genesis filter header")
				if err := hi.addFilterHeaders(tx, fh, true); err != nil {
					return err
				}
			}
		}

		if !verify {
		} else if err := hi.CheckConnectivity(tx); err != nil {
			log.Warnf("CheckConnectivity failed [%v] resyncing header chain", err)
			hi.deleteBuckets(tx)
			retry = true
		}

		return nil
	}); err != nil {
		return nil, err
	}
	if retry {
		return NewNeutrinoDBStore(db, netParams, verify)
	}
	return hi, nil
}

// CheckConnectivity cycles through all of the block headers on disk, from last
// to first, and makes sure they all connect to each other. Additionally, at
// each block header, we also ensure that the index entry for that height and
// hash also match up properly.
func (h *NeutrinoDBStore) CheckConnectivity(tx walletdb.ReadTx) er.R {
	log.Info("[1] CheckConnectivity()")

	if he, err := h.chainTip(tx, bucketNameBlockTip); err != nil {
		return err
	} else if fhe, err := h.chainTip(tx, bucketNameFilterTip); err != nil {
		return err
	} else {
		blockHash := he.Header.blockHeader.BlockHash()
		for {
			if he.Height <= fhe.Height {
				// If it should have a filter header, it does
				// We don't check the filter header is at all sane
				if he.Header.filterHeader == nil {
					return er.Errorf("missing filter header at height [%d] filter tip is [%d]",
						he.Height, fhe.Height)
				}
			} else if he.Header.filterHeader != nil {
				return er.Errorf("unexpected filter header at height [%d] filter tip is [%d]",
					he.Height, fhe.Height)
			}

			// Check that we can access the header by it's hash
			if heights, err := h.headerHeightsByHash(tx, &blockHash); err != nil {
				return err
			} else {
				ok := false
				for _, h := range heights {
					if h == he.Height {
						ok = true
					}
				}
				if !ok {
					return er.Errorf("unable to get height from hash for [%s]", blockHash.String())
				}
			}

			if he.Height == 0 {
				break
			}

			// Get the previous header
			if hen1, err := h.readHeader(tx, he.Height-1); err != nil {
				return err
			} else {
				n1Hash := hen1.Header.blockHeader.BlockHash()
				if he.Header.blockHeader.PrevBlock != n1Hash {
					return er.Errorf("hash mismatch at height [%d -> %d]: want [%s] got [%s]",
						he.Height, he.Height-1, he.Header.blockHeader.PrevBlock, n1Hash)
				}
				he = hen1
				blockHash = n1Hash
			}
		}
	}
	return nil
}

func (h *NeutrinoDBStore) createBuckets(tx walletdb.ReadWriteTx) er.R {
	if bkt, err := h.rwBucket(tx); err != nil {
		return err
	} else if _, err := bkt.CreateBucketIfNotExists(bucketNameHeaderByHeight); err != nil {
		return err
	} else if _, err := bkt.CreateBucketIfNotExists(bucketNameHeightsByHashPfx); err != nil {
		return err
	} else {
		return nil
	}
}

func rootRwBucket(tx walletdb.ReadWriteTx) (walletdb.ReadWriteBucket, er.R) {
	root := tx.ReadWriteBucket(newRootBucket)
	if root == nil {
		if r, err := tx.CreateTopLevelBucket(newRootBucket); err != nil {
			return nil, err
		} else {
			root = r
		}
	}
	return root, nil
}

func (h *NeutrinoDBStore) deleteBuckets(tx walletdb.ReadWriteTx) er.R {
	root, err := rootRwBucket(tx)
	if err != nil {
		return err
	}
	if err := root.DeleteNestedBucket(h.dbBucketName); err != nil && !walletdb.ErrBucketNotFound.Is(err) {
		return err
	}
	return nil
}

func (h *NeutrinoDBStore) rwBucket(tx walletdb.ReadWriteTx) (walletdb.ReadWriteBucket, er.R) {
	root, err := rootRwBucket(tx)
	if err != nil {
		return nil, err
	}
	sub := root.NestedReadWriteBucket(h.dbBucketName)
	if sub == nil {
		if s, err := root.CreateBucket(h.dbBucketName); err != nil {
			return nil, err
		} else {
			sub = s
		}
	}
	return sub, nil
}

func (h *NeutrinoDBStore) roBucket(tx walletdb.ReadTx) (walletdb.ReadBucket, er.R) {
	root := tx.ReadBucket(newRootBucket)
	if root == nil {
		return nil, walletdb.ErrBucketNotFound.Default()
	}
	sub := root.NestedReadBucket(h.dbBucketName)
	if sub == nil {
		return nil, walletdb.ErrBucketNotFound.Default()
	}
	return sub, nil
}

func (he *headerEntry) Bytes() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, TotalSize))
	he.blockHeader.BtcEncode(buf, 0, wire.BaseEncoding)
	if he.filterHeader != nil {
		buf.Write(he.filterHeader[:])
	}
	return buf.Bytes()
}

func decodeHeaderEntry(b []byte) (*headerEntry, er.R) {
	if len(b) != TotalSize && len(b) != BlockHeaderSize {
		return nil, er.Errorf("headerEntry has size [%d] can not decode header", len(b))
	}
	var bh wire.BlockHeader
	err := bh.Deserialize(bytes.NewReader(b[0:BlockHeaderSize]))
	if err != nil {
		return nil, err
	}
	var hash *chainhash.Hash
	if len(b) == TotalSize {
		h, err := chainhash.NewHash(b[BlockHeaderSize:TotalSize])
		if err != nil {
			return nil, err
		}
		hash = h
	}
	return &headerEntry{
		blockHeader:  bh,
		filterHeader: hash,
	}, nil
}

func heightBin(height uint32) []byte {
	heightBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(heightBytes[:], height)
	return heightBytes
}

func binHeight(height []byte) uint32 {
	return binary.BigEndian.Uint32(height[:])
}

// Return headerEntry
// NOTE: process is take the first 4 bytes of the hash, lookup in heightByHashPrefix
// Take each match, load the header, take the BlockHeader and do BlockHash(), compare hash, match => return

// Return a slice of possible heights for the header given the hash, we need to access each
// one in order to figure out which one is actually correct.
func (h *NeutrinoDBStore) headerHeightsByHash(tx walletdb.ReadTx, hash *chainhash.Hash) ([]uint32, er.R) {
	rootBucket, err := h.roBucket(tx)
	if err != nil {
		return nil, err
	}
	heightsByHashBucket := rootBucket.NestedReadBucket(bucketNameHeightsByHashPfx)
	heights := heightsByHashBucket.Get(hash[0:4])
	if heights == nil {
		// not found, return empty array (nil)
		return nil, nil
	}
	if len(heights)%4 != 0 {
		return nil, er.Errorf("DB corruption: heights length for entry [%s] is [%d] which is not a multiple of 4",
			hex.EncodeToString(hash[0:4]), len(heights))
	}
	out := make([]uint32, 0, 4)
	for i := 0; i < len(heights); i += 4 {
		out = append(out, binHeight(heights[i:i+4]))
	}
	return out, nil
}
func (h *NeutrinoDBStore) modifyHeightsByHashPfx(
	tx walletdb.ReadWriteTx,
	hash *chainhash.Hash,
	height uint32,
	delete bool,
) er.R {
	if heights, err := h.headerHeightsByHash(tx, hash); err != nil {
		return err
	} else {
		newHeights := make([]byte, 0, (len(heights)+1)*4)
		if delete {
			ok := false
			for _, h := range heights {
				if h != height {
					newHeights = append(newHeights, heightBin(h)...)
				} else {
					ok = true
				}
			}
			if !ok {
				return er.Errorf("Error deleting hash entry [%s @ %d] from prefix table, not found",
					hash, height)
			}
		} else {
			newHeights = append(newHeights, heightBin(height)...)
			for _, h := range heights {
				if h != height {
					newHeights = append(newHeights, heightBin(h)...)
				} else {
					return er.Errorf("Error adding hash entry [%s @ %d] to prefix table, exists",
						hash, height)
				}
			}
		}
		rootBucket, err := h.rwBucket(tx)
		if err != nil {
			return err
		}
		hbhp := rootBucket.NestedReadWriteBucket(bucketNameHeightsByHashPfx)
		return hbhp.Put(hash[0:4], newHeights)
	}
}

func (h *NeutrinoDBStore) headerEntryByHash(tx walletdb.ReadTx, hash *chainhash.Hash) (*headerEntryWithHeight, er.R) {
	if heights, err := h.headerHeightsByHash(tx, hash); err != nil {
		return nil, err
	} else {
		for _, height := range heights {
			if he, err := h.readHeader(tx, height); err != nil {
				return nil, err
			} else if he.Header.blockHeader.BlockHash() == *hash {
				return he, nil
			}
		}
	}
	return nil, ErrHashNotFound.New("", er.Errorf("With hash %v", hash))
}
func (h *NeutrinoDBStore) FetchBlockHeader1(tx walletdb.ReadTx, hash *chainhash.Hash) (
	*wire.BlockHeader, uint32, er.R,
) {
	if he, err := h.headerEntryByHash(tx, hash); err != nil {
		return nil, 0, err
	} else {
		return &he.Header.blockHeader, he.Height, nil
	}
}
func (h *NeutrinoDBStore) HeightFromHash1(tx walletdb.ReadTx, hash *chainhash.Hash) (uint32, er.R) {
	if he, err := h.headerEntryByHash(tx, hash); err != nil {
		return 0, err
	} else {
		return he.Height, nil
	}
}

func (h *NeutrinoDBStore) readHeader(tx walletdb.ReadTx, height uint32) (*headerEntryWithHeight, er.R) {
	rootBucket, err := h.roBucket(tx)
	if err != nil {
		return nil, err
	}
	byheight := rootBucket.NestedReadBucket(bucketNameHeaderByHeight)
	hb := heightBin(height)
	if entryBytes := byheight.Get(hb[:]); entryBytes == nil {
		// If the hash wasn't found, then we don't know of this
		// hash within the index.
		return nil, ErrHashNotFound.New("", er.Errorf("height: %v", height))
	} else if he, err := decodeHeaderEntry(entryBytes); err != nil {
		return nil, err
	} else {
		return &headerEntryWithHeight{
			Header: he,
			Height: height,
		}, nil
	}
}
func (h *NeutrinoDBStore) FetchBlockHeaderByHeight1(
	tx walletdb.ReadTx,
	height uint32,
) (*wire.BlockHeader, er.R) {
	if he, err := h.readHeader(tx, height); err != nil {
		return nil, err
	} else {
		return &he.Header.blockHeader, nil
	}
}

func (h *NeutrinoDBStore) chainTip(tx walletdb.ReadTx, chainTipType []byte) (*headerEntryWithHeight, er.R) {
	rootBucket, err := h.roBucket(tx)
	if err != nil {
		return nil, err
	}
	tipHeightBytes := rootBucket.Get(chainTipType)
	if tipHeightBytes == nil {
		return nil, er.Errorf("no chain tip found in %s", chainTipType)
	}
	if len(tipHeightBytes) < 4 {
		tipHeightBytes = []byte{0x00, 0x00, 0x00, 0x00}
	}
	tipHeight := binHeight(tipHeightBytes)
	he, err := h.readHeader(tx, tipHeight)
	if err != nil {
		return nil, err
	}
	return he, nil
}
func (h *NeutrinoDBStore) FilterChainTip1(tx walletdb.ReadTx) (*chainhash.Hash, uint32, er.R) {
	if he, err := h.chainTip(tx, bucketNameFilterTip); err != nil {
		return nil, 0, err
	} else if he.Header.filterHeader == nil {
		return nil, 0, ErrHashNotFound.New("", er.Errorf("height: %v", he.Height))
	} else {
		return he.Header.filterHeader, he.Height, nil
	}
}
func (h *NeutrinoDBStore) BlockChainTip1(tx walletdb.ReadTx) (*wire.BlockHeader, uint32, er.R) {
	if he, err := h.chainTip(tx, bucketNameBlockTip); err != nil {
		return nil, 0, err
	} else {
		return &he.Header.blockHeader, he.Height, nil
	}
}

func (h *NeutrinoDBStore) truncateBlockIndex(tx walletdb.ReadWriteTx) (*headerEntryWithHeight, er.R) {
	rootBucket, err := h.rwBucket(tx)
	if err != nil {
		return nil, err
	}

	//rolls back both block and filter
	if ct, err := h.chainTip(tx, bucketNameBlockTip); err != nil {
		return nil, err
	} else if err := rootBucket.Put(bucketNameBlockTip, heightBin(ct.Height-1)); err != nil {
		return nil, err
	} else {
		log.Warnf("ROLLBACK in neutrino removes [%s @ %d]", ct.Header.blockHeader.BlockHash(), ct.Height)
		// In case the filter tip is equal to the block tip, we must roll that one back as well
		filterTipBytes := rootBucket.Get(bucketNameFilterTip)
		if filterTipBytes != nil && binHeight(filterTipBytes) >= ct.Height {
			if err := rootBucket.Put(bucketNameFilterTip, heightBin(ct.Height-1)); err != nil {
				return nil, err
			}
		}

		blockHash := ct.Header.blockHeader.BlockHash()
		if err := h.modifyHeightsByHashPfx(tx, &blockHash, ct.Height, true); err != nil {
			return nil, err
		}
		hdrGroup := rootBucket.NestedReadWriteBucket(bucketNameHeaderByHeight)
		if err := hdrGroup.Delete(heightBin(ct.Height)); err != nil {
			return nil, err
		}
		return h.chainTip(tx, bucketNameBlockTip)
	}
}

// addHeaders writes a batch of header entries in a single atomic batch
func (h *NeutrinoDBStore) addBlockHeaders(tx walletdb.ReadWriteTx, batch headerWithHeightBatch, isGenesis bool) er.R {
	// If we're writing a 0-length batch, make no changes and return.
	if len(batch) == 0 {
		return nil
	}

	rootBucket, err := h.rwBucket(tx)
	if err != nil {
		return err
	}
	headerBucket := rootBucket.NestedReadWriteBucket(bucketNameHeaderByHeight)

	sort.Sort(batch)
	var tip *headerEntryWithHeight
	if !isGenesis {
		he, err := h.chainTip(tx, bucketNameBlockTip)
		if err != nil {
			return err
		}
		tip = he
	} else {
		tip = &headerEntryWithHeight{}
	}
	var heightBytes []byte
	for _, header := range batch {
		if !isGenesis && header.Height != tip.Height+1 {
			log.Warnf("Unable to add block header at height %v because tip is %v", header.Height, tip.Height)
			return er.Errorf("Unable to add block header at height %v because tip is %v", header.Height, tip.Height)
		}
		he := headerEntry{blockHeader: header.Header.blockHeader, filterHeader: nil}
		value := he.Bytes()
		heightBytes = heightBin(header.Height)
		if err := headerBucket.Put(heightBytes, value); err != nil {
			return err
		}
		blockHash := header.Header.blockHeader.BlockHash()
		if err := h.modifyHeightsByHashPfx(tx, &blockHash, header.Height, false); err != nil {
			return err
		}
		tip.Height = header.Height
		tip.Header = header.Header
	}
	return rootBucket.Put(bucketNameBlockTip, heightBytes)
}

func (h *NeutrinoDBStore) addFilterHeaders(tx walletdb.ReadWriteTx, batch filterHeaderBatch, isGenesis bool) er.R {
	// If we're writing a 0-length batch, make no changes and return.
	if len(batch) == 0 {
		return nil
	}

	rootBucket, err := h.rwBucket(tx)
	if err != nil {
		return err
	}
	headerBucket := rootBucket.NestedReadWriteBucket(bucketNameHeaderByHeight)

	sort.Sort(batch)
	var tip *headerEntryWithHeight
	if !isGenesis {
		he, err := h.chainTip(tx, bucketNameFilterTip)
		if err != nil {
			return err
		}
		tip = he
	} else {
		tip = &headerEntryWithHeight{}
	}
	var heightBytes []byte
	for _, header := range batch {
		if !isGenesis && header.Height != tip.Height+1 {
			log.Warnf("Unable to add filter header at height %v because tip is %v", header.Height, tip.Height)
			break
		}
		if he, err := h.readHeader(tx, header.Height); err != nil {
			return err
		} else {
			blockHash := he.Header.blockHeader.BlockHash()
			if blockHash != header.HeaderHash {
				return er.Errorf("Unable to add filter header, block header mismatch at height %v, want %s got %s",
					header.Height, blockHash.String(), header.HeaderHash.String())
			}
			he.Header.filterHeader = &header.FilterHash
			heightBytes = heightBin(header.Height)
			value := he.Header.Bytes()
			if err := headerBucket.Put(heightBytes, value); err != nil {
				return err
			}
			tip.Header = he.Header
			tip.Height = he.Height
		}
	}
	return rootBucket.Put(bucketNameFilterTip, heightBytes)
}
