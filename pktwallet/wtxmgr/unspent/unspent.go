package unspent

import (
	"encoding/hex"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/pktlog/log"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
	"github.com/pkt-cash/pktd/pktwallet/wtxmgr/dbstructs"
	"github.com/pkt-cash/pktd/pktwallet/wtxmgr/utilfun"
	"github.com/pkt-cash/pktd/wire"
)

var bucketUnspentOld = []byte("u")
var bucketUnspent = []byte("u2")

var UnspentErr = er.GenericErrorType.Code("UnspentErr")

// The unspent index records all outpoints for mined credits which are not spent
// by any other mined transaction records (but may be spent by a mempool
// transaction).
//
// Keys are use the canonical outpoint serialization:
//
//   [0:32]  Transaction hash (32 bytes)
//   [32:36] Output index (4 bytes)
//
// Values WERE serialized as such, now are json.
//
//   [0:4]   Block height (4 bytes)
//   [4:36]  Block hash (32 bytes)

func Put(ns walletdb.ReadWriteBucket, u *dbstructs.Unspent) er.R {
	k := utilfun.CanonicalOutPoint(&u.OutPoint.Hash, u.OutPoint.Index)
	if v, err := jsoniter.Marshal(u); err != nil {
		return UnspentErr.New("cannot marshal unspent", er.E(err))
	} else if err := ns.NestedReadWriteBucket(bucketUnspent).Put(k, v); err != nil {
		return UnspentErr.New("cannot put unspent", err)
	}
	return nil
}

func Get(ns walletdb.ReadBucket, outPoint *wire.OutPoint) (*dbstructs.Unspent, er.R) {
	k := utilfun.CanonicalOutPoint(&outPoint.Hash, outPoint.Index)
	bytes := ns.NestedReadBucket(bucketUnspent).Get(k)
	if len(bytes) == 0 {
		return nil, nil
	}
	var uns dbstructs.Unspent
	if err := decode(bytes, &uns); err != nil {
		return nil, UnspentErr.New(fmt.Sprintf("Unable to parse unspent with key [%s]",
			hex.EncodeToString(k)), err)
	}
	return &uns, nil
}

func decode(v []byte, uns *dbstructs.Unspent) er.R {
	if err := jsoniter.Unmarshal(v, uns); err != nil {
		return UnspentErr.New(fmt.Sprintf("Unable to parse JSON [%s]", string(v)), er.E(err))
	}
	return nil
}

func Delete(ns walletdb.ReadWriteBucket, outPoint *wire.OutPoint) er.R {
	k := utilfun.CanonicalOutPoint(&outPoint.Hash, outPoint.Index)
	err := ns.NestedReadWriteBucket(bucketUnspent).Delete(k)
	if err != nil {
		return UnspentErr.New("failed to delete unspent", err)
	}
	return nil
}

func ForEachUnspentOutput(
	ns walletdb.ReadBucket,
	beginKey []byte,
	visitor func(key []byte, unspent *dbstructs.Unspent) er.R,
) er.R {
	bu := ns.NestedReadBucket(bucketUnspent)
	return bu.ForEachBeginningWith(beginKey, func(k, v []byte) er.R {
		var unspent dbstructs.Unspent
		if err := decode(v, &unspent); err != nil {
			return err
		} else if err := visitor(k, &unspent); err != nil {
			return err
		}
		return nil
	})
}

func CreateBuckets(ns walletdb.ReadWriteBucket) er.R {
	_, err := ns.CreateBucket(bucketUnspent)
	return err
}
func DeleteBuckets(ns walletdb.ReadWriteBucket) er.R {
	return ns.DeleteNestedBucket(bucketUnspent)
}

func ExtendUnspents(ns walletdb.ReadWriteBucket, extend func(u *dbstructs.Unspent) er.R) er.R {
	if _, err := ns.CreateBucket(bucketUnspent); err != nil {
		return err
	}
	bu := ns.NestedReadBucket(bucketUnspentOld)
	if bucketUnspentOld == nil {
		log.Warn("There is no bucketUnspentOld bucket")
	}
	i := 0
	if err := bu.ForEach(func(k, v []byte) er.R {
		var unspent dbstructs.Unspent
		if err := utilfun.ReadCanonicalOutPoint(k, &unspent.OutPoint); err != nil {
			return err
		}
		if err := utilfun.ReadUnspentBlock(v, &unspent.Block); err != nil {
			return err
		}
		if err := extend(&unspent); err != nil {
			return err
		}
		if err := Put(ns, &unspent); err != nil {
			return err
		}
		if i%1000 == 0 {
			log.Infof("Migrating UTXO [%d] ([%d]%%)", i, int(k[0])*100/256)
		}
		i += 1
		return nil
	}); err != nil {
		return err
	}
	log.Info("Deleting old UTXO bucket")
	return ns.DeleteNestedBucket(bucketUnspentOld)
}
