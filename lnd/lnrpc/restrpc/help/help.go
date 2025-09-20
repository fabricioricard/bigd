////////////////////////////////////////////////////////////////////////////////
//	restrpc/help/help.go  -  Apr-26-2022  -  aldebap
//
//	Data structs to build the help for REST APIs
////////////////////////////////////////////////////////////////////////////////

package help

import "github.com/pkt-cash/pktd/lnd/pkthelp"

const (
	URI_prefix = "/api/v1"
)

//	constants for categories and subcategories of commands
const (
	CategoryLightning             = "Lightning"
	SubcategoryChannel            = "Channel"
	SubSubCategoryBackup          = "Backup"
	SubCategoryGraph              = "Graph"
	SubCategoryInvoice            = "Invoice"
	SubCategoryPayment            = "Payment"
	SubCategoryPeer               = "Peer"
	CategoryMeta                  = "Meta"
	CategoryWallet                = "Wallet"
	SubCategoryNetworkStewardVote = "Network Steward Vote"
	SubCategoryTransaction        = "Transaction"
	SubCategoryUnspent            = "Unspent"
	SubSubCategoryLock            = "Lock"
	SubCategoryAddress            = "Address"
	CategoryNeutrino              = "Neutrino"
	CategoryUtil                  = "Util"
	SubCategorySeed               = "Seed"
	CategoryWatchtower            = "Watchtower"
)

//	mapping with the description for categories and subcategories of commands
var (
	CategoryDescription map[string][]string = map[string][]string{
		CategoryLightning:  {"The Lightning Network component of the wallet"},
		SubcategoryChannel: {"Management of lightning channels to direct peers of this pld node"},
		SubSubCategoryBackup: {"Backup and recovery of the state of active Lightning channels",
			"to and from this pld node"},
		SubCategoryGraph:   {"Information about the global known Lightning Network"},
		SubCategoryInvoice: {"Management of invoices which are used to request payment over Lightning"},
		SubCategoryPayment: {"Lightning network payments which have been made, or have been forwarded, through this node"},
		SubCategoryPeer:    {"Connections to other nodes in the Lightning Network"},
		CategoryMeta:       {"API endpoints which are relevant to the entire pld node, not any specific part"},
		CategoryWallet: {"APIs for management of on-chain (non-Lightning) payments,",
			"seed export and recovery, and on-chain transaction detection"},
		SubCategoryNetworkStewardVote: {"Control how this wallet votes on PKT Network Steward"},
		SubCategoryTransaction:        {"Create and manage on-chain transactions with the wallet"},
		SubCategoryUnspent:            {"Detected unspent transactions associated with one of our wallet addresses"},
		SubSubCategoryLock: {"Manipulation of unspent outputs which are 'locked'",
			"and therefore will not be used to source funds for any transaction"},
		SubCategoryAddress: {"Management of individual wallet addresses"},
		CategoryNeutrino:   {"Management of the Neutrino interface which is used to communicate with the p2p nodes in the network"},
		CategoryUtil:       {"Stateless utility functions which do not affect, not query, the node in any way"},
		SubCategorySeed:    {"Manipulation of mnemonic seed phrases which represent wallet keys"},
		CategoryWatchtower: {"Interact with the watchtower client"},
	}
)

