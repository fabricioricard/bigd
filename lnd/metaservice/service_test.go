////////////////////////////////////////////////////////////////////////////////
//	lnd/metaservice/service_test.go  -  Jan-26-2022  -  aldebap
//
//	unit tests for lnd/metaservice/service.go
////////////////////////////////////////////////////////////////////////////////

package metaservice

import (
	"context"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/btcutil/util"
	"github.com/pkt-cash/pktd/chaincfg"
	"github.com/pkt-cash/pktd/lnd/channeldb/kvdb"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
	"github.com/pkt-cash/pktd/lnd/lnwallet/btcwallet"
	"github.com/pkt-cash/pktd/lnd/macaroons"
	"github.com/pkt-cash/pktd/neutrino"
	"github.com/pkt-cash/pktd/pktlog/log"
	"github.com/pkt-cash/pktd/pktwallet/snacl"
	"github.com/pkt-cash/pktd/pktwallet/waddrmgr"
	"github.com/pkt-cash/pktd/pktwallet/wallet"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
	"github.com/stretchr/testify/require"
)

const (
	testWalletFilename string = "wallet.db"
)

var (
	testPassword = []byte("test-password")
	testSeed     = []byte("test-seed-123456789")

	testNetParams = &chaincfg.MainNetParams

	defaultRootKeyIDContext = macaroons.ContextWithRootKeyID(
		context.Background(), macaroons.DefaultRootKeyID,
	)
)

//	Test that as error occurs on an attempt to change the password for a non-existing  wallet
func TestChangePasswordForNonExistingWallet(t *testing.T) {
	t.Parallel()

	log.Debugf(">>>>> running TestChangePasswordForNonExistingWallet()")

	//	create a temporary directory, initialize an empty walletdb with an SPV chain
	//	namespace, and create a configuration for the ChainService
	testDir, err := ioutil.TempDir("", "neutrino")
	require.NoError(t, err)
	defer func() {
		_ = os.RemoveAll(testDir)
	}()

	db, errr := walletdb.Create("bdb", testDir+"/testNeutrino.db", true)
	util.RequireNoErr(t, errr)
	defer db.Close()

	config := neutrino.Config{
		DataDir:     testDir,
		Database:    db,
		ChainParams: *testNetParams,
	}

	testChainService, errr := neutrino.NewChainService(config)
	util.RequireNoErr(t, errr)

	//	create a new MetaService with our test file
	metaService := NewMetaService(testChainService)
	metaService.walletPath = btcwallet.NetworkDir(testDir, testNetParams)
	metaService.walletFile = testWalletFilename

	ctx := context.Background()

	//	changing the password to a non-existing wallet should fail
	log.Debugf("[1] attempt to change password for a non-existing wallet")

	newPassword := []byte("hunter2???")
	req := &lnrpc.ChangePasswordRequest{
		CurrentPasswordBin: testPassword,
		NewPassphraseBin:   newPassword,
	}

	_, err = metaService.ChangePassword(ctx, req)
	require.Error(t, err)
	require.Contains(t, err.Error(), "wallet not found")
}

