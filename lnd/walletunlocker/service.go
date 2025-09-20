package walletunlocker

import (
	"context"
	"strings"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/chaincfg"
	"github.com/pkt-cash/pktd/lnd/chanbackup"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
	"github.com/pkt-cash/pktd/lnd/lnwallet/btcwallet"
	"github.com/pkt-cash/pktd/pktwallet/wallet"
	"github.com/pkt-cash/pktd/pktwallet/wallet/seedwords"
)

var (
	// ErrUnlockTimeout signals that we did not get the expected unlock
	// message before the timeout occurred.
	ErrUnlockTimeout = er.GenericErrorType.CodeWithDetail("ErrUnlockTimeout",
		"got no unlock message before timeout")
)

// ChannelsToRecover wraps any set of packed (serialized+encrypted) channel
// back ups together. These can be passed in when unlocking the wallet, or
// creating a new wallet for the first time with an existing seed.
type ChannelsToRecover struct {
	// PackedMultiChanBackup is an encrypted and serialized multi-channel
	// backup.
	PackedMultiChanBackup chanbackup.PackedMulti

	// PackedSingleChanBackups is a series of encrypted and serialized
	// single-channel backup for one or more channels.
	PackedSingleChanBackups chanbackup.PackedSingles
}

// WalletInitMsg is a message sent by the UnlockerService when a user wishes to
// set up the internal wallet for the first time. The user MUST provide a
// passphrase, but is also able to provide their own source of entropy. If
// provided, then this source of entropy will be used to generate the wallet's
// HD seed. Otherwise, the wallet will generate one itself.
type WalletInitMsg struct {
	// Passphrase is the passphrase that will be used to encrypt the wallet
	// itself. This MUST be at least 8 characters.
	Passphrase []byte

	// WalletSeed is the deciphered cipher seed that the wallet should use
	// to initialize itself.
	Seed *seedwords.Seed

	// RecoveryWindow is the address look-ahead used when restoring a seed
	// with existing funds. A recovery window zero indicates that no
	// recovery should be attempted, such as after the wallet's initial
	// creation.
	RecoveryWindow uint32

	// ChanBackups a set of static channel backups that should be received
	// after the wallet has been initialized.
	ChanBackups ChannelsToRecover

	// StatelessInit signals that the user requested the daemon to be
	// initialized stateless, which means no unencrypted macaroons should be
	// written to disk.
	StatelessInit bool

	WalletName string
}

// WalletUnlockMsg is a message sent by the UnlockerService when a user wishes
// to unlock the internal wallet after initial setup. The user can optionally
// specify a recovery window, which will resume an interrupted rescan for used
// addresses.
type WalletUnlockMsg struct {
	// Passphrase is the passphrase that will be used to encrypt the wallet
	// itself. This MUST be at least 8 characters.
	Passphrase []byte

	// RecoveryWindow is the address look-ahead used when restoring a seed
	// with existing funds. A recovery window zero indicates that no
	// recovery should be attempted, such as after the wallet's initial
	// creation, but before any addresses have been created.
	RecoveryWindow uint32

	// Wallet is the loaded and unlocked Wallet. This is returned through
	// the channel to avoid it being unlocked twice (once to check if the
	// password is correct, here in the WalletUnlocker and again later when
	// lnd actually uses it). Because unlocking involves scrypt which is
	// resource intensive, we want to avoid doing it twice.
	Wallet *wallet.Wallet

	// ChanBackups a set of static channel backups that should be received
	// after the wallet has been unlocked.
	ChanBackups ChannelsToRecover

	// UnloadWallet is a function for unloading the wallet, which should
	// be called on shutdown.
	UnloadWallet func() er.R

	// StatelessInit signals that the user requested the daemon to be
	// initialized stateless, which means no unencrypted macaroons should be
	// written to disk.
	StatelessInit bool
}

