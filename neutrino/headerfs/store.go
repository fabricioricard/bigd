package headerfs

import (
	"bytes"
	"encoding/hex"
	"sync"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/pktlog/log"

	"github.com/pkt-cash/pktd/blockchain"
	"github.com/pkt-cash/pktd/chaincfg/chainhash"
	"github.com/pkt-cash/pktd/pktwallet/waddrmgr"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
	"github.com/pkt-cash/pktd/wire"
)

// headerBufPool is a pool of bytes.Buffer that will be re-used by the various
// headerStore implementations to batch their header writes to disk. By
// utilizing this variable we can minimize the total number of allocations when
// writing headers to disk.
var headerBufPool = sync.Pool{
	New: func() interface{} { return new(bytes.Buffer) },
}

// type NeutrinoDBStore_ struct {
// 	Db          walletdb.DB
// 	headerStore *NeutrinoDBStore
// 	//blockHeaderIndex  *headerIndex
// 	//filterHeaderIndex *headerIndex
// }

type RollbackHeader struct {
	BlockHeader  *waddrmgr.BlockStamp
	FilterHeader *chainhash.Hash
}

// FetchHeader attempts to retrieve a block header determined by the passed
// block height.
//
// NOTE: Part of the BlockHeaderStore interface.
func (h *NeutrinoDBStore) FetchBlockHeader(hash *chainhash.Hash) (*wire.BlockHeader, uint32, er.R) {
	var header *wire.BlockHeader
	var height uint32
	return header, height, walletdb.View(h.Db, func(tx walletdb.ReadTx) er.R {
		var err er.R
		header, height, err = h.FetchBlockHeader1(tx, hash)
		return err
	})
}

// FetchHeaderByHeight attempts to retrieve a target block header based on a
// block height.
//
// NOTE: Part of the BlockHeaderStore interface.
func (h *NeutrinoDBStore) FetchBlockHeaderByHeight(height uint32) (*wire.BlockHeader, er.R) {
	var header *wire.BlockHeader
	return header, walletdb.View(h.Db, func(tx walletdb.ReadTx) er.R {
		var err er.R
		header, err = h.FetchBlockHeaderByHeight1(tx, height)
		return err
	})
}

// FetchHeaderAncestors fetches the numHeaders block headers that are the
// ancestors of the target stop hash. A total of numHeaders+1 headers will be
// returned, as we'll walk back numHeaders distance to collect each header,
// then return the final header specified by the stop hash. We'll also return
// the starting height of the header range as well so callers can compute the
// height of each header without knowing the height of the stop hash.
//
// NOTE: Part of the BlockHeaderStore interface.
func (h *NeutrinoDBStore) FetchBlockHeaderAncestors(
	numHeaders uint32,
	stopHash *chainhash.Hash,
) ([]wire.BlockHeader, uint32, er.R) {
	var headers []wire.BlockHeader
	var startHeight uint32
	return headers, startHeight, walletdb.View(h.Db, func(tx walletdb.ReadTx) er.R {
		// First, we'll find the final header in the range, this will be the
		// ending height of our scan.
		endEntry, err := h.headerEntryByHash(tx, stopHash)
		if err != nil {
			return err
		}

		startHeight = endEntry.Height - numHeaders
		if headers, err = h.readBlockHeaderRange(tx, startHeight, endEntry.Height); err != nil {
			return err
		} else if len(headers) == 0 {
			return er.Errorf("Fetching %v headers up to %v - no results",
				numHeaders, stopHash)
		} else if realHash := headers[len(headers)-1].BlockHash(); realHash != endEntry.Header.blockHeader.BlockHash() {
			return er.Errorf("Fetching %v headers up to %v - hash mismatch, got %v",
				numHeaders, stopHash, realHash)
		}
		return err
	})
}

// HeightFromHash returns the height of a particular block header given its
// hash.
//
// NOTE: Part of the BlockHeaderStore interface.
func (h *NeutrinoDBStore) HeightFromHash(hash *chainhash.Hash) (uint32, er.R) {
	var height uint32
	return height, walletdb.View(h.Db, func(tx walletdb.ReadTx) er.R {
		if he, err := h.headerEntryByHash(tx, hash); err != nil {
			return err
		} else {
			height = he.Height
			return nil
		}
	})
}