//	Test that we can successfully change the wallet's password needed to unlock
//	it and rotate the root key for the macaroons in the same process.
/*
func TestChangeWalletPasswordNewRootkey(t *testing.T) {
	t.Parallel()

	log.Debugf(">>>>> running TestChangeWalletPasswordNewRootkey()")

	//	create a temporary directory, initialize an empty walletdb with an SPV chain
	//	namespace, and create a configuration for the ChainService
	testDir, err := ioutil.TempDir("", "neutrino")
	require.NoError(t, err)
	defer func() {
		_ = os.RemoveAll(testDir)
	}()

	walletDb, errr := walletdb.Create("bdb", testDir+"/testNeutrino.db", true)
	util.RequireNoErr(t, errr)
	defer walletDb.Close()

	config := neutrino.Config{
		DataDir:     testDir,
		Database:    walletDb,
		ChainParams: *testNetParams,
	}
	testChainService, errr := neutrino.NewChainService(config)
	util.RequireNoErr(t, errr)

	//	create a new MetaService with our test file
	metaService := NewMetaService(testChainService)
	metaService.walletPath = btcwallet.NetworkDir(testDir, testNetParams)
	metaService.walletFile = testWalletFilename

	ctx := context.Background()

	// Create a wallet to test changing the password.
	loader := createTestWallet(t, testDir, testNetParams)

	//	unload wallet
	errr = loader.UnloadWallet()
	util.RequireNoErr(t, errr)

	if metaService.Wallet != nil {
		log.Debugf("[2.5] wallet not nil")
	}

	//	changing the wallet's password using an incorrect current password should fail
	newPassword := []byte("hunter2???")

	log.Debugf("[3] attempt to change passwaord using an incorrect current password")
	wrongReq := &lnrpc.ChangePasswordRequest{
		CurrentPasswordBin: []byte("wrong-ofc"),
		NewPassphraseBin:   newPassword,
	}
	_, err = metaService.ChangePassword(ctx, wrongReq)
	require.Error(t, err)
	//require.Contains(t, err.Error(), "invalid passphrase for master public key")
	require.Contains(t, err.Error(), "unable to change wallet passphrase: ")

	//	changing the wallet's password using an invalid new password should fail
	log.Debugf("[4] attempt to change passwaord using an invalid new password")

	wrongReq.NewPassphraseBin = []byte("8")
	_, err = metaService.ChangePassword(ctx, wrongReq)
	require.Error(t, err)
	require.Contains(t, err.Error(), "custom password must have at least 8 characters")

	//	when providing the correct wallet's current password and a valid new password,
	//	the password change should succeed
	log.Debugf("[5] finally change passwaord")

	req := &lnrpc.ChangePasswordRequest{
		CurrentPasswordBin: testPassword,
		NewPassphraseBin:   newPassword,
	}

	_, errr = changePassword(metaService, testDir, req)
	util.RequireNoErr(t, errr)
}
*/

//	Test that we can successfully change the wallet's password needed to unlock
//	it and rotate the root key for the macaroons in the same process.
func TestChangeWalletPasswordWithWrongPassphrase(t *testing.T) {
	t.Parallel()

	log.Debugf(">>>>> running TestChangeWalletPasswordWithWrongPassphrase()")

	//	create a temporary directory, initialize an empty walletdb with an SPV chain
	//	namespace, and create a configuration for the ChainService
	testDir, err := ioutil.TempDir("", "neutrino")
	require.NoError(t, err)
	defer func() {
		_ = os.RemoveAll(testDir)
	}()

	walletDb, errr := walletdb.Create("bdb", testDir+"/testNeutrino.db", true)
	util.RequireNoErr(t, errr)
	defer walletDb.Close()

	config := neutrino.Config{
		DataDir:     testDir,
		Database:    walletDb,
		ChainParams: *testNetParams,
	}
	testChainService, errr := neutrino.NewChainService(config)
	util.RequireNoErr(t, errr)

	//	create a new MetaService with our test file
	metaService := NewMetaService(testChainService)
	metaService.walletPath = btcwallet.NetworkDir(testDir, testNetParams)
	metaService.walletFile = testWalletFilename

	ctx := context.Background()

	// Create a wallet to test changing the password.
	loader := createTestWallet(t, testDir, testNetParams)

	//	unload wallet
	errr = loader.UnloadWallet()
	util.RequireNoErr(t, errr)

	//	changing the wallet's password using an incorrect current password should fail
	newPassword := []byte("hunter2???")

	wrongReq := &lnrpc.ChangePasswordRequest{
		CurrentPasswordBin: []byte("wrong-ofc"),
		NewPassphraseBin:   newPassword,
	}
	_, err = metaService.ChangePassword(ctx, wrongReq)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unable to change wallet passphrase: ")
}