// UnlockerService implements the WalletUnlocker service used to provide lnd
// with a password for wallet encryption at startup. Additionally, during
// initial setup, users can provide their own source of entropy which will be
// used to generate the seed that's ultimately used within the wallet.
type UnlockerService struct {
	// InitMsgs is a channel that carries all wallet init messages.
	InitMsgs chan *WalletInitMsg

	// UnlockMsgs is a channel where unlock parameters provided by the rpc
	// client to be used to unlock and decrypt an existing wallet will be
	// sent.
	UnlockMsgs chan *WalletUnlockMsg

	// MacResponseChan is the channel for sending back the admin macaroon to
	// the WalletUnlocker service.
	MacResponseChan chan []byte

	chainDir       string
	noFreelistSync bool
	netParams      *chaincfg.Params

	// macaroonFiles is the path to the three generated macaroons with
	// different access permissions. These might not exist in a stateless
	// initialization of lnd.
	macaroonFiles []string

	walletFile string
	walletPath string
}

var _ lnrpc.WalletUnlockerServer = (*UnlockerService)(nil)

// New creates and returns a new UnlockerService.
func New(chainDir string, params *chaincfg.Params, noFreelistSync bool,
	macaroonFiles []string, walletPath string, walletFilename string) *UnlockerService {

	return &UnlockerService{
		InitMsgs:   make(chan *WalletInitMsg, 1),
		UnlockMsgs: make(chan *WalletUnlockMsg, 1),

		// Make sure we buffer the channel is buffered so the main lnd
		// goroutine isn't blocking on writing to it.
		MacResponseChan: make(chan []byte, 1),
		chainDir:        chainDir,
		noFreelistSync:  noFreelistSync,
		netParams:       params,
		macaroonFiles:   macaroonFiles,
		walletFile:      walletFilename,
		walletPath:      walletPath,
	}
}

func (u *UnlockerService) GenSeed(_ context.Context,
	in *lnrpc.GenSeedRequest) (*lnrpc.GenSeedResponse, error) {
	//	TODO: should replace the nil context by context.TODO()
	res, err := u.GenSeed0(nil, in)
	return res, er.Native(err)
}

// GenSeed is the first method that should be used to instantiate a new lnd
// instance. This method allows a caller to generate a new aezeed cipher seed
// given an optional passphrase. If provided, the passphrase will be necessary
// to decrypt the cipherseed to expose the internal wallet seed.
//
// Once the cipherseed is obtained and verified by the user, the InitWallet
// method should be used to commit the newly generated seed, and create the
// wallet.
func (u *UnlockerService) GenSeed0(_ context.Context,
	in *lnrpc.GenSeedRequest) (*lnrpc.GenSeedResponse, er.R) {

	//var entropy [aezeed.EntropySize]byte

	if len(in.SeedEntropy) != 0 {
		return nil, er.Errorf("seed input entropy is not supported")
	}

	// Now that we have our set of entropy, we'll create a new cipher seed
	// instance.
	//
	cipherSeed, err := seedwords.RandomSeed()
	if err != nil {
		return nil, err
	}

	//	fetch seed passphrase from request
	var seedPassphrase []byte

	if len(in.SeedPassphraseBin) > 0 {
		seedPassphrase = in.SeedPassphraseBin
	} else {
		if len(in.SeedPassphrase) > 0 {
			seedPassphrase = []byte(in.SeedPassphrase)
		}
	}
	encipheredSeed := cipherSeed.Encrypt(seedPassphrase)

	mnemonic, err := encipheredSeed.Words("english")
	if err != nil {
		return nil, err
	}

	return &lnrpc.GenSeedResponse{
		Seed: strings.Split(mnemonic, " "),
	}, nil
}

// extractChanBackups is a helper function that extracts the set of channel
// backups from the proto into a format that we'll pass to higher level
// sub-systems.
func extractChanBackups(chanBackups *lnrpc.ChanBackupSnapshot) *ChannelsToRecover {
	// If there aren't any populated channel backups, then we can exit
	// early as there's nothing to extract.
	if chanBackups == nil || (chanBackups.SingleChanBackups == nil &&
		chanBackups.MultiChanBackup == nil) {
		return nil
	}

	// Now that we know there's at least a single back up populated, we'll
	// extract the multi-chan backup (if it's there).
	var backups ChannelsToRecover
	if chanBackups.MultiChanBackup != nil {
		multiBackup := chanBackups.MultiChanBackup
		backups.PackedMultiChanBackup = chanbackup.PackedMulti(
			multiBackup.MultiChanBackup,
		)
	}

	if chanBackups.SingleChanBackups == nil {
		return &backups
	}

	// Finally, we can extract all the single chan backups as well.
	for _, backup := range chanBackups.SingleChanBackups.ChanBackups {
		singleChanBackup := backup.ChanBackup

		backups.PackedSingleChanBackups = append(
			backups.PackedSingleChanBackups, singleChanBackup,
		)
	}

	return &backups
}