//	constants for every pld command
const (
	//	lightning/channel subCategory commands
	CommandOpenChannel      = "OpenChannel"
	CommandCloseChannel     = "CloseChannel"
	CommandAbandonChannel   = "AbandonChannel"
	CommandChannelBalance   = "ChannelBalance"
	CommandPendingChannels  = "PendingChannels"
	CommandListChannels     = "ListChannels"
	CommandClosedChannels   = "ClosedChannels"
	CommandGetNetworkInfo   = "GetNetworkInfo"
	CommandFeeReport        = "FeeReport"
	CommandUpdateChanPolicy = "UpdateChannelPolicy"
	//	lightning/channel/backup subCategory commands
	CommandExportChanBackup  = "ExportChannelBackup"
	CommandVerifyChanBackup  = "VerifyChanBackup"
	CommandRestoreChanBackup = "RestoreChannelBackups"
	//	lightning/graph subCategory commands
	CommandDescribeGraph  = "DescribeGraph"
	CommandGetNodeMetrics = "GetNodeMetrics"
	CommandGetChanInfo    = "GetChanInfo"
	CommandGetNodeInfo    = "GetNodeInfo"
	//	lightning/invoice subCategory commands
	CommandAddInvoice    = "AddInvoice"
	CommandLookupInvoice = "LookupInvoice"
	CommandListInvoices  = "ListInvoices"
	CommandDecodePayreq  = "DecodePayReq"
	//	lightning/payment subCategory command
	CommandSendPayment   = "SendPaymentV2"
	CommandPayInvoice    = "payinvoice"
	CommandSendToRoute   = "SendToRouteV2"
	CommandListPayments  = "ListPayments"
	CommandTrackPayment  = "TrackPaymentV2"
	CommandQueryRoutes   = "QueryRoutes"
	CommandFwdingHistory = "ForwardingHistory"
	CommandQueryMc       = "QueryMissionControl"
	CommandQueryProb     = "QueryProbability"
	CommandResetMc       = "ResetMissionControl"
	CommandBuildRoute    = "BuildRoute"
	//	lightning/peer subCategory command
	CommandConnectPeer    = "ConnectPeer"
	CommandDisconnectPeer = "DisconnectPeer"
	CommandListPeers      = "ListPeers"
	//	meta category command
	CommandDebugLevel = "DebugLevel"
	CommandGetInfo    = "GetInfo2"
	CommandStop       = "StopDaemon"
	CommandVersion    = "GetVersion"
	CommandCrash      = "ForceCrash"
	//	wallet category command
	CommandWalletBalance    = "WalletBalance"
	CommandChangePassphrase = "ChangePassword"
	CommandCheckPassphrase  = "CheckPassword"
	CommandCreateWallet     = "InitWallet"
	CommandGetSecret        = "GetSecret"
	CommandGetWalletSeed    = "GetWalletSeed"
	CommandUnlockWallet     = "UnlockWallet"
	//	wallet/networkstewardvote subCategory command
	CommandGetNetworkStewardVote = "GetNetworkStewardVote"
	CommandSetNetworkStewardVote = "SetNetworkStewardVote"
	//	wallet/transaction subCategory command
	CommandGetTransaction    = "GetTransaction"
	CommandCreateTransaction = "CreateTransaction"
	CommandQueryTransactions = "GetTransactions"
	CommandSendCoins         = "SendCoins"
	CommandSendFrom          = "SendFrom"
	CommandSendMany          = "SendMany"
	//	wallet/unspent subCategory command
	CommandListUnspent = "ListUnspent"
	CommandResync      = "ReSync"
	CommandStopResync  = "StopReSync"
	//	wallet/unspent/lock subCategory command
	CommandListLockUnspent = "ListLockUnspent"
	CommandLockUnspent     = "LockUnspent"
	//	wallet/address subCategory command
	CommandGetAddressBalances = "GetAddressBalances"
	CommandNewAddress         = "GetNewAddress"
	CommandDumpPrivkey        = "DumpPrivKey"
	CommandImportPrivkey      = "ImportPrivKey"
	CommandSignMessage        = "SignMessage"
	//	neutrino category command
	CommandBcastTransaction = "BcastTransaction"
	CommandEstimateFee      = "EstimateFee"
	// util category command
	CommandDecodeRawTransaction = "DecodeRawTransaction"
	//	util/seed subCategory command
	CommandChangeSeedPassphrase = "ChangeSeedPassphrase"
	CommandGenSeed              = "GenSeed"
	//	wtclient/tower subCategory command
	CommandCreateWatchTower = "AddTower"
	CommandRemoveTower      = "RemoveTower"
	CommandListTowers       = "ListTowers"
	CommandGetTowerInfo     = "GetTowerInfo"
	CommandGetTowerStats    = "Stats"
	CommandGetTowerPolicy   = "Policy"
)

type CommandInfo struct {
	Command     string
	Category    string
	Description string
	Path        string
	AllowGet    bool
	HelpInfo    func() pkthelp.Method
}

