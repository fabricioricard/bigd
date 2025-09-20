package metaservice

import (
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/pkt-cash/pktd/btcjson"
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/chaincfg"
	"github.com/pkt-cash/pktd/connmgr/banmgr"
	"github.com/pkt-cash/pktd/lnd/lncfg"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
	"github.com/pkt-cash/pktd/lnd/lnwallet"
	"github.com/pkt-cash/pktd/lnd/lnwallet/btcwallet"
	"github.com/pkt-cash/pktd/lnd/macaroons"
	"github.com/pkt-cash/pktd/neutrino"
	"github.com/pkt-cash/pktd/pktlog/log"
	"github.com/pkt-cash/pktd/pktwallet/waddrmgr"
	"github.com/pkt-cash/pktd/pktwallet/wallet"
	"google.golang.org/grpc"
)

type MetaService struct {
	Neutrino *neutrino.ChainService
	Wallet   *wallet.Wallet

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

var _ lnrpc.MetaServiceServer = (*MetaService)(nil)

// New creates and returns a new MetaService
func NewMetaService(neutrino *neutrino.ChainService) *MetaService {
	return &MetaService{
		Neutrino: neutrino,
	}
}

func (m *MetaService) SetWallet(wallet *wallet.Wallet) {
	m.Wallet = wallet
}

func (m *MetaService) Init(MacResponseChan chan []byte, chainDir string,
	noFreelistSync bool, netParams *chaincfg.Params, macaroonFiles []string, walletFile, walletPath string) {
	m.MacResponseChan = MacResponseChan
	m.chainDir = chainDir
	m.netParams = netParams
	m.macaroonFiles = macaroonFiles
	m.walletFile = walletFile
	m.walletPath = walletPath
}

func (m *MetaService) GetInfo2(ctx context.Context,
	in *lnrpc.GetInfo2Request) (*lnrpc.GetInfo2Response, error) {
	res, err := m.GetInfo20(ctx, in)
	return res, er.Native(err)
}

func getClientConn(ctx *context.Context, skipMacaroons bool) *grpc.ClientConn {
	var defaultRPCPort = "10009"
	var maxMsgRecvSize = grpc.MaxCallRecvMsgSize(1 * 1024 * 1024 * 200)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	genericDialer := lncfg.ClientAddressDialer(defaultRPCPort)
	opts = append(opts, grpc.WithContextDialer(genericDialer))
	opts = append(opts, grpc.WithDefaultCallOptions(maxMsgRecvSize))

	conn, errr := grpc.Dial("localhost:10009", opts...)
	if errr != nil {
		log.Errorf("unable to connect to RPC server: %v", errr)
		return nil
	}

	return conn
}

func getClient(ctx *context.Context) (lnrpc.LightningClient, func()) {
	conn := getClientConn(ctx, false)

	cleanUp := func() {
		conn.Close()
	}

	return lnrpc.NewLightningClient(conn), cleanUp
}

func (m *MetaService) GetInfo20(ctx context.Context,
	in *lnrpc.GetInfo2Request) (*lnrpc.GetInfo2Response, er.R) {

	var ni lnrpc.NeutrinoInfo
	neutrinoPeers := m.Neutrino.Peers()
	for i := range neutrinoPeers {
		var peerDesc lnrpc.PeerDesc
		neutrinoPeer := neutrinoPeers[i]

		peerDesc.BytesReceived = neutrinoPeer.BytesReceived()
		peerDesc.BytesSent = neutrinoPeer.BytesSent()
		peerDesc.LastRecv = neutrinoPeer.LastRecv().String()
		peerDesc.LastSend = neutrinoPeer.LastSend().String()
		peerDesc.Connected = neutrinoPeer.Connected()
		peerDesc.Addr = neutrinoPeer.Addr()
		peerDesc.Inbound = neutrinoPeer.Inbound()
		na := neutrinoPeer.NA()
		if na != nil {
			peerDesc.Na = na.IP.String() + ":" + strconv.Itoa(int(na.Port))
		}
		peerDesc.Id = neutrinoPeer.ID()
		peerDesc.UserAgent = neutrinoPeer.UserAgent()
		peerDesc.Services = neutrinoPeer.Services().String()
		peerDesc.VersionKnown = neutrinoPeer.VersionKnown()
		peerDesc.AdvertisedProtoVer = neutrinoPeer.Describe().AdvertisedProtoVer
		peerDesc.ProtocolVersion = neutrinoPeer.ProtocolVersion()
		peerDesc.SendHeadersPreferred = neutrinoPeer.Describe().SendHeadersPreferred
		peerDesc.VerAckReceived = neutrinoPeer.VerAckReceived()
		peerDesc.WitnessEnabled = neutrinoPeer.Describe().WitnessEnabled
		peerDesc.WireEncoding = strconv.Itoa(int(neutrinoPeer.Describe().WireEncoding))
		peerDesc.TimeOffset = neutrinoPeer.TimeOffset()
		peerDesc.TimeConnected = neutrinoPeer.Describe().TimeConnected.String()
		peerDesc.StartingHeight = neutrinoPeer.StartingHeight()
		peerDesc.LastBlock = neutrinoPeer.LastBlock()
		if neutrinoPeer.LastAnnouncedBlock() != nil {
			peerDesc.LastAnnouncedBlock = neutrinoPeer.LastAnnouncedBlock().CloneBytes()
		}
		peerDesc.LastPingNonce = neutrinoPeer.LastPingNonce()
		peerDesc.LastPingTime = neutrinoPeer.LastPingTime().String()
		peerDesc.LastPingMicros = neutrinoPeer.LastPingMicros()

		ni.Peers = append(ni.Peers, &peerDesc)
	}
	m.Neutrino.BanMgr().ForEachIp(func(bi banmgr.BanInfo) er.R {
		ban := lnrpc.NeutrinoBan{}
		ban.Addr = bi.Addr
		ban.Reason = bi.Reason
		ban.EndTime = bi.BanExpiresTime.String()
		ban.BanScore = bi.BanScore

		ni.Bans = append(ni.Bans, &ban)
		return nil
	})

	neutrionoQueries := m.Neutrino.GetActiveQueries()
	for i := range neutrionoQueries {
		nq := lnrpc.NeutrinoQuery{}
		query := neutrionoQueries[i]
		nq.Peer = query.Peer.String()
		nq.Command = query.Command
		nq.ReqNum = query.ReqNum
		nq.CreateTime = query.CreateTime
		nq.LastRequestTime = query.LastRequestTime
		nq.LastResponseTime = query.LastResponseTime

		ni.Queries = append(ni.Queries, &nq)
	}

	bb, err := m.Neutrino.BestBlock()
	if err != nil {
		return nil, err
	}
	ni.BlockHash = bb.Hash.String()
	ni.Height = bb.Height
	ni.BlockTimestamp = bb.Timestamp.String()
	ni.IsSyncing = !m.Neutrino.IsCurrent()

	mgrStamp := waddrmgr.BlockStamp{}
	walletInfo := &lnrpc.WalletInfo{}

	if m.Wallet != nil {
		mgrStamp = m.Wallet.Manager.SyncedTo()
		walletStats := &lnrpc.WalletStats{}
		m.Wallet.ReadStats(func(ws *btcjson.WalletStats) {
			walletStats.MaintenanceInProgress = ws.MaintenanceInProgress
			walletStats.MaintenanceName = ws.MaintenanceName
			walletStats.MaintenanceCycles = int32(ws.MaintenanceCycles)
			walletStats.MaintenanceLastBlockVisited = int32(ws.MaintenanceLastBlockVisited)
			walletStats.Syncing = ws.Syncing
			if ws.SyncStarted != nil {
				walletStats.SyncStarted = ws.SyncStarted.String()
			}
			walletStats.SyncRemainingSeconds = ws.SyncRemainingSeconds
			walletStats.SyncCurrentBlock = ws.SyncCurrentBlock
			walletStats.SyncFrom = ws.SyncFrom
			walletStats.SyncTo = ws.SyncTo
			walletStats.BirthdayBlock = ws.BirthdayBlock
		})
		walletInfo = &lnrpc.WalletInfo{
			CurrentBlockHash:      mgrStamp.Hash.String(),
			CurrentHeight:         mgrStamp.Height,
			CurrentBlockTimestamp: mgrStamp.Timestamp.String(),
			WalletVersion:         int32(waddrmgr.LatestMgrVersion),
			WalletStats:           walletStats,
		}
	} else {
		walletInfo = nil
	}
	//Get Lightning info

	ctxb := context.Background()
	client, cleanUp := getClient(&ctx)
	defer cleanUp()
	inforeq := &lnrpc.GetInfoRequest{}
	inforesp, infoerr := client.GetInfo(ctxb, inforeq)
	if infoerr != nil {
		inforesp = nil
	}

	return &lnrpc.GetInfo2Response{
		Neutrino:  &ni,
		Wallet:    walletInfo,
		Lightning: inforesp,
	}, nil
}

func (u *MetaService) ChangePassword(ctx context.Context,
	in *lnrpc.ChangePasswordRequest) (*lnrpc.ChangePasswordResponse, error) {
	res, err := u.ChangePassword0(ctx, in)
	return res, er.Native(err)
}

// ChangePassword changes the password of the wallet and sends the new password
// across the UnlockPasswords channel to automatically unlock the wallet if
// successful.
func (m *MetaService) ChangePassword0(ctx context.Context,
	in *lnrpc.ChangePasswordRequest) (*lnrpc.ChangePasswordResponse, er.R) {

	//	fetch current wallet passphrase from request
	var walletPassphrase []byte

	if len(in.CurrentPasswordBin) > 0 {
		walletPassphrase = in.CurrentPasswordBin
	} else {
		if len(in.CurrentPassphrase) > 0 {
			walletPassphrase = []byte(in.CurrentPassphrase)
		} else {
			// If the current password is blank, we'll assume the user is coming
			// from a --noseedbackup state, so we'll use the default passwords.
			walletPassphrase = []byte(lnwallet.DefaultPrivatePassphrase)
		}
	}

	//	fetch new wallet passphrase from request
	var newWalletPassphrase []byte

	if len(in.NewPassphraseBin) > 0 {
		newWalletPassphrase = in.NewPassphraseBin
	} else {
		if len(in.NewPassphrase) > 0 {
			newWalletPassphrase = []byte(in.NewPassphrase)
		} else {
			newWalletPassphrase = []byte(lnwallet.DefaultPrivatePassphrase)
		}
	}

	publicPw := []byte(wallet.InsecurePubPassphrase)
	newPubPw := []byte(wallet.InsecurePubPassphrase)

	walletFile := m.walletFile
	if m.Wallet == nil || m.Wallet.Locked() {
		if in.WalletName != "" {
			walletFile = in.WalletName
		}
		loader := wallet.NewLoader(m.netParams, m.walletPath, walletFile, m.noFreelistSync, 0)

		// First, we'll make sure the wallet exists for the specific chain and
		// network.
		walletExists, err := loader.WalletExists()
		if err != nil {
			return nil, err
		}

		if !walletExists {
			return nil, er.New("wallet not found")
		}

		// Load the existing wallet in order to proceed with the password change.
		w, err := loader.OpenExistingWallet(publicPw, false)
		if err != nil {
			return nil, err
		}
		m.Wallet = w
		// Now that we've opened the wallet, we need to close it in case of an
		// error. But not if we succeed, then the caller must close it.
		orderlyReturn := false
		defer func() {
			if !orderlyReturn {
				_ = loader.UnloadWallet()
			}
		}()

		// Before we actually change the password, we need to check if all flags
		// were set correctly. The content of the previously generated macaroon
		// files will become invalid after we generate a new root key. So we try
		// to delete them here and they will be recreated during normal startup
		// later. If they are missing, this is only an error if the
		// stateless_init flag was not set.

		//	since we don't have macarrons anymore, 'stateless_init' will always be true
		for _, file := range m.macaroonFiles {
			err := os.Remove(file)
			if err != nil {
				return nil, er.Errorf("could not remove macaroon file: %v", err)
			}
		}
	} else if (in.WalletName != "") && (in.WalletName != m.walletFile) {
		walletFile = in.WalletName
		loader := wallet.NewLoader(m.netParams, m.walletPath, walletFile, m.noFreelistSync, 0)

		// First, we'll make sure the wallet exists for the specific chain and
		// network.
		walletExists, err := loader.WalletExists()
		if err != nil {
			return nil, err
		}

		if !walletExists {
			return nil, er.New("wallet not found")
		}

		// Load the existing wallet in order to proceed with the password change.
		w, err := loader.OpenExistingWallet(publicPw, false)
		if err != nil {
			return nil, err
		}
		m.Wallet = w
		// Now that we've opened the wallet, we need to close it in case of an
		// error. But not if we succeed, then the caller must close it.
		orderlyReturn := false
		defer func() {
			if !orderlyReturn {
				_ = loader.UnloadWallet()
			}
		}()
		// Before we actually change the password, we need to check if all flags
		// were set correctly. The content of the previously generated macaroon
		// files will become invalid after we generate a new root key. So we try
		// to delete them here and they will be recreated during normal startup
		// later. If they are missing, this is only an error if the
		// stateless_init flag was not set.

		//	since we don't have macarrons anymore, 'stateless_init' will always be true
		for _, file := range m.macaroonFiles {
			err := os.Remove(file)
			if err != nil {
				return nil, er.Errorf("could not remove macaroon file: %v", err)
			}
		}
	}

	// Attempt to change both the public and private passphrases for the
	// wallet. This will be done atomically in order to prevent one
	// passphrase change from being successful and not the other.
	err := m.Wallet.ChangePassphrases(
		publicPw, newPubPw, walletPassphrase, newWalletPassphrase,
	)
	if err != nil {
		return nil, er.Errorf("unable to change wallet passphrase: "+
			"%v", err)
	}

	adminMac := []byte{}
	// Check if macaroonFiles is populated, if not it's due to noMacaroon flag is set
	//so we do not need the service
	if len(m.macaroonFiles) > 0 {
		netDir := btcwallet.NetworkDir(m.chainDir, m.netParams)
		// The next step is to load the macaroon database, change the password
		// then close it again.
		// Attempt to open the macaroon DB, unlock it and then change
		// the passphrase.
		macaroonService, err := macaroons.NewService(
			netDir, "lnd", true,
		)
		if err != nil {
			return nil, err
		}

		err = macaroonService.CreateUnlock(&walletPassphrase)
		if err != nil {
			closeErr := macaroonService.Close()
			if closeErr != nil {
				return nil, er.Errorf("could not create unlock: %v "+
					"--> follow-up error when closing: %v", err,
					closeErr)
			}
			return nil, err
		}
		err = macaroonService.ChangePassword(walletPassphrase, newWalletPassphrase)
		if err != nil {
			closeErr := macaroonService.Close()
			if closeErr != nil {
				return nil, er.Errorf("could not change password: %v "+
					"--> follow-up error when closing: %v", err,
					closeErr)
			}
			return nil, err
		}

		err = macaroonService.Close()
		if err != nil {
			return nil, er.Errorf("could not close macaroon service: %v",
				err)
		}
		adminMac = <-m.MacResponseChan
	}
	_ = adminMac

	return &lnrpc.ChangePasswordResponse{}, nil
}

func (u *MetaService) CheckPassword(ctx context.Context, req *lnrpc.CheckPasswordRequest) (*lnrpc.CheckPasswordResponse, error) {

	res, err := u.CheckPassword0(ctx, req)

	return res, er.Native(err)
}

//	CheckPassword just verifies if the password of the wallet is valid, and is
//	meant to be used independent of wallet's state being unlocked or locked.
func (m *MetaService) CheckPassword0(ctx context.Context, req *lnrpc.CheckPasswordRequest) (*lnrpc.CheckPasswordResponse, er.R) {

	//	fetch current wallet passphrase from request
	var walletPassphrase []byte

	if len(req.WalletPasswordBin) > 0 {
		walletPassphrase = req.WalletPasswordBin
	} else {
		if len(req.WalletPassphrase) > 0 {
			walletPassphrase = []byte(req.WalletPassphrase)
		} else {
			// If the current password is blank, we'll assume the user is coming
			// from a --noseedbackup state, so we'll use the default passwords.
			walletPassphrase = []byte(lnwallet.DefaultPrivatePassphrase)
		}
	}

	publicPw := []byte(wallet.InsecurePubPassphrase)

	//	if wallet is locked, temporary unlock it just to check the passphrase
	var walletAux *wallet.Wallet = m.Wallet
	//if wallet_name not passed then try to unlock the default
	walletFile := m.walletFile
	if walletAux == nil || walletAux.Locked() {
		if req.WalletName != "" {
			walletFile = req.WalletName
		}
		loader := wallet.NewLoader(m.netParams, m.walletPath, walletFile, m.noFreelistSync, 0)

		// First, we'll make sure the wallet exists for the specific chain and network.
		walletExists, err := loader.WalletExists()
		if err != nil {
			return nil, err
		}

		if !walletExists {
			return nil, er.New("wallet " + walletFile + " not found")
		}

		// Load the existing wallet in order to proceed with the password change.
		walletAux, err = loader.OpenExistingWallet(publicPw, false)
		if err != nil {
			return nil, err
		}
		log.Info("Wallet " + walletFile + " temporary opened with success")

		// Now that we've opened the wallet, we need to close it before exit
		defer func() {
			if walletAux != m.Wallet {
				_ = loader.UnloadWallet()
				log.Info("Wallet unloaded with success")
			}
		}()
	} else if (req.WalletName != "") && (req.WalletName != m.walletFile) {
		walletFile = req.WalletName
		//wallet is unlocked but not the requested one, so we unlock the wallet_name now
		loader := wallet.NewLoader(m.netParams, m.walletPath, walletFile, m.noFreelistSync, 0)
		// First, we'll make sure the wallet exists for the specific chain and network.
		walletExists, err := loader.WalletExists()
		if err != nil {
			return nil, err
		}

		if !walletExists {
			return nil, er.New("wallet " + walletFile + " not found")
		}

		// Load the existing wallet in order to proceed with the password change.
		walletAux, err = loader.OpenExistingWallet(publicPw, false)
		if err != nil {
			return nil, err
		}
		log.Info("Wallet " + walletFile + " temporary opened with success")

		// Now that we've opened the wallet, we need to close it before exit
		defer func() {
			if walletAux != m.Wallet {
				_ = loader.UnloadWallet()
				log.Info("Wallet unloaded with success")
			}
		}()
	}

	//	attempt to check the private passphrases for the wallet.
	err := walletAux.CheckPassphrase(publicPw, walletPassphrase)
	if err != nil {
		if !strings.HasSuffix(err.Message(), "invalid passphrase for master private key") {
			return nil, err
		}

		return &lnrpc.CheckPasswordResponse{
			ValidPassphrase: false,
		}, nil
	}

	return &lnrpc.CheckPasswordResponse{
		ValidPassphrase: true,
	}, nil
}

func (u *MetaService) ForceCrash(ctx context.Context, req *lnrpc.CrashRequest) (*lnrpc.CrashResponse, error) {

	var someVariable *string = nil

	//	dereference o nil pointer to force a core dump
	if len(*someVariable) == 0 {
		return nil, nil
	}

	return &lnrpc.CrashResponse{}, nil
}