func (u *UnlockerService) InitWallet(ctx context.Context,
	in *lnrpc.InitWalletRequest) (*lnrpc.InitWalletResponse, error) {
	res, err := u.InitWallet0(ctx, in)
	return res, er.Native(err)
}

// InitWallet is used when lnd is starting up for the first time to fully
// initialize the daemon and its internal wallet. At the very least a wallet
// password must be provided. This will be used to encrypt sensitive material
// on disk.
//
// In the case of a recovery scenario, the user can also specify their aezeed
// mnemonic and passphrase. If set, then the daemon will use this prior state
// to initialize its internal wallet.
//
// Alternatively, this can be used along with the GenSeed RPC to obtain a
// seed, then present it to the user. Once it has been verified by the user,
// the seed can be fed into this RPC in order to commit the new wallet.
func (u *UnlockerService) InitWallet0(ctx context.Context,
	in *lnrpc.InitWalletRequest) (*lnrpc.InitWalletResponse, er.R) {

	//	fetch wallet passphrase from request
	var walletPassphrase []byte

	if len(in.WalletPassphraseBin) > 0 {
		walletPassphrase = in.WalletPassphraseBin
	} else {
		if len(in.WalletPassphrase) > 0 {
			walletPassphrase = []byte(in.WalletPassphrase)
		}
	}

	// Require that the recovery window be non-negative.
	recoveryWindow := in.RecoveryWindow
	if recoveryWindow < 0 {
		return nil, er.Errorf("recovery window %d must be "+
			"non-negative", recoveryWindow)
	}

	// We'll then open up the directory that will be used to store the
	// wallet's files so we can check if the wallet already exists.
	netDir := btcwallet.NetworkDir(u.chainDir, u.netParams)
	if u.walletPath != "" {
		netDir = u.walletPath
	}
	//If wallet_name passed use it instead of default
	walletFile := u.walletFile
	if in.WalletName != "" {
		walletFile = in.WalletName
	}
	loader := wallet.NewLoader(u.netParams, netDir, walletFile, u.noFreelistSync, uint32(recoveryWindow))

	walletExists, err := loader.WalletExists()
	if err != nil {
		return nil, err
	}

	// If the wallet already exists, then we'll exit early as we can't
	// create the wallet if it already exists!
	if walletExists {
		return nil, er.Errorf("wallet already exists")
	}

	mnemonic := strings.Join(in.WalletSeed, " ")
	seedEnc, err := seedwords.SeedFromWords(mnemonic)
	if err != nil {
		return nil, err
	}

	//	fetch seed passphrase from request
	var seedPassphrase []byte

	if len(in.SeedPassphraseBin) > 0 {
		seedPassphrase = in.SeedPassphraseBin
	} else {
		if len(in.SeedPassphrase) > 0 {
			seedPassphrase = []byte(in.SeedPassphrase)
		}
	}

	seed, err := seedEnc.Decrypt(seedPassphrase, false)
	if err != nil {
		return nil, err
	}

	// With the cipher seed deciphered, and the auth service created, we'll
	// now send over the wallet password and the seed. This will allow the
	// daemon to initialize itself and startup.
	initMsg := &WalletInitMsg{
		Passphrase:     walletPassphrase,
		Seed:           seed,
		RecoveryWindow: uint32(recoveryWindow),
		StatelessInit:  true,
		WalletName:     walletFile,
	}

	// Before we return the unlock payload, we'll check if we can extract
	// any channel backups to pass up to the higher level sub-system.
	chansToRestore := extractChanBackups(in.ChannelBackups)
	if chansToRestore != nil {
		initMsg.ChanBackups = *chansToRestore
	}

	// Deliver the initialization message back to the main daemon.
	select {
	case u.InitMsgs <- initMsg:
		// We need to read from the channel to let the daemon continue
		// its work and to get the admin macaroon. Once the response
		// arrives, we directly forward it to the client.
		select {
		case <-u.MacResponseChan:
			return &lnrpc.InitWalletResponse{}, nil

		case <-ctx.Done():
			return nil, ErrUnlockTimeout.Default()
		}

	case <-ctx.Done():
		return nil, ErrUnlockTimeout.Default()
	}
}

