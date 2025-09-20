package walletunlocker_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/btcutil/util"
	"github.com/pkt-cash/pktd/chaincfg"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
	"github.com/pkt-cash/pktd/lnd/lnwallet/btcwallet"
	"github.com/pkt-cash/pktd/lnd/walletunlocker"
	"github.com/pkt-cash/pktd/pktlog/log"
	"github.com/pkt-cash/pktd/pktwallet/snacl"
	"github.com/pkt-cash/pktd/pktwallet/waddrmgr"
	"github.com/pkt-cash/pktd/pktwallet/wallet"
	"github.com/pkt-cash/pktd/pktwallet/wallet/seedwords"
	"github.com/stretchr/testify/require"
)

const (
	testWalletFilename string = "wallet.db"
)

var (
	testPassword = []byte("test-password")
	testSeed     = []byte("test-seed-123456789")
	testMac      = []byte("fakemacaroon")

	testNetParams = &chaincfg.MainNetParams

	testRecoveryWindow uint32 = 150

	//	due to the cipher routines, low timeout is making some test cases to fail when running on Github Actions
	defaultTestTimeout = 10 * time.Second
)

func createTestWallet(t *testing.T, dir string, netParams *chaincfg.Params) {
	createTestWalletWithPw(t, []byte(wallet.InsecurePubPassphrase), testPassword, dir, netParams)
}

func createTestWalletWithPw(t *testing.T, pubPw, privPw []byte, dir string,
	netParams *chaincfg.Params) {

	// Instruct waddrmgr to use the cranked down scrypt parameters when
	// creating new wallet encryption keys.
	fastScrypt := waddrmgr.FastScryptOptions
	keyGen := func(passphrase *[]byte, config *waddrmgr.ScryptOptions) (
		*snacl.SecretKey, er.R) {

		return snacl.NewSecretKey(
			passphrase, fastScrypt.N, fastScrypt.R, fastScrypt.P,
		)
	}
	waddrmgr.SetSecretKeyGen(keyGen)

	// Create a new test wallet that uses fast scrypt as KDF.
	netDir := btcwallet.NetworkDir(dir, netParams)
	loader := wallet.NewLoader(netParams, netDir, testWalletFilename, true, 0)
	_, err := loader.CreateNewWallet(
		pubPw, privPw, []byte(hex.EncodeToString(testSeed)), time.Time{}, nil,
	)
	util.RequireNoErr(t, err)

	realWalletPathname := wallet.WalletDbPath(netDir, testWalletFilename)
	log.Debugf(">>> createTestWalletWithPw [1] wallet path: %s", realWalletPathname)
	walletFileExists := true
	_, errr := os.Stat(realWalletPathname)
	if err != nil {
		if os.IsNotExist(errr) {
			walletFileExists = false
		} else {
			require.NoError(t, errr)
		}
	}
	log.Debugf(">>> createTestWalletWithPw [2] after loader.CreateNewWallet() the wallet file exists: %t", walletFileExists)

	err = loader.UnloadWallet()
	util.RequireNoErr(t, err)
}

func createSeedAndMnemonic(t *testing.T, pass []byte) (*seedwords.Seed, string) {

	cipherSeed, err := seedwords.RandomSeed()
	util.RequireNoErr(t, err)

	encipheredSeed := cipherSeed.Encrypt(pass)

	// With the new seed created, we'll convert it into a mnemonic phrase
	// that we'll send over to initialize the wallet.
	mnemonic, err := encipheredSeed.Words("english")
	util.RequireNoErr(t, err)

	return cipherSeed, mnemonic
}

// TestGenSeedUserEntropy tests that the gen seed method generates a valid
// cipher seed mnemonic phrase and user provided source of entropy.
func TestGenSeed(t *testing.T) {
	t.Parallel()

	log.Debugf(">>>>> running TestGenSeed()")
	// First, we'll create a new test directory and unlocker service for
	// that directory.
	testDir, errr := ioutil.TempDir("", "testcreate")
	require.NoError(t, errr)
	defer func() {
		_ = os.RemoveAll(testDir)
	}()

	service := walletunlocker.New(testDir, testNetParams, true, nil, "", testWalletFilename)

	// Now that the service has been created, we'll ask it to generate a
	// new seed for us given a test passphrase.
	seedPass := []byte("kek")
	genSeedReq := &lnrpc.GenSeedRequest{
		SeedPassphraseBin: seedPass,
		SeedEntropy:       make([]byte, 0),
	}

	ctx := context.Background()
	seedResp, errr := service.GenSeed(ctx, genSeedReq)
	require.NoError(t, errr)

	// We should then be able to take the generated mnemonic, and properly
	// decipher both it.
	mnemonic := strings.Join(seedResp.Seed, " ")
	_, err := seedwords.SeedFromWords(mnemonic)
	util.RequireNoErr(t, err)
}

