package dbstructs

import (
	"github.com/pkt-cash/pktd/chaincfg/chainhash"
	"github.com/pkt-cash/pktd/wire"
)

// dbstructs.Block contains the minimum amount of data to uniquely identify any block on
// either the best or side chain.
type Block struct {
	Hash   chainhash.Hash `json:"h"`
	Height int32          `json:"hei"`
}

type Unspent struct {
	OutPoint     wire.OutPoint `json:"op"`
	Block        Block         `json:"blk"`
	Address      string        `json:"addr"`
	Value        int64         `json:"val"`
	FromCoinBase bool          `json:"cb"`
	PkScript     []byte        `json:"scr"`
}
