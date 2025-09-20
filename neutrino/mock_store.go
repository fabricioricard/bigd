package neutrino

import (
	"github.com/pkt-cash/pktd/blockchain"
	"github.com/pkt-cash/pktd/btcutil/er"

	"github.com/pkt-cash/pktd/chaincfg/chainhash"
	"github.com/pkt-cash/pktd/neutrino/headerfs"
	"github.com/pkt-cash/pktd/pktwallet/waddrmgr"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
	"github.com/pkt-cash/pktd/wire"
)

type BlockHeader struct {
	headers map[chainhash.Hash]wire.BlockHeader
	heights map[uint32]wire.BlockHeader
}

type FilterHeader struct {
	HeaderHash chainhash.Hash
	FilterHash chainhash.Hash
	Height     uint32
}

type mockneutrinoDBStore struct {
	Db				  walletdb.DB
	blockHeaderIndex  *BlockHeader
	filterHeaderIndex *FilterHeader
}

// NewMockBlockHeaderStore returns a version of the BlockHeaderStore that's
// backed by an in-memory map. This instance is meant to be used by callers
// outside the package to unit test components that require a BlockHeaderStore
// interface.
func newMockNeutrinoDBStore() *mockneutrinoDBStore {
	return &mockneutrinoDBStore{
		blockHeaderIndex: &BlockHeader{
			headers: make(map[chainhash.Hash]wire.BlockHeader),
			heights: make(map[uint32]wire.BlockHeader),
		},
	}
}

func (m *mockneutrinoDBStore) BlockChainTip() (*wire.BlockHeader,
	uint32, er.R) {
	return nil, 0, nil
}
func (m *mockneutrinoDBStore) BlockChainTip1(tx walletdb.ReadTx) (*wire.BlockHeader,
	uint32, er.R) {
	return nil, 0, nil
}
func (m *mockneutrinoDBStore) FilterChainTip() (*chainhash.Hash, uint32, er.R) {
	return nil, 0, nil
}
func (m *mockneutrinoDBStore) FilterChainTip1(tx walletdb.ReadTx) (*chainhash.Hash, uint32, er.R) {
	return nil, 0, nil
}

func (m *mockneutrinoDBStore) FetchBlockHeaderByHeight(height uint32) (
	*wire.BlockHeader, er.R) {

	if header, ok := m.blockHeaderIndex.heights[height]; ok {
		return &header, nil
	}

	return nil, headerfs.ErrHeightNotFound.Default()
}

func (m *mockneutrinoDBStore) FetchBlockHeaderByHeight1(tx walletdb.ReadTx, height uint32) (
	*wire.BlockHeader, er.R) {
	return m.FetchBlockHeaderByHeight(height)
}

func (m *mockneutrinoDBStore) FetchFilterHeaderByHeight(height uint32) (*chainhash.Hash, er.R) {
	return nil, nil
}

func (m *mockneutrinoDBStore) FetchFilterHeader(hash *chainhash.Hash) (*chainhash.Hash, er.R) {
	return nil, nil
}

func (m *mockneutrinoDBStore) FetchFilterHeader1(tx walletdb.ReadTx, hash *chainhash.Hash) (*chainhash.Hash, er.R) {
	return nil, nil
}

func (m *mockneutrinoDBStore) FetchBlockHeaderAncestors(uint32,
	*chainhash.Hash) ([]wire.BlockHeader, uint32, er.R) {

	return nil, 0, nil
}

func (m *mockneutrinoDBStore) FetchFilterHeaderAncestors(numHeaders uint32, stopHash *chainhash.Hash) ([]chainhash.Hash, uint32, er.R) {
	return nil, 0, nil
}

func (m *mockneutrinoDBStore) HeightFromHash(*chainhash.Hash) (uint32, er.R) {
	return 0, nil

}
func (m *mockneutrinoDBStore) RollbackLastHeaderBlock(tx walletdb.ReadWriteTx) (*waddrmgr.BlockStamp,
	er.R) {
	return nil, nil
}

func (m *mockneutrinoDBStore) RollbackLastFilterBlock(tx walletdb.ReadWriteTx) (*waddrmgr.BlockStamp, er.R) {
	return nil, nil
}

func (m *mockneutrinoDBStore) FetchBlockHeader(h *chainhash.Hash) (
	*wire.BlockHeader, uint32, er.R) {
	if header, ok := m.blockHeaderIndex.headers[*h]; ok {
		return &header, 0, nil
	}
	return nil, 0, er.Errorf("not found")
}
func (m *mockneutrinoDBStore) FetchBlockHeader1(tx walletdb.ReadTx, h *chainhash.Hash) (
	*wire.BlockHeader, uint32, er.R) {
	if header, ok := m.blockHeaderIndex.headers[*h]; ok {
		return &header, 0, nil
	}
	return nil, 0, er.Errorf("not found")
}

func (m *mockneutrinoDBStore) WriteBlockHeaders(tx walletdb.ReadWriteTx, headers ...headerfs.BlockHeader) er.R {
	for _, h := range headers {
		m.blockHeaderIndex.headers[h.BlockHash()] = *h.BlockHeader
	}

	return nil
}

func (m *mockneutrinoDBStore) WriteFilterHeaders(tx walletdb.ReadWriteTx, hdrs ...headerfs.FilterHeader) er.R {
	return nil
}

func (m *mockneutrinoDBStore) GetWalletDB() *walletdb.DB {
	return nil
}

func (m *mockneutrinoDBStore) LatestBlockLocator() (blockchain.BlockLocator, er.R) {
	return nil, nil
}