// TestGenSeedInvalidEntropy tests that the gen seed method generates a valid
// cipher seed mnemonic pass phrase even when the user doesn't provide its own
// source of entropy.
//	the following test makes no sense anymore, since the seedwords package doesn't support entropy
/*
func TestGenSeedGenerateEntropy(t *testing.T) {
	t.Parallel()

	log.Debugf(">>>>> running TestGenSeedGenerateEntropy()")
	// First, we'll create a new test directory and unlocker service for
	// that directory.
	testDir, errr := ioutil.TempDir("", "testcreate")
	require.NoError(t, errr)
	defer func() {
		_ = os.RemoveAll(testDir)
	}()
	service := walletunlocker.New(testDir, testNetParams, true, nil, testDir, testWalletFilename)

	// Now that the service has been created, we'll ask it to generate a
	// new seed for us given a test passphrase. Note that we don't actually
	aezeedPass := []byte("kek")
	genSeedReq := &lnrpc.GenSeedRequest{
		AezeedPassphrase: aezeedPass,
	}

	ctx := context.Background()
	seedResp, errr := service.GenSeed(ctx, genSeedReq)
	require.NoError(t, errr)

	// We should then be able to take the generated mnemonic, and properly
	// decipher both it.
	mnemonic := strings.Join(seedResp.CipherSeedMnemonic, " ")
	_, err := seedwords.SeedFromWords(mnemonic)
	util.RequireNoErr(t, err)
}
*/

// TestGenSeedInvalidEntropy tests that if a user attempt to create a seed with
// a non empty initial entropy, then the proper error is returned.
func TestGenSeedInvalidEntropy(t *testing.T) {
	t.Parallel()

	log.Debugf(">>>>> running TestGenSeedInvalidEntropy()")
	// First, we'll create a new test directory and unlocker service for
	// that directory.
	testDir, errr := ioutil.TempDir("", "testcreate")
	require.NoError(t, errr)
	defer func() {
		_ = os.RemoveAll(testDir)
	}()
	service := walletunlocker.New(testDir, testNetParams, true, nil, "", testWalletFilename)

	// Now that the service has been created, we'll ask it to generate a
	// new seed for us given a test passphrase. However, we'll be using an
	// invalid set of entropy that's 55 bytes, instead of 15 bytes.
	seedPass := []byte("kek")
	genSeedReq := &lnrpc.GenSeedRequest{
		SeedPassphraseBin: seedPass,
		SeedEntropy:       bytes.Repeat([]byte("a"), 55),
	}

	// We should get an error now since the entropy source was invalid.
	ctx := context.Background()
	_, errr = service.GenSeed(ctx, genSeedReq)
	require.Error(t, errr)
	require.Contains(t, errr.Error(), "custom seed input entropy is not supported")
}

// TestInitWallet tests that the user is able to properly initialize the wallet
// given an existing cipher seed passphrase.
func TestInitWallet(t *testing.T) {
	t.Parallel()

	log.Debugf(">>>>> running TestInitWallet()")
	// testDir is empty, meaning wallet was not created from before.
	testDir, errr := ioutil.TempDir("", "testcreate")
	require.NoError(t, errr)
	defer func() {
		_ = os.RemoveAll(testDir)
	}()

	// Create new UnlockerService.
	service := walletunlocker.New(testDir, testNetParams, true, nil, "", testWalletFilename)

	// Once we have the unlocker service created, we'll now instantiate a
	// new cipher seed and its mnemonic.
	pass := []byte("test")
	cipherSeed, mnemonic := createSeedAndMnemonic(t, pass)

	// Now that we have all the necessary items, we'll now issue the Init
	// command to the wallet. This should check the validity of the cipher
	// seed, then send over the initialization information over the init
	// channel.
	ctx := context.Background()
	req := &lnrpc.InitWalletRequest{
		WalletPassphraseBin: testPassword,
		WalletSeed:          strings.Split(mnemonic, " "),
		SeedPassphraseBin:   pass,
		RecoveryWindow:      int32(testRecoveryWindow),
	}

	errChan := make(chan er.R, 1)

	go func() {
		_, err := service.InitWallet(ctx, req)
		if err != nil {
			errChan <- er.E(err)
			return
		}
		log.Debugf(">>> TestInitWallet [1] InitWallet() finished with success")
	}()

	// The same user passphrase, and also the plaintext cipher seed
	// should be sent over and match exactly.
	select {
	case err := <-errChan:
		t.Fatalf("InitWallet call failed: %v", err)

	case msg := <-service.InitMsgs:
		log.Debugf(">>> TestInitWallet [2] initialization message received")
		msgSeed := msg.Seed
		require.Equal(t, testPassword, msg.Passphrase)
		require.Equal(
			t, cipherSeed.Version(), msgSeed.Version(),
		)
		require.Equal(t, cipherSeed.Birthday(), msgSeed.Birthday())
		require.Equal(t, testRecoveryWindow, msg.RecoveryWindow)

		// Send a fake macaroon that should finish the async code above.
		log.Debugf(">>> TestInitWallet [3] fake macaroon sent back")
		service.MacResponseChan <- testMac

	case <-time.After(defaultTestTimeout):
		t.Fatalf("password not received")
	}

	// Create a wallet in testDir.
	createTestWallet(t, testDir, testNetParams)

	// Now calling InitWallet should fail, since a wallet already exists in
	// the directory.
	_, errr = service.InitWallet(ctx, req)
	require.Error(t, errr)
	require.Contains(t, errr.Error(), "wallet already exists")

	// Similarly, if we try to do GenSeed again, we should get an error as
	// the wallet already exists.
	_, errr = service.GenSeed(ctx, &lnrpc.GenSeedRequest{})
	require.Error(t, errr)
	require.Contains(t, errr.Error(), "wallet already exists")
}