//	mapping with the category, description and path for every command
var (
	CommandInfoData []CommandInfo = []CommandInfo{
		//	lightning/channel subCategory commands
		{Command: CommandOpenChannel, Path: "/lightning/channel/open"},
		{Command: CommandCloseChannel, Path: "/lightning/channel/close"},
		{Command: CommandAbandonChannel, Path: "/lightning/channel/abandon"},
		{Command: CommandChannelBalance, Path: "/lightning/channel/balance"},
		{Command: CommandPendingChannels, Path: "/lightning/channel/pending"},
		{Command: CommandListChannels, Path: "/lightning/channel", AllowGet: true},
		{Command: CommandClosedChannels, Path: "/lightning/channel/closed", AllowGet: true},
		{Command: CommandGetNetworkInfo, Path: "/lightning/channel/networkinfo"},
		{Command: CommandFeeReport, Path: "/lightning/channel/feereport"},
		{Command: CommandUpdateChanPolicy, Path: "/lightning/channel/policy"},
		//	lightning/channel/backup subCategory commands
		{Command: CommandExportChanBackup, Path: "/lightning/channel/backup/export"},
		{Command: CommandVerifyChanBackup, Path: "/lightning/channel/backup/verify"},
		{Command: CommandRestoreChanBackup, Path: "/lightning/channel/backup/restore"},
		//	lightning/graph subCategory commands
		{Command: CommandDescribeGraph, Path: "/lightning/graph", AllowGet: true},
		{Command: CommandGetNodeMetrics, Path: "/lightning/graph/nodemetrics", AllowGet: true},
		{Command: CommandGetChanInfo, Path: "/lightning/graph/channel"},
		{Command: CommandGetNodeInfo, Path: "/lightning/graph/nodeinfo"},
		//	lightning/invoice subCategory commands
		{Command: CommandAddInvoice, Path: "/lightning/invoice/create"},
		{Command: CommandLookupInvoice, Path: "/lightning/invoice/lookup"},
		{Command: CommandListInvoices, Path: "/lightning/invoice", AllowGet: true},
		{Command: CommandDecodePayreq, Path: "/lightning/invoice/decodepayreq"},
		//	lightning/payment subCategory command
		{Command: CommandSendPayment, Path: "/lightning/payment/send"},
		//	TODO: need to understand what command does the PayInvoice functionality
		{Command: CommandPayInvoice, Description: "Pay an invoice over lightning", Path: "/lightning/payment/payinvoice"},
		{Command: CommandSendToRoute, Path: "/lightning/payment/sendtoroute"},
		{Command: CommandListPayments, Path: "/lightning/payment", AllowGet: true},
		{Command: CommandTrackPayment, Path: "/lightning/payment/track"},
		{Command: CommandQueryRoutes, Path: "/lightning/payment/queryroutes"},
		{Command: CommandFwdingHistory, Path: "/lightning/payment/fwdinghistory"},
		{Command: CommandQueryMc, Path: "/lightning/payment/querymc"},
		{Command: CommandQueryProb, Path: "/lightning/payment/queryprob"},
		{Command: CommandResetMc, Path: "/lightning/payment/resetmc"},
		{Command: CommandBuildRoute, Path: "/lightning/payment/buildroute"},
		//	lightning/peer subCategory command
		{Command: CommandConnectPeer, Path: "/lightning/peer/connect"},
		{Command: CommandDisconnectPeer, Path: "/lightning/peer/disconnect"},
		{Command: CommandListPeers, Path: "/lightning/peer", AllowGet: true},
		//	meta category command
		{Command: CommandDebugLevel, Path: "/meta/debuglevel"},
		{Command: CommandGetInfo, Path: "/meta/getinfo"},
		{Command: CommandStop, Path: "/meta/stop"},
		{Command: CommandVersion, Path: "/meta/version"},
		{Command: CommandCrash, Path: "/meta/crash"},
		//	wallet category command
		{Command: CommandWalletBalance, Path: "/wallet/balance"},
		{Command: CommandChangePassphrase, Path: "/wallet/changepassphrase"},
		{Command: CommandCheckPassphrase, Path: "/wallet/checkpassphrase"},
		{Command: CommandCreateWallet, Path: "/wallet/create"},
		{Command: CommandGetSecret, Path: "/wallet/getsecret"},
		{Command: CommandGetWalletSeed, Path: "/wallet/seed"},
		{Command: CommandUnlockWallet, Path: "/wallet/unlock"},
		//	wallet/networkstewardvote subCategory command
		{Command: CommandGetNetworkStewardVote, Path: "/wallet/networkstewardvote"},
		{Command: CommandSetNetworkStewardVote, Path: "/wallet/networkstewardvote/set"},
		//	wallet/transaction subCategory command
		{Command: CommandGetTransaction, Path: "/wallet/transaction"},
		{Command: CommandCreateTransaction, Path: "/wallet/transaction/create"},
		{Command: CommandQueryTransactions, Path: "/wallet/transaction/query", AllowGet: true},
		{Command: CommandSendCoins, Path: "/wallet/transaction/sendcoins"},
		{Command: CommandSendFrom, Path: "/wallet/transaction/sendfrom"},
		{Command: CommandSendMany, Path: "/wallet/transaction/sendmany"},
		//	wallet/unspent subCategory command
		{Command: CommandListUnspent, Path: "/wallet/unspent", AllowGet: true},
		{Command: CommandResync, Path: "/wallet/unspent/resync"},
		{Command: CommandStopResync, Path: "/wallet/unspent/stopresync"},
		//	wallet/unspent/lock subCategory command
		{Command: CommandListLockUnspent, Path: "/wallet/unspent/lock"},
		{Command: CommandLockUnspent, Path: "/wallet/unspent/lock/create"},
		// TODO: need to create a /wallet/unspent/lock/delete command
		//	wallet/address subCategory command
		{Command: CommandGetAddressBalances, Path: "/wallet/address/balances", AllowGet: true},
		{Command: CommandNewAddress, Path: "/wallet/address/create"},
		{Command: CommandDumpPrivkey, Path: "/wallet/address/dumpprivkey"},
		{Command: CommandImportPrivkey, Path: "/wallet/address/import"},
		{Command: CommandSignMessage, Path: "/wallet/address/signmessage"},
		//	neutrino category command
		{Command: CommandBcastTransaction, Path: "/neutrino/bcasttransaction"},
		{Command: CommandEstimateFee, Path: "/neutrino/estimatefee"},
		//	util category command
		{Command: CommandDecodeRawTransaction, Path: "/util/transaction/decode"},
		//	util/seed subCategory command
		{Command: CommandChangeSeedPassphrase, Path: "/util/seed/changepassphrase"},
		{Command: CommandGenSeed, Path: "/util/seed/create"},
		//	wtclient/tower subCategory command
		{Command: CommandCreateWatchTower, Path: "/wtclient/tower/create"},
		{Command: CommandRemoveTower, Path: "/wtclient/tower/remove"},
		{Command: CommandListTowers, Path: "/wtclient/tower", AllowGet: true},
		{Command: CommandGetTowerInfo, Path: "/wtclient/tower/getinfo"},
		{Command: CommandGetTowerStats, Path: "/wtclient/tower/stats"},
		{Command: CommandGetTowerPolicy, Path: "/wtclient/tower/policy"},
	}
)