func (u *UnlockerService) UnlockWallet(ctx context.Context,
	in *lnrpc.UnlockWalletRequest) (*lnrpc.UnlockWalletResponse, error) {
	res, err := u.UnlockWallet0(ctx, in)
	return res, er.Native(err)
}

// UnlockWallet sends the password provided by the incoming UnlockWalletRequest
// over the UnlockMsgs channel in case it successfully decrypts an existing
// wallet found in the chain's wallet database directory.
func (u *UnlockerService) UnlockWallet0(ctx context.Context,
	in *lnrpc.UnlockWalletRequest) (*lnrpc.UnlockWalletResponse, er.R) {

	//	fetch wallet passphrase from request
	var walletPassphrase []byte

	if len(in.WalletPassphraseBin) > 0 {
		walletPassphrase = in.WalletPassphraseBin
	} else {
		if len(in.WalletPassphrase) > 0 {
			walletPassphrase = []byte(in.WalletPassphrase)
		}
	}

	pubpassword := []byte(wallet.InsecurePubPassphrase)

	recoveryWindow := uint32(in.RecoveryWindow)

	netDir := btcwallet.NetworkDir(u.chainDir, u.netParams)
	if u.walletPath != "" {
		netDir = u.walletPath
	}
	walletFile := u.walletFile
	if in.WalletName != "" {
		walletFile = in.WalletName
	}
	loader := wallet.NewLoader(u.netParams, netDir, walletFile, u.noFreelistSync, recoveryWindow)

	// Check if wallet already exists.
	walletExists, err := loader.WalletExists()
	if err != nil {
		return nil, err
	}

	if !walletExists {
		// Cannot unlock a wallet that does not exist!
		return nil, er.Errorf("wallet not found at path [%s/wallet.db]", netDir)
	}

	// Try opening the existing wallet with the provided password.
	unlockedWallet, err := loader.OpenExistingWallet(pubpassword, false)
	if err != nil {
		// Could not open wallet, most likely this means that provided
		// password was incorrect.
		return nil, err
	}
	//Also test against private password
	err = unlockedWallet.Unlock(walletPassphrase, nil)
	if err != nil {
		//unload wallet so future unlock calls can be processed
		loader.UnloadWallet()
		return nil, err
	}
	// We successfully opened the wallet and pass the instance back to
	// avoid it needing to be unlocked again.
	walletUnlockMsg := &WalletUnlockMsg{
		Passphrase:     walletPassphrase,
		RecoveryWindow: recoveryWindow,
		Wallet:         unlockedWallet,
		UnloadWallet:   loader.UnloadWallet,
		StatelessInit:  true,
	}

	// Before we return the unlock payload, we'll check if we can extract
	// any channel backups to pass up to the higher level sub-system.
	chansToRestore := extractChanBackups(in.ChannelBackups)
	if chansToRestore != nil {
		walletUnlockMsg.ChanBackups = *chansToRestore
	}

	// At this point we were able to open the existing wallet with the
	// provided password. We send the password over the UnlockMsgs
	// channel, such that it can be used by lnd to open the wallet.
	select {
	case u.UnlockMsgs <- walletUnlockMsg:
		// We need to read from the channel to let the daemon continue
		// its work. But we don't need the returned macaroon for this
		// operation, so we read it but then discard it.
		select {
		case <-u.MacResponseChan:
			return &lnrpc.UnlockWalletResponse{}, nil

		case <-ctx.Done():
			return nil, ErrUnlockTimeout.Default()
		}

	case <-ctx.Done():
		return nil, ErrUnlockTimeout.Default()
	}
}