//	Test that we can successfully change the wallet's password needed to unlock
//	it and rotate the root key for the macaroons in the same process.
func TestChangeWalletPasswordNewRootkey(t *testing.T) {
	t.Parallel()

	log.Debugf(">>>>> running TestChangeWalletPasswordNewRootkey()")

	//	create a temporary directory, initialize an empty walletdb with an SPV chain
	//	namespace, and create a configuration for the ChainService
	testDir, err := ioutil.TempDir("", "neutrino")
	require.NoError(t, err)
	defer func() {
		_ = os.RemoveAll(testDir)
	}()

	walletDb, errr := walletdb.Create("bdb", testDir+"/testNeutrino.db", true)
	util.RequireNoErr(t, errr)
	defer walletDb.Close()

	config := neutrino.Config{
		DataDir:     testDir,
		Database:    walletDb,
		ChainParams: *testNetParams,
	}
	testChainService, errr := neutrino.NewChainService(config)
	util.RequireNoErr(t, errr)

	//	create a new MetaService with our test file
	metaService := NewMetaService(testChainService)
	metaService.walletPath = btcwallet.NetworkDir(testDir, testNetParams)
	metaService.walletFile = testWalletFilename

	// Create a wallet to test changing the password.
	loader := createTestWallet(t, testDir, testNetParams)

	//	unload wallet
	errr = loader.UnloadWallet()
	util.RequireNoErr(t, errr)

	//	when providing the correct wallet's current password and a valid new password,
	//	the password change should succeed
	newPassword := []byte("hunter2???")

	log.Debugf("[5] finally change passwaord")

	req := &lnrpc.ChangePasswordRequest{
		CurrentPasswordBin: testPassword,
		NewPassphraseBin:   newPassword,
	}

	_, errr = changePassword(metaService, testDir, req)
	util.RequireNoErr(t, errr)
}

//	execute a password change
func changePassword(metaService *MetaService, macTestDir string, req *lnrpc.ChangePasswordRequest) (*lnrpc.ChangePasswordResponse, er.R) {

	//	when providing the correct wallet's current password and a valid new password,
	//	the password change should succeed
	ctx := context.Background()
	response, err := metaService.ChangePassword(ctx, req)
	if err != nil {
		return nil, er.Errorf("could not change password: %v", err)
	}

	//	close the macaroon DB and try to open it and read the root key with the
	//	new password
	store, errr := openOrCreateTestMacStore(macTestDir, &testPassword, testNetParams)
	if errr != nil {
		return nil, er.Errorf("could not create test store: %v", err)
	}
	_, _, err = store.RootKey(defaultRootKeyIDContext)
	if err != nil {
		return nil, er.Errorf("could not get root key: %v", errr)
	}
	errr = store.Close()
	if errr != nil {
		return nil, er.Errorf("could not close store: %v", err)
	}

	// Do cleanup now. Since we are in a go func, the defer at the
	// top of the outer would not work, because it would delete
	// the directory before we could check the content in here.
	err = os.RemoveAll(macTestDir)
	if errr != nil {
		return nil, er.E(err)
	}

	return response, nil
}

// openOrCreateTestMacStore opens or creates a bbolt DB and then initializes a
// root key storage for that DB and then unlocks it, creating a root key in the
// process.
func openOrCreateTestMacStore(tempDir string, pw *[]byte,
	netParams *chaincfg.Params) (*macaroons.RootKeyStorage, er.R) {

	netDir := btcwallet.NetworkDir(tempDir, netParams)
	errr := os.MkdirAll(netDir, 0700)
	if errr != nil {
		return nil, er.E(errr)
	}
	db, err := kvdb.Create(
		kvdb.BoltBackendName, path.Join(netDir, macaroons.DBFilename),
		true,
	)
	if err != nil {
		return nil, err
	}

	store, err := macaroons.NewRootKeyStorage(db)
	if err != nil {
		_ = db.Close()
		return nil, err
	}

	err = store.CreateUnlock(pw)
	if err != nil {
		_ = store.Close()
		return nil, err
	}
	_, _, errr = store.RootKey(defaultRootKeyIDContext)
	if errr != nil {
		_ = store.Close()
		return nil, er.E(errr)
	}

	return store, nil
}

func createTestWallet(t *testing.T, dir string, netParams *chaincfg.Params) *wallet.Loader {
	return createTestWalletWithPw(t, []byte(wallet.InsecurePubPassphrase), testPassword, dir, netParams)
}

func createTestWalletWithPw(t *testing.T, pubPw, privPw []byte, dir string, netParams *chaincfg.Params) *wallet.Loader {

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
	log.Debugf("[1] wallet path: %s", realWalletPathname)
	walletFileExists := true
	_, errr := os.Stat(realWalletPathname)
	if err != nil {
		if os.IsNotExist(errr) {
			walletFileExists = false
		} else {
			require.NoError(t, errr)
		}
	}
	log.Debugf("[2] after loader.CreateNewWallet() the wallet file exists: %t", walletFileExists)

	return loader
}