//	init command info data by getting meta fields from help based on command name
func init() {

	var commandHelpFunctions []func() pkthelp.Method = []func() pkthelp.Method{
		pkthelp.Lightning_OpenChannel,
		pkthelp.Lightning_CloseChannel,
		pkthelp.Lightning_AbandonChannel,
		pkthelp.Lightning_ChannelBalance,
		pkthelp.Lightning_PendingChannels,
		pkthelp.Lightning_ListChannels,
		pkthelp.Lightning_ClosedChannels,
		pkthelp.Lightning_GetNetworkInfo,
		pkthelp.Lightning_FeeReport,
		pkthelp.Lightning_UpdateChannelPolicy,

		pkthelp.Lightning_ExportChannelBackup,
		pkthelp.Lightning_VerifyChanBackup,
		pkthelp.Lightning_RestoreChannelBackups,

		pkthelp.Lightning_DescribeGraph,
		pkthelp.Lightning_GetNodeMetrics,
		pkthelp.Lightning_GetChanInfo,
		pkthelp.Lightning_GetNodeInfo,

		pkthelp.Lightning_AddInvoice,
		pkthelp.Lightning_LookupInvoice,
		pkthelp.Lightning_ListInvoices,
		pkthelp.Lightning_DecodePayReq,

		pkthelp.Router_SendPaymentV2,
		//	PayInvoice
		pkthelp.Router_SendToRouteV2,
		pkthelp.Lightning_ListPayments,
		pkthelp.Router_TrackPaymentV2,
		pkthelp.Lightning_QueryRoutes,
		pkthelp.Lightning_ForwardingHistory,
		pkthelp.Router_QueryMissionControl,
		pkthelp.Router_QueryProbability,
		pkthelp.Router_ResetMissionControl,
		pkthelp.Router_BuildRoute,

		pkthelp.Lightning_ConnectPeer,
		pkthelp.Lightning_DisconnectPeer,
		pkthelp.Lightning_ListPeers,

		pkthelp.Lightning_DebugLevel,
		pkthelp.MetaService_GetInfo2,
		pkthelp.Lightning_StopDaemon,
		pkthelp.Versioner_GetVersion,
		pkthelp.MetaService_ForceCrash,

		pkthelp.Lightning_WalletBalance,
		pkthelp.MetaService_ChangePassword,
		pkthelp.MetaService_CheckPassword,
		pkthelp.WalletUnlocker_InitWallet,
		pkthelp.Lightning_GetSecret,
		pkthelp.Lightning_GetWalletSeed,
		pkthelp.WalletUnlocker_UnlockWallet,

		pkthelp.Lightning_GetNetworkStewardVote,
		pkthelp.Lightning_SetNetworkStewardVote,

		pkthelp.Lightning_GetTransaction,
		pkthelp.Lightning_CreateTransaction,
		pkthelp.Lightning_GetTransactions,
		pkthelp.Lightning_SendCoins,
		pkthelp.Lightning_SendFrom,
		pkthelp.Lightning_SendMany,

		pkthelp.WalletKit_ListUnspent,
		pkthelp.Lightning_ReSync,
		pkthelp.Lightning_StopReSync,

		pkthelp.Lightning_ListLockUnspent,
		pkthelp.Lightning_LockUnspent,

		pkthelp.Lightning_GetAddressBalances,
		pkthelp.Lightning_GetNewAddress,
		pkthelp.Lightning_DumpPrivKey,
		pkthelp.Lightning_ImportPrivKey,
		pkthelp.Signer_SignMessage,

		pkthelp.Lightning_BcastTransaction,
		pkthelp.Lightning_EstimateFee,

		pkthelp.Lightning_DecodeRawTransaction,
		pkthelp.Lightning_ChangeSeedPassphrase,
		pkthelp.WalletUnlocker_GenSeed,

		pkthelp.WatchtowerClient_AddTower,
		pkthelp.WatchtowerClient_RemoveTower,
		pkthelp.WatchtowerClient_ListTowers,
		pkthelp.WatchtowerClient_GetTowerInfo,
		pkthelp.WatchtowerClient_Stats,
		pkthelp.WatchtowerClient_Policy,
	}

	//	for each command help fuction set meta data to commandInfoData
	for _, helpFunction := range commandHelpFunctions {

		var commandMethod = helpFunction()

		for i, commandInfo := range CommandInfoData {
			if commandMethod.Name == commandInfo.Command {
				CommandInfoData[i].Category = commandMethod.Category
				CommandInfoData[i].Description = commandMethod.ShortDescription
				CommandInfoData[i].HelpInfo = helpFunction
			}
		}
	}
}

