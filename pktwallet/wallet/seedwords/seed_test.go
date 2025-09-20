// Copyright (c) 2020 Anode LLC
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package seedwords_test

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/pkt-cash/pktd/btcutil/util"
	"github.com/pkt-cash/pktd/pktwallet/wallet/seedwords"
	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	seed, err := seedwords.RandomSeed()
	if err != nil {
		t.Error(err)
		return
	}
	t0 := time.Now()
	se := seed.Encrypt([]byte("password"))
	t1 := time.Now()
	fmt.Printf("Time spent encrypting: %s\n", t1.Sub(t0))
	if seed1, err := se.Decrypt([]byte("password"), false); err != nil {
		t.Error(err)
		return
	} else if !bytes.Equal(seed1.Bytes(), seed.Bytes()) {
		t.Error("Seed decrypt is not the same")
	}
}

//	Test to check if a result seed get by deciphering a seed keeps the original attributes:w
func TestReversibleEncrypt(t *testing.T) {

	t.Run("decipher a ciphered seed using the mnemonic", func(t *testing.T) {

		//	create a new seed
		seed, err := seedwords.RandomSeed()
		util.RequireNoErr(t, err)

		//	cipher the seed
		cipherPass := []byte("kek")
		cipheredSeed := seed.Encrypt(cipherPass)

		//	get the mnemonic for the cyphered seed
		mnemonic, err := cipheredSeed.Words("english")
		util.RequireNoErr(t, err)

		//	try to reverse
		seedEnc, err := seedwords.SeedFromWords(mnemonic)
		util.RequireNoErr(t, err)

		decipheredSeed, err := seedEnc.Decrypt(cipherPass, false)

		//	load the configuration file
		require.Equal(t, seed.Version(), decipheredSeed.Version())
		require.Equal(t, seed.Birthday(), decipheredSeed.Birthday())
	})
}