func (h *NeutrinoDBStore) RollbackLastBlock(tx walletdb.ReadWriteTx) (*RollbackHeader, er.R) {
	result := RollbackHeader{}
	prev, err := h.truncateBlockIndex(tx)
	if err != nil {
		result.BlockHeader = nil
		result.FilterHeader = nil
		return &result, err
	} else {
		result.BlockHeader = &waddrmgr.BlockStamp{}
		result.FilterHeader = &chainhash.Hash{}

		result.BlockHeader.Hash = prev.Header.blockHeader.BlockHash()
		result.BlockHeader.Height = int32(prev.Height)
		result.FilterHeader = prev.Header.filterHeader
	}
	return &result, nil
}

// BlockHeader is a Bitcoin block header that also has its height included.
type BlockHeader struct {
	*wire.BlockHeader
	// Height is the height of this block header within the current main
	// chain.
	Height uint32
}

// toIndexEntry converts the BlockHeader into a matching headerEntry. This
// method is used when a header is to be written to disk.
func (b *BlockHeader) toIndexEntry() *headerEntry {
	var buf [80]byte
	hb := bytes.NewBuffer(buf[:])
	hb.Reset()

	// Finally, decode the raw bytes into a proper bitcoin header.
	if err := b.Serialize(hb); err != nil {
		panic(er.Errorf("Failed to serialize header %v", err))
	}
	return &headerEntry{
		blockHeader: *b.BlockHeader,
	}
}

func blockHeaderFromHe(he *headerEntryWithHeight) (*BlockHeader, er.R) {
	var ret wire.BlockHeader
	if err := ret.Deserialize(bytes.NewReader(he.Header.Bytes())); err != nil {
		return nil, err
	}
	return &BlockHeader{&ret, he.Height}, nil
}

// WriteHeaders writes a set of headers to disk.
//
// NOTE: Part of the BlockHeaderStore interface.
func (h *NeutrinoDBStore) WriteBlockHeaders(tx walletdb.ReadWriteTx, hdrs ...BlockHeader) er.R {
	headerLocs := make([]headerEntryWithHeight, len(hdrs))
	for i, header := range hdrs {
		headerLocs[i].Header = header.toIndexEntry()
		headerLocs[i].Height = header.Height
	}
	return h.addBlockHeaders(tx, headerLocs, false)
}

// blockLocatorFromHash takes a given block hash and then creates a block
// locator using it as the root of the locator. We'll start by taking a single
// step backwards, then keep doubling the distance until genesis after we get
// 10 locators.
//
// TODO(roasbeef): make into single transaction.
func (h *NeutrinoDBStore) blockLocatorFromHash(tx walletdb.ReadTx, he *headerEntry) (
	blockchain.BlockLocator, er.R) {

	var locator blockchain.BlockLocator
	hash := he.blockHeader.BlockHash()
	// Append the initial hash
	locator = append(locator, &hash)

	// If hash isn't found in DB or this is the genesis block, return the
	// locator as is
	hewh, err := h.headerEntryByHash(tx, &hash)
	if err != nil {
		//???
	}
	height := hewh.Height
	if height == 0 {
		return locator, nil
	}

	decrement := uint32(1)
	for height > 0 && len(locator) < wire.MaxBlockLocatorsPerMsg {
		// Decrement by 1 for the first 10 blocks, then double the jump
		// until we get to the genesis hash
		if len(locator) > 10 {
			decrement *= 2
		}

		if decrement > height {
			height = 0
		} else {
			height -= decrement
		}

		he, err := h.readHeader(tx, height)
		if err != nil {
			return locator, err
		}
		headerHash := he.Header.blockHeader.BlockHash()

		locator = append(locator, &headerHash)
	}

	return locator, nil
}