//	the category help in REST master
func RESTCategory_help(category string, subCategory []*RestCommandCategory) *RestCommandCategory {
	restCommandCategory := &RestCommandCategory{
		Name:        category,
		Description: CategoryDescription[category],
		Endpoints:   make([]*RestEndpoint, 0, 10),
		Subcategory: make([]*RestCommandCategory, 0, 10),
	}

	//	add all endpoints for the category
	for _, commandInfo := range CommandInfoData {

		if commandInfo.Category == category {
			restCommandCategory.Endpoints = append(restCommandCategory.Endpoints, &RestEndpoint{
				URI:              URI_prefix + commandInfo.Path,
				ShortDescription: commandInfo.Description,
			})
		}
	}

	//	add sub categories
	restCommandCategory.Subcategory = subCategory

	return restCommandCategory
}

//	return the REST master help messsage
func RESTMaster_help() *RestMasterHelpResponse {
	masterHelpResp := &RestMasterHelpResponse{
		Name: "pld - Lightning Network Daemon REST interface (pld)",
		Description: []string{
			"General information about PLD",
		},
		Category: []*RestCommandCategory{
			RESTCategory_help(CategoryLightning, []*RestCommandCategory{
				RESTCategory_help(SubcategoryChannel, []*RestCommandCategory{
					RESTCategory_help(SubSubCategoryBackup, []*RestCommandCategory{}),
				}),
				RESTCategory_help(SubCategoryGraph, []*RestCommandCategory{}),
				RESTCategory_help(SubCategoryInvoice, []*RestCommandCategory{}),
				RESTCategory_help(SubCategoryPayment, []*RestCommandCategory{}),
				RESTCategory_help(SubCategoryPeer, []*RestCommandCategory{}),
			}),
			RESTCategory_help(CategoryMeta, []*RestCommandCategory{}),
			RESTCategory_help(CategoryWallet, []*RestCommandCategory{
				RESTCategory_help(SubCategoryNetworkStewardVote, []*RestCommandCategory{}),
				RESTCategory_help(SubCategoryTransaction, []*RestCommandCategory{}),
				RESTCategory_help(SubCategoryUnspent, []*RestCommandCategory{
					RESTCategory_help(SubSubCategoryLock, []*RestCommandCategory{}),
				}),
				RESTCategory_help(SubCategoryAddress, []*RestCommandCategory{}),
			}),
			RESTCategory_help(CategoryNeutrino, []*RestCommandCategory{}),
			RESTCategory_help(CategoryUtil, []*RestCommandCategory{
				RESTCategory_help(SubCategorySeed, []*RestCommandCategory{}),
			}),
			RESTCategory_help(CategoryWatchtower, []*RestCommandCategory{}),
		},
	}

	return masterHelpResp
}