// TestInitWalletInvalidCipherSeed tests that if we attempt to create a wallet
// with an invalid cipher seed, then we'll receive an error.
func TestCreateWalletInvalidEntropy(t *testing.T) {
	t.Parallel()

	log.Debugf(">>>>> running TestCreateWalletInvalidEntropy()")
	// testDir is empty, meaning wallet was not created from before.
	testDir, errr := ioutil.TempDir("", "testcreate")
	require.NoError(t, errr)
	defer func() {
		_ = os.RemoveAll(testDir)
	}()

	// Create new UnlockerService.
	service := walletunlocker.New(testDir, testNetParams, true, nil, "", testWalletFilename)

	// We'll attempt to init the wallet with an invalid cipher seed and
	// passphrase.
	req := &lnrpc.InitWalletRequest{
		WalletPassphraseBin: testPassword,
		WalletSeed:          []string{"invalid", "seed"},
		SeedPassphraseBin:   []byte("fake pass"),
	}

	ctx := context.Background()
	_, errr = service.InitWallet(ctx, req)
	require.Error(t, errr)
	require.Contains(t, errr.Error(), "Expected a 15 word seed")
}

// TestUnlockWallet checks that trying to unlock non-existing wallet fail, that
// unlocking existing wallet with wrong passphrase fails, and that unlocking
// existing wallet with correct passphrase succeeds.
func TestUnlockWallet(t *testing.T) {
	t.Parallel()

	log.Debugf(">>>>> running TestUnlockWallet()")
	// testDir is empty, meaning wallet was not created from before.
	testDir, errr := ioutil.TempDir("", "testunlock")
	require.NoError(t, errr)
	defer func() {
		_ = os.RemoveAll(testDir)
	}()

	// Create new UnlockerService.
	service := walletunlocker.New(testDir, testNetParams, true, nil, "", testWalletFilename)

	ctx := context.Background()
	req := &lnrpc.UnlockWalletRequest{
		WalletPassphraseBin: testPassword,
		RecoveryWindow:      int32(testRecoveryWindow),
	}

	// Should fail to unlock non-existing wallet.
	_, err := service.UnlockWallet(ctx, req)
	require.Error(t, err)
	require.Contains(t, err.Error(), "wallet not found")

	// Create a wallet we can try to unlock.
	createTestWallet(t, testDir, testNetParams)

	// Try unlocking this wallet with the wrong passphrase.
	wrongReq := &lnrpc.UnlockWalletRequest{
		WalletPassphraseBin: []byte("wrong-ofc"),
	}
	_, err = service.UnlockWallet(ctx, wrongReq)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid passphrase for master private key")

	// With the correct password, we should be able to unlock the wallet.
	errChan := make(chan er.R, 1)
	go func() {
		_, err := service.UnlockWallet(ctx, req)
		if err != nil {
			errChan <- er.E(err)
		}
	}()

	// Password and recovery window should be sent over the channel.
	select {
	case err := <-errChan:
		t.Fatalf("UnlockWallet call failed: %v", err)

	case unlockMsg := <-service.UnlockMsgs:
		require.Equal(t, testPassword, unlockMsg.Passphrase)
		require.Equal(t, testRecoveryWindow, unlockMsg.RecoveryWindow)
		require.Equal(t, true, unlockMsg.StatelessInit)

		// Send a fake macaroon that should be returned in the response
		// in the async code above.
		service.MacResponseChan <- testMac

	case <-time.After(defaultTestTimeout):
		t.Fatalf("password not received")
	}
}