// LatestBlockLocator returns the latest block locator object based on the tip
// of the current main chain from the PoV of the database and flat files.
//
// NOTE: Part of the BlockHeaderStore interface.
func (h *NeutrinoDBStore) LatestBlockLocator() (blockchain.BlockLocator, er.R) {
	var locator blockchain.BlockLocator
	return locator, walletdb.View(h.Db, func(tx walletdb.ReadTx) er.R {
		if ct, err := h.chainTip(tx, bucketNameBlockTip); err != nil {
			return err
		} else {
			he := headerEntry{
				blockHeader:  ct.Header.blockHeader,
				filterHeader: ct.Header.filterHeader,
			}
			locator, err = h.blockLocatorFromHash(tx, &he)
			return err
		}
	})
}

// maybeResetHeaderState will reset the header state if the header assertion
// fails, but only if the target height is found. The boolean returned indicates
// that header state was reset.
func (f *NeutrinoDBStore) maybeResetHeaderState(
	tx walletdb.ReadWriteTx,
	headerStateAssertion *FilterHeader,
	nhs *NeutrinoDBStore,
) (bool, er.R) {

	failed := false

	if headerStateAssertion != nil {
		// First, we'll attempt to locate the header at this height. If no such
		// header is found, then we'll exit early.
		assertedHeader, err := f.FetchFilterHeaderByHeight(headerStateAssertion.Height)
		if assertedHeader == nil {
			if !ErrHeaderNotFound.Is(err) {
				return false, err
			}
		} else if *assertedHeader != headerStateAssertion.FilterHash {
			log.Warnf("Filter header at height %v is not %v, assertion failed, resyncing filters",
				headerStateAssertion.Height, headerStateAssertion.HeaderHash)
			failed = true
		}
	}

	if !failed && nhs != nil {
		hdr, err := f.chainTip(tx, bucketNameFilterTip)
		if err != nil {
			return false, err
		}
		for {
			hdrhash := hdr.Header.blockHeader.BlockHash()
			he := hdr.Header
			if bh, err := nhs.FetchBlockHeaderByHeight1(tx, hdr.Height); err != nil {
				if ErrHashNotFound.Is(err) {
					log.Warnf("We have filter header number %v but no block header, "+
						"resetting filter headers", hdr.Height)
					failed = true
					break
				}
				return false, err
			} else if bh := bh.BlockHash(); !hdrhash.IsEqual(&bh) {
				log.Warnf("Filter header / block header mismatch at height %v: %v != %v",
					hdr.Height, hdrhash, bh)
				failed = true
				break
			} else if len(he.Bytes()) != 32 {
				log.Warnf("Filter header at height %v is not 32 bytes: %v",
					hdr.Height, hex.EncodeToString(he.Bytes()))
				failed = true
				break
			} else if hdr.Height == 0 {
				break
			}
			height := hdr.Height - 1
			hdr, err = f.readHeader(tx, height)
			if err != nil {
				log.Warnf("Filter header missing at height %v (%v), resyncing filter headers",
					height, err)
				failed = true
				break
			}
		}
	}

	// If our on disk state and the provided header assertion don't match,
	// then we'll purge this state so we can sync it anew once we fully
	// start up.
	if failed {
		if err := f.deleteBuckets(tx); err != nil {
			return true, err
		} else {
			return true, f.createBuckets(tx)
		}
	}

	return false, nil
}

// FetchHeader returns the filter header that corresponds to the passed block
// height.
func (f *NeutrinoDBStore) FetchFilterHeader(hash *chainhash.Hash) (*chainhash.Hash, er.R) {
	var out *chainhash.Hash
	return out, walletdb.View(f.Db, func(tx walletdb.ReadTx) er.R {
		var err er.R
		out, err = f.FetchFilterHeader1(tx, hash)
		return err
	})
}
func (f *NeutrinoDBStore) FetchFilterHeader1(tx walletdb.ReadTx, hash *chainhash.Hash) (*chainhash.Hash, er.R) {
	if hdr, err := f.headerEntryByHash(tx, hash); err != nil {
		return nil, err
	} else if h, err := chainhash.NewHash(hdr.Header.filterHeader[:]); err != nil {
		return nil, err
	} else {
		return h, nil
	}
}

