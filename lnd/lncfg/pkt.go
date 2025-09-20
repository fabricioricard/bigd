package lncfg

// Bitcoind holds the configuration options for the daemon's connection to
// bitcoind.
type Pkt struct {
	Dir       string `long:"dir" description:"The base directory that contains the node's data, logs, configuration file, etc."`
	WalletDir string `long:"walletDir" description:"the sub directory where the wallet.db file resides"`
	RPCHost   string `long:"rpchost" description:"The daemon's rpc listening address. If a port is omitted, then the default port for the selected chain parameters will be used."`
	RPCUser   string `long:"rpcuser" description:"Username for RPC connections"`
	RPCPass   string `long:"rpcpass" default-mask:"-" description:"Password for RPC connections"`
}
