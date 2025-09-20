package utilfun

import (
	"encoding/binary"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/chaincfg/chainhash"
	"github.com/pkt-cash/pktd/pktwallet/wtxmgr/dbstructs"
	"github.com/pkt-cash/pktd/wire"
)

// Several data structures are given canonical serialization formats as either
// keys or values.  These common formats allow keys and values to be reused
// across different buckets.
//
// The canonical outpoint serialization format is:
//
//   [0:32]  Trasaction hash (32 bytes)
//   [32:36] Output index (4 bytes)
//
// The canonical transaction hash serialization is simply the hash.

func CanonicalOutPoint(txHash *chainhash.Hash, index uint32) []byte {
	k := make([]byte, 36)
	copy(k, txHash[:])
	binary.BigEndian.PutUint32(k[32:36], index)
	return k
}

func ReadCanonicalOutPoint(k []byte, op *wire.OutPoint) er.R {
	if len(k) < 36 {
		return er.New("short canonical outpoint")
	}
	copy(op.Hash[:], k)
	op.Index = binary.BigEndian.Uint32(k[32:36])
	return nil
}

func ReadUnspentBlock(v []byte, block *dbstructs.Block) er.R {
	if len(v) < 36 {
		return er.New("short unspent value")
	}
	block.Height = int32(binary.BigEndian.Uint32(v))
	copy(block.Hash[:], v[4:36])
	return nil
}

func CreditKey(txHash *chainhash.Hash, index uint32, block *dbstructs.Block) []byte {
	k := make([]byte, 72)
	copy(k, txHash[:])
	binary.BigEndian.PutUint32(k[32:36], uint32(block.Height))
	copy(k[36:68], block.Hash[:])
	binary.BigEndian.PutUint32(k[68:72], index)
	return k
}

func CreditKeyForUnspent(uns *dbstructs.Unspent) []byte {
	return CreditKey(&uns.OutPoint.Hash, uns.OutPoint.Index, &uns.Block)
}