// FetchHeaderByHeight returns the filter header for a particular block height.
func (f *NeutrinoDBStore) FetchFilterHeaderByHeight(height uint32) (*chainhash.Hash, er.R) {
	var hash *chainhash.Hash
	return hash, walletdb.View(f.Db, func(tx walletdb.ReadTx) er.R {
		var h *chainhash.Hash
		if hdr, err := f.readHeader(tx, height); err != nil {
			return err
		} else if hdr.Header.filterHeader != nil {
			h, err = chainhash.NewHash(hdr.Header.filterHeader[:])
			if err != nil {
				return err
			}
		}
		hash = h
		return nil
	})
}

// FetchHeaderAncestors fetches the numHeaders filter headers that are the
// ancestors of the target stop block hash. A total of numHeaders+1 headers will be
// returned, as we'll walk back numHeaders distance to collect each header,
// then return the final header specified by the stop hash. We'll also return
// the starting height of the header range as well so callers can compute the
// height of each header without knowing the height of the stop hash.
func (f *NeutrinoDBStore) FetchFilterHeaderAncestors(
	numHeaders uint32,
	stopHash *chainhash.Hash,
) ([]chainhash.Hash, uint32, er.R) {
	var hashes []chainhash.Hash
	var height uint32
	return hashes, height, walletdb.View(f.Db, func(tx walletdb.ReadTx) er.R {
		// First, we'll find the final header in the range, this will be the
		// ending height of our scan.
		endEntry, err := f.headerEntryByHash(tx, stopHash)
		if err != nil {
			return err
		}
		startHeight := endEntry.Height - numHeaders
		hashes, err = f.readFilterHeaderRange(tx, startHeight, endEntry.Height)
		if err != nil {
			return err
		}
		// for i, h := range hashes {
		// 	log.Debugf("Load filter header %d => [%s]", startHeight+uint32(i), h)
		// }
		if !bytes.Equal(hashes[len(hashes)-1][:], endEntry.Header.filterHeader[:]) {
			return er.Errorf("Hash mismatch on %v: %v %v", endEntry.Height,
				hashes[len(hashes)-1], endEntry.Header.Bytes())
		}
		return nil
	})
}

// FilterHeader represents a filter header (basic or extended). The filter
// header itself is coupled with the block height and hash of the filter's
// block.
type FilterHeader struct {
	// HeaderHash is the hash of the block header that this filter header
	// corresponds to.
	HeaderHash chainhash.Hash

	// FilterHash is the filter header itself.
	FilterHash chainhash.Hash

	// Height is the block height of the filter header in the main chain.
	Height uint32
}

// WriteHeaders writes a batch of filter headers to persistent storage. The
// headers themselves are appended to the flat file, and then the index updated
// to reflect the new entires.
func (f *NeutrinoDBStore) WriteFilterHeaders(tx walletdb.ReadWriteTx, hdrs ...FilterHeader) er.R {
	return f.addFilterHeaders(tx, hdrs, false)
}

/////
/////

// ChainTip returns the best known block header and height for the
// blockHeaderStore.
//
// NOTE: Part of the BlockHeaderStore interface.
func (h *NeutrinoDBStore) BlockChainTip() (*wire.BlockHeader, uint32, er.R) {
	var bh *wire.BlockHeader
	var height uint32
	return bh, height, walletdb.View(h.Db, func(tx walletdb.ReadTx) er.R {
		var err er.R
		bh, height, err = h.BlockChainTip1(tx)
		return err
	})
}

// ChainTip returns the latest filter header and height known to the
// FilterHeaderStore.
func (f *NeutrinoDBStore) FilterChainTip() (*chainhash.Hash, uint32, er.R) {
	var hash *chainhash.Hash
	var height uint32
	return hash, height, walletdb.View(f.Db, func(tx walletdb.ReadTx) er.R {
		var err er.R
		hash, height, err = f.FilterChainTip1(tx)
		return err
	})
}
