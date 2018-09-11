package utils

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/kowala-tech/kcoin/client/consensus/konsensus"
	"github.com/kowala-tech/kcoin/client/stats"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/accounts/keystore"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/fdlimit"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/state"
	"github.com/kowala-tech/kcoin/client/core/vm"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/kcoindb"
	"github.com/kowala-tech/kcoin/client/knode"
	"github.com/kowala-tech/kcoin/client/knode/downloader"
	"github.com/kowala-tech/kcoin/client/knode/gasprice"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/metrics"
	"github.com/kowala-tech/kcoin/client/metrics/influxdb"
	"github.com/kowala-tech/kcoin/client/node"
	"github.com/kowala-tech/kcoin/client/p2p"
	"github.com/kowala-tech/kcoin/client/p2p/discover"
	"github.com/kowala-tech/kcoin/client/p2p/discv5"
	"github.com/kowala-tech/kcoin/client/p2p/nat"
	"github.com/kowala-tech/kcoin/client/p2p/netutil"
	"github.com/kowala-tech/kcoin/client/params"

	"gopkg.in/urfave/cli.v1"
)

var (
	CommandHelpTemplate = `{{.cmd.Name}}{{if .cmd.Subcommands}} command{{end}}{{if .cmd.Flags}} [command options]{{end}} [arguments...]
{{if .cmd.Description}}{{.cmd.Description}}
{{end}}{{if .cmd.Subcommands}}
SUBCOMMANDS:
	{{range .cmd.Subcommands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
	{{end}}{{end}}{{if .categorizedFlags}}
{{range $idx, $categorized := .categorizedFlags}}{{$categorized.Name}} OPTIONS:
{{range $categorized.Flags}}{{"\t"}}{{.}}
{{end}}
{{end}}{{end}}`
)

func init() {
	cli.AppHelpTemplate = `{{.Name}} {{if .Flags}}[global options] {{end}}command{{if .Flags}} [command options]{{end}} [arguments...]

VERSION:
   {{.Version}}

COMMANDS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`

	cli.CommandHelpTemplate = CommandHelpTemplate
}

// NewApp creates an app with sane defaults.
func NewApp(gitCommit, usage string) *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Author = ""
	//app.Authors = nil
	app.Email = ""
	app.Version = params.Version
	if len(gitCommit) >= 8 {
		app.Version += "-" + gitCommit[:8]
	}
	app.Usage = usage
	return app
}

// These are all the command line flags we support.
// If you add to this list, please remember to include the
// flag in the appropriate command definition.
//
// The flags are defined here so their names and help texts
// are the same for all commands.

var (
	// General settings
	DataDirFlag = DirectoryFlag{
		Name:  "datadir",
		Usage: "Data directory for the databases and keystore",
		Value: DirectoryString{node.DefaultDataDir()},
	}
	KeyStoreDirFlag = DirectoryFlag{
		Name:  "keystore",
		Usage: "Directory for the keystore (default = inside the datadir)",
	}
	NoUSBFlag = cli.BoolFlag{
		Name:  "nousb",
		Usage: "Disables monitoring for and managing USB hardware wallets",
	}
	NetworkIdFlag = cli.Uint64Flag{
		Name:  "networkid",
		Usage: "Network identifier (integer, 1=MainNet, 2=TestNet)",
		Value: knode.DefaultConfig.NetworkId,
	}
	TestnetFlag = cli.BoolFlag{
		Name:  "testnet",
		Usage: "Zygote network: pre-configured proof-of-stake test network",
	}
	DevModeFlag = cli.BoolFlag{
		Name:  "dev",
		Usage: "Developer mode: pre-configured private test network",
	}
	IdentityFlag = cli.StringFlag{
		Name:  "identity",
		Usage: "Custom node name",
	}
	CurrencyFlag = cli.StringFlag{
		Name:  "currency",
		Usage: "Currency to use with the client",
		Value: knode.DefaultConfig.Currency,
	}
	DocRootFlag = DirectoryFlag{
		Name:  "docroot",
		Usage: "Document Root for HTTPClient file scheme",
		Value: DirectoryString{homeDir()},
	}
	VersionRepository = cli.StringFlag{
		Name:  "repository",
		Usage: "Version repository url",
		Value: "https://s3.amazonaws.com/releases.kowala.tech",
	}
	SelfUpdateEnabledFlag = cli.BoolFlag{
		Name:  "selfupdate",
		Usage: "Auto-updated binary to latest version",
	}
	FastSyncFlag = cli.BoolFlag{
		Name:  "fast",
		Usage: "Enable fast syncing through state downloads (replaced by --syncmode)",
	}
	LightModeFlag = cli.BoolFlag{
		Name:  "light",
		Usage: "Enable light client mode (replaced by --syncmode)",
	}
	defaultSyncMode = knode.DefaultConfig.SyncMode
	SyncModeFlag    = TextMarshalerFlag{
		Name:  "syncmode",
		Usage: `Blockchain sync mode ("fast", "full", or "light")`,
		Value: &defaultSyncMode,
	}
	GCModeFlag = cli.StringFlag{
		Name:  "gcmode",
		Usage: `Blockchain garbage collection mode ("full", "archive")`,
		Value: "full",
	}
	LightServFlag = cli.IntFlag{
		Name:  "lightserv",
		Usage: "Maximum percentage of time allowed for serving LES requests (0-90)",
		Value: 0,
	}
	LightPeersFlag = cli.IntFlag{
		Name:  "lightpeers",
		Usage: "Maximum number of LES client peers",
		Value: 20,
	}
	LightKDFFlag = cli.BoolFlag{
		Name:  "lightkdf",
		Usage: "Reduce key-derivation RAM & CPU usage at some expense of KDF strength",
	}
	// Transaction pool settings
	TxPoolNoLocalsFlag = cli.BoolFlag{
		Name:  "txpool.nolocals",
		Usage: "Disables price exemptions for locally submitted transactions",
	}
	TxPoolJournalFlag = cli.StringFlag{
		Name:  "txpool.journal",
		Usage: "Disk journal for local transaction to survive node restarts",
		Value: core.DefaultTxPoolConfig.Journal,
	}
	TxPoolRejournalFlag = cli.DurationFlag{
		Name:  "txpool.rejournal",
		Usage: "Time interval to regenerate the local transaction journal",
		Value: core.DefaultTxPoolConfig.Rejournal,
	}
	TxPoolPriceLimitFlag = cli.Uint64Flag{
		Name:  "txpool.pricelimit",
		Usage: "Minimum gas price limit to enforce for acceptance into the pool",
		Value: knode.DefaultConfig.TxPool.PriceLimit,
	}
	TxPoolPriceBumpFlag = cli.Uint64Flag{
		Name:  "txpool.pricebump",
		Usage: "Price bump percentage to replace an already existing transaction",
		Value: knode.DefaultConfig.TxPool.PriceBump,
	}
	TxPoolAccountSlotsFlag = cli.Uint64Flag{
		Name:  "txpool.accountslots",
		Usage: "Minimum number of executable transaction slots guaranteed per account",
		Value: knode.DefaultConfig.TxPool.AccountSlots,
	}
	TxPoolGlobalSlotsFlag = cli.Uint64Flag{
		Name:  "txpool.globalslots",
		Usage: "Maximum number of executable transaction slots for all accounts",
		Value: knode.DefaultConfig.TxPool.GlobalSlots,
	}
	TxPoolAccountQueueFlag = cli.Uint64Flag{
		Name:  "txpool.accountqueue",
		Usage: "Maximum number of non-executable transaction slots permitted per account",
		Value: knode.DefaultConfig.TxPool.AccountQueue,
	}
	TxPoolGlobalQueueFlag = cli.Uint64Flag{
		Name:  "txpool.globalqueue",
		Usage: "Maximum number of non-executable transaction slots for all accounts",
		Value: knode.DefaultConfig.TxPool.GlobalQueue,
	}
	TxPoolLifetimeFlag = cli.DurationFlag{
		Name:  "txpool.lifetime",
		Usage: "Maximum amount of time non-executable transaction are queued",
		Value: knode.DefaultConfig.TxPool.Lifetime,
	}
	// Performance tuning settings
	CacheFlag = cli.IntFlag{
		Name:  "cache",
		Usage: "Megabytes of memory allocated to internal caching (min 16MB / database forced)",
		Value: 128,
	}
	CacheDatabaseFlag = cli.IntFlag{
		Name:  "cache.database",
		Usage: "Percentage of cache memory allowance to use for database io",
		Value: 75,
	}
	CacheGCFlag = cli.IntFlag{
		Name:  "cache.gc",
		Usage: "Percentage of cache memory allowance to use for trie pruning",
		Value: 25,
	}
	TrieCacheGenFlag = cli.IntFlag{
		Name:  "trie-cache-gens",
		Usage: "Number of trie node generations to keep in memory",
		Value: int(state.MaxTrieCacheGen),
	}
	// Consensus Validator settings
	ValidationEnabledFlag = cli.BoolFlag{
		Name:  "validate",
		Usage: "Enable consensus validation",
	}

	ValidatorDepositFlag = cli.Uint64Flag{
		Name:  "deposit",
		Usage: "Deposit at stake",
		// @TODO (rgeraldes) - default could be set to the minimum required
	}

	TargetGasLimitFlag = cli.Uint64Flag{
		Name:  "targetgaslimit",
		Usage: "Target gas limit sets the artificial target gas floor for the blocks to mine",
		Value: params.GenesisGasLimit,
	}
	CoinbaseFlag = cli.StringFlag{
		Name:  "coinbase",
		Usage: "Public address for block validation rewards (default = first account created)",
		Value: "0",
	}

	// @TODO(rgeraldes) - review
	GasPriceFlag = BigFlag{
		Name:  "gasprice",
		Usage: "Minimal gas price to accept for mining a transactions",
		Value: knode.DefaultConfig.GasPrice,
	}
	ExtraDataFlag = cli.StringFlag{
		Name:  "extradata",
		Usage: "Block extra data set by the consensus validator (default = client version)",
	}
	// Account settings
	UnlockedAccountFlag = cli.StringFlag{
		Name:  "unlock",
		Usage: "Comma separated list of accounts to unlock",
		Value: "",
	}
	PasswordFileFlag = cli.StringFlag{
		Name:  "password",
		Usage: "Password file to use for non-inteactive password input",
		Value: "",
	}

	VMEnableDebugFlag = cli.BoolFlag{
		Name:  "vmdebug",
		Usage: "Record information useful for VM and contract debugging",
	}
	// Logging and debug settings
	ShipLogzioFlag = cli.StringFlag{
		Name:  "logzioapi",
		Usage: "Logzio API key for shipping logs",
		Value: "",
	}
	VerbosityFlag = cli.IntFlag{
		Name:  "verbosity",
		Usage: "sets the verbosity level",
	}
	KowalaStatsURLFlag = cli.StringFlag{
		Name:  "stats",
		Usage: "Reporting URL of a stats service (nodename:secret@host:port)",
	}
	NoCompactionFlag = cli.BoolFlag{
		Name:  "nocompaction",
		Usage: "Disables db compaction after import",
	}
	// RPC settings
	RPCEnabledFlag = cli.BoolFlag{
		Name:  "rpc",
		Usage: "Enable the HTTP-RPC server",
	}
	RPCListenAddrFlag = cli.StringFlag{
		Name:  "rpcaddr",
		Usage: "HTTP-RPC server listening interface",
		Value: node.DefaultHTTPHost,
	}
	RPCPortFlag = cli.IntFlag{
		Name:  "rpcport",
		Usage: "HTTP-RPC server listening port",
		Value: node.DefaultHTTPPort,
	}
	RPCCORSDomainFlag = cli.StringFlag{
		Name:  "rpccorsdomain",
		Usage: "Comma separated list of domains from which to accept cross origin requests (browser enforced)",
		Value: "",
	}
	RPCVirtualHostsFlag = cli.StringFlag{
		Name:  "rpcvhosts",
		Usage: "Comma separated list of virtual hostnames from which to accept requests (server enforced). Accepts '*' wildcard.",
		Value: strings.Join(node.DefaultConfig.HTTPVirtualHosts, ","),
	}
	RPCApiFlag = cli.StringFlag{
		Name:  "rpcapi",
		Usage: "API's offered over the HTTP-RPC interface",
		Value: "",
	}
	IPCDisabledFlag = cli.BoolFlag{
		Name:  "ipcdisable",
		Usage: "Disable the IPC-RPC server",
	}
	IPCPathFlag = DirectoryFlag{
		Name:  "ipcpath",
		Usage: "Filename for IPC socket/pipe within the datadir (explicit paths escape it)",
	}
	WSEnabledFlag = cli.BoolFlag{
		Name:  "ws",
		Usage: "Enable the WS-RPC server",
	}
	WSListenAddrFlag = cli.StringFlag{
		Name:  "wsaddr",
		Usage: "WS-RPC server listening interface",
		Value: node.DefaultWSHost,
	}
	WSPortFlag = cli.IntFlag{
		Name:  "wsport",
		Usage: "WS-RPC server listening port",
		Value: node.DefaultWSPort,
	}
	WSApiFlag = cli.StringFlag{
		Name:  "wsapi",
		Usage: "API's offered over the WS-RPC interface",
		Value: "",
	}
	WSAllowedOriginsFlag = cli.StringFlag{
		Name:  "wsorigins",
		Usage: "Origins from which to accept websockets requests",
		Value: "",
	}
	ExecFlag = cli.StringFlag{
		Name:  "exec",
		Usage: "Execute JavaScript statement",
	}
	PreloadJSFlag = cli.StringFlag{
		Name:  "preload",
		Usage: "Comma separated list of JavaScript files to preload into the console",
	}

	// Network Settings
	MaxPeersFlag = cli.IntFlag{
		Name:  "maxpeers",
		Usage: "Maximum number of network peers (network disabled if set to 0)",
		Value: 100,
	}
	MaxPendingPeersFlag = cli.IntFlag{
		Name:  "maxpendpeers",
		Usage: "Maximum number of pending connection attempts (defaults used if set to 0)",
		Value: 0,
	}
	ListenPortFlag = cli.IntFlag{
		Name:  "port",
		Usage: "Network listening port",
		Value: 22334,
	}
	BootnodesFlag = cli.StringFlag{
		Name:  "bootnodes",
		Usage: "Comma separated enode URLs for P2P discovery bootstrap (set v4+v5 instead for light servers)",
		Value: "",
	}
	BootnodesV4Flag = cli.StringFlag{
		Name:  "bootnodesv4",
		Usage: "Comma separated enode URLs for P2P v4 discovery bootstrap (light server, full nodes)",
		Value: "",
	}
	BootnodesV5Flag = cli.StringFlag{
		Name:  "bootnodesv5",
		Usage: "Comma separated enode URLs for P2P v5 discovery bootstrap (light server, light nodes)",
		Value: "",
	}
	NodeKeyFileFlag = cli.StringFlag{
		Name:  "nodekey",
		Usage: "P2P node key file",
	}
	NodeKeyHexFlag = cli.StringFlag{
		Name:  "nodekeyhex",
		Usage: "P2P node key as hex (for testing)",
	}
	NATFlag = cli.StringFlag{
		Name:  "nat",
		Usage: "NAT port mapping mechanism (any|none|upnp|pmp|extip:<IP>)",
		Value: "any",
	}
	NoDiscoverFlag = cli.BoolFlag{
		Name:  "nodiscover",
		Usage: "Disables the peer discovery mechanism (manual peer addition)",
	}
	NetrestrictFlag = cli.StringFlag{
		Name:  "netrestrict",
		Usage: "Restricts network communication to the given IP networks (CIDR masks)",
	}

	// ATM the url is left to the user and deployment to
	JSpathFlag = cli.StringFlag{
		Name:  "jspath",
		Usage: "JavaScript root path for `loadScript`",
		Value: ".",
	}

	// Gas price oracle settings
	GpoBlocksFlag = cli.IntFlag{
		Name:  "gpoblocks",
		Usage: "Number of recent blocks to check for gas prices",
		Value: knode.DefaultConfig.GPO.Blocks,
	}
	GpoPercentileFlag = cli.IntFlag{
		Name:  "gpopercentile",
		Usage: "Suggested gas price is the given percentile of a set of recent transaction gas prices",
		Value: knode.DefaultConfig.GPO.Percentile,
	}

	MetricsEnabledFlag = cli.BoolFlag{
		Name:  metrics.MetricsEnabledFlag,
		Usage: "Enable metrics collection and reporting",
	}
	MetricsEnableInfluxDBFlag = cli.BoolFlag{
		Name:  "metrics.influxdb",
		Usage: "Enable metrics export/push to an external InfluxDB database",
	}
	MetricsInfluxDBEndpointFlag = cli.StringFlag{
		Name:  "metrics.influxdb.endpoint",
		Usage: "InfluxDB API endpoint to report metrics to",
		Value: "http://localhost:8086",
	}
	MetricsInfluxDBDatabaseFlag = cli.StringFlag{
		Name:  "metrics.influxdb.database",
		Usage: "InfluxDB database name to push reported metrics to",
		Value: "geth",
	}
	MetricsInfluxDBUsernameFlag = cli.StringFlag{
		Name:  "metrics.influxdb.username",
		Usage: "Username to authorize access to the database",
		Value: "test",
	}
	MetricsInfluxDBPasswordFlag = cli.StringFlag{
		Name:  "metrics.influxdb.password",
		Usage: "Password to authorize access to the database",
		Value: "test",
	}
	// The `host` tag is part of every measurement sent to InfluxDB. Queries on tags are faster in InfluxDB.
	// It is used so that we can group all nodes and average a measurement across all of them, but also so
	// that we can select a specific node and inspect its measurements.
	// https://docs.influxdata.com/influxdb/v1.4/concepts/key_concepts/#tag-key
	MetricsInfluxDBHostTagFlag = cli.StringFlag{
		Name:  "metrics.influxdb.host.tag",
		Usage: "InfluxDB `host` tag attached to all measurements",
		Value: "localhost",
	}

	MetricsPrometheusAddressFlag = cli.StringFlag{
		Name:  metrics.MetricsPrometheusAddressFlag,
		Usage: "Set the address for Prometheus reporting",
		Value: ":8080",
	}
	MetricsPrometheusSubsystemFlag = cli.StringFlag{
		Name:  metrics.MetricsPrometheusSubsystemFlag,
		Usage: "Set the subsystem name for Prometheus reporting",
		Value: "node",
	}
)

// MakeDataDir retrieves the currently requested data directory, terminating
// if none (or the empty string) is specified. If the node is starting a testnet,
// the a subdirectory of the specified datadir will be used.
func MakeDataDir(ctx *cli.Context) string {
	if path := ctx.GlobalString(DataDirFlag.Name); path != "" {
		return path
	}
	Fatalf("Cannot determine default data directory, please set manually (--datadir)")
	return ""
}

// setNodeKey creates a node key from set command line flags, either loading it
// from a file or as a specified hex value. If neither flags were provided, this
// method returns nil and an emphemeral key is to be generated.
func setNodeKey(ctx *cli.Context, cfg *p2p.Config) {
	var (
		hex  = ctx.GlobalString(NodeKeyHexFlag.Name)
		file = ctx.GlobalString(NodeKeyFileFlag.Name)
		key  *ecdsa.PrivateKey
		err  error
	)
	switch {
	case file != "" && hex != "":
		Fatalf("Options %q and %q are mutually exclusive", NodeKeyFileFlag.Name, NodeKeyHexFlag.Name)
	case file != "":
		if key, err = crypto.LoadECDSA(file); err != nil {
			Fatalf("Option %q: %v", NodeKeyFileFlag.Name, err)
		}
		cfg.PrivateKey = key
	case hex != "":
		if key, err = crypto.HexToECDSA(hex); err != nil {
			Fatalf("Option %q: %v", NodeKeyHexFlag.Name, err)
		}
		cfg.PrivateKey = key
	}
}

// setNodeUserIdent creates the user identifier from CLI flags.
func setNodeUserIdent(ctx *cli.Context, cfg *node.Config) {
	if identity := ctx.GlobalString(IdentityFlag.Name); len(identity) > 0 {
		cfg.UserIdent = identity
	}
}

// setBootstrapNodes creates a list of bootstrap nodes from the command line
// flags, reverting to pre-configured ones if none have been specified.
func setBootstrapNodes(ctx *cli.Context, cfg *p2p.Config) {
	urls := params.MainnetBootnodes
	switch {
	case ctx.GlobalIsSet(BootnodesFlag.Name) || ctx.GlobalIsSet(BootnodesV4Flag.Name):
		if ctx.GlobalIsSet(BootnodesV4Flag.Name) {
			urls = strings.Split(ctx.GlobalString(BootnodesV4Flag.Name), ",")
		} else {
			urls = strings.Split(ctx.GlobalString(BootnodesFlag.Name), ",")
		}
	case ctx.GlobalBool(TestnetFlag.Name):
		urls = params.TestnetBootnodes
	case ctx.GlobalBool(DevModeFlag.Name):
		urls = params.DevnetBootnodes
	}

	cfg.BootstrapNodes = make([]*discover.Node, 0, len(urls))
	for _, url := range urls {
		node, err := discover.ParseNode(url)
		if err != nil {
			log.Error("Bootstrap URL invalid", "enode", url, "err", err)
			continue
		}
		cfg.BootstrapNodes = append(cfg.BootstrapNodes, node)
	}
}

// setBootstrapNodesV5 creates a list of bootstrap nodes from the command line
// flags, reverting to pre-configured ones if none have been specified.
func setBootstrapNodesV5(ctx *cli.Context, cfg *p2p.Config) {
	urls := params.MainnetDiscoveryV5Bootnodes

	switch {
	case ctx.GlobalBool(TestnetFlag.Name):
		urls = params.TestnetDiscoveryV5Bootnodes
	case ctx.GlobalBool(DevModeFlag.Name):
		urls = params.DevnetDiscoveryV5Bootnodes
	}

	switch {
	case ctx.GlobalIsSet(BootnodesFlag.Name) || ctx.GlobalIsSet(BootnodesV5Flag.Name):
		if ctx.GlobalIsSet(BootnodesV5Flag.Name) {
			urls = strings.Split(ctx.GlobalString(BootnodesV5Flag.Name), ",")
		} else {
			urls = strings.Split(ctx.GlobalString(BootnodesFlag.Name), ",")
		}
	case cfg.BootstrapNodesV5 != nil:
		return // already set, don't apply defaults.
	}

	cfg.BootstrapNodesV5 = make([]*discv5.Node, 0, len(urls))
	for _, url := range urls {
		node, err := discv5.ParseNode(url)
		if err != nil {
			log.Error("Bootstrap URL invalid", "enode", url, "err", err)
			continue
		}
		cfg.BootstrapNodesV5 = append(cfg.BootstrapNodesV5, node)
	}
}

// setListenAddress creates a TCP listening address string from set command
// line flags.
func setListenAddress(ctx *cli.Context, cfg *p2p.Config) {
	port := ListenPortFlag.Value
	if ctx.GlobalIsSet(ListenPortFlag.Name) {
		port = ctx.GlobalInt(ListenPortFlag.Name)
	}
	cfg.ListenAddr = fmt.Sprintf(":%d", port)
}

// setDiscoveryV5Address creates a UDP listening address string from set command
// line flags for the V5 discovery protocol.
// fixme: uses ListenPortFlag
func setDiscoveryV5Address(ctx *cli.Context, cfg *p2p.Config) {
	if len(cfg.BootstrapNodesV5) > 0 {
		port := ListenPortFlag.Value
		if ctx.GlobalIsSet(ListenPortFlag.Name) {
			port = ctx.GlobalInt(ListenPortFlag.Name)
		}
		cfg.DiscoveryV5Addr = fmt.Sprintf(":%d", port+1)
	}
}

// setNAT creates a port mapper from command line flags.
func setNAT(ctx *cli.Context, cfg *p2p.Config) {
	if ctx.GlobalIsSet(NATFlag.Name) {
		natif, err := nat.Parse(ctx.GlobalString(NATFlag.Name))
		if err != nil {
			Fatalf("Option %s: %v", NATFlag.Name, err)
		}
		cfg.NAT = natif
	}
}

// splitAndTrim splits input separated by a comma
// and trims excessive white space from the substrings.
func splitAndTrim(input string) []string {
	result := strings.Split(input, ",")
	for i, r := range result {
		result[i] = strings.TrimSpace(r)
	}
	return result
}

// setHTTP creates the HTTP RPC listener interface string from the set
// command line flags, returning empty if the HTTP endpoint is disabled.
func setHTTP(ctx *cli.Context, cfg *node.Config) {
	if ctx.GlobalBool(RPCEnabledFlag.Name) && cfg.HTTPHost == "" {
		cfg.HTTPHost = "127.0.0.1"
		if ctx.GlobalIsSet(RPCListenAddrFlag.Name) {
			cfg.HTTPHost = ctx.GlobalString(RPCListenAddrFlag.Name)
		}
	}

	if ctx.GlobalIsSet(RPCPortFlag.Name) {
		cfg.HTTPPort = ctx.GlobalInt(RPCPortFlag.Name)
	}
	if ctx.GlobalIsSet(RPCCORSDomainFlag.Name) {
		cfg.HTTPCors = splitAndTrim(ctx.GlobalString(RPCCORSDomainFlag.Name))
	}
	if ctx.GlobalIsSet(RPCApiFlag.Name) {
		cfg.HTTPModules = splitAndTrim(ctx.GlobalString(RPCApiFlag.Name))
	}
	if ctx.GlobalIsSet(RPCVirtualHostsFlag.Name) {
		cfg.HTTPVirtualHosts = splitAndTrim(ctx.GlobalString(RPCVirtualHostsFlag.Name))
	}
}

// setWS creates the WebSocket RPC listener interface string from the set
// command line flags, returning empty if the HTTP endpoint is disabled.
func setWS(ctx *cli.Context, cfg *node.Config) {
	if ctx.GlobalBool(WSEnabledFlag.Name) && cfg.WSHost == "" {
		cfg.WSHost = "127.0.0.1"
		if ctx.GlobalIsSet(WSListenAddrFlag.Name) {
			cfg.WSHost = ctx.GlobalString(WSListenAddrFlag.Name)
		}
	}

	if ctx.GlobalIsSet(WSPortFlag.Name) {
		cfg.WSPort = ctx.GlobalInt(WSPortFlag.Name)
	}
	if ctx.GlobalIsSet(WSAllowedOriginsFlag.Name) {
		cfg.WSOrigins = splitAndTrim(ctx.GlobalString(WSAllowedOriginsFlag.Name))
	}
	if ctx.GlobalIsSet(WSApiFlag.Name) {
		cfg.WSModules = splitAndTrim(ctx.GlobalString(WSApiFlag.Name))
	}
}

// setIPC creates an IPC path configuration from the set command line flags,
// returning an empty string if IPC was explicitly disabled, or the set path.
func setIPC(ctx *cli.Context, cfg *node.Config) {
	checkExclusive(ctx, IPCDisabledFlag, IPCPathFlag)
	switch {
	case ctx.GlobalBool(IPCDisabledFlag.Name):
		cfg.IPCPath = ""
	case ctx.GlobalIsSet(IPCPathFlag.Name):
		cfg.IPCPath = ctx.GlobalString(IPCPathFlag.Name)
	}
}

// makeDatabaseHandles raises out the number of allowed file handles per process
// for kcoin and returns half of the allowance to assign to the database.
func makeDatabaseHandles() int {
	limit, err := fdlimit.Current()
	if err != nil {
		Fatalf("Failed to retrieve file descriptor allowance: %v", err)
	}
	if limit < 2048 {
		if err := fdlimit.Raise(2048); err != nil {
			Fatalf("Failed to raise file descriptor allowance: %v", err)
		}
	}
	if limit > 2048 { // cap database file descriptors even if more is available
		limit = 2048
	}
	return limit / 2 // Leave half for networking and other stuff
}

// MakeAddress converts an account specified directly as a hex encoded string or
// a key index in the key store to an internal account representation.
func MakeAddress(ks *keystore.KeyStore, account string) (accounts.Account, error) {
	// If the specified account is a valid address, return it
	if common.IsHexAddress(account) {
		return accounts.Account{Address: common.HexToAddress(account)}, nil
	}
	// Otherwise try to interpret the account as a keystore index
	index, err := strconv.Atoi(account)
	if err != nil || index < 0 {
		return accounts.Account{}, fmt.Errorf("invalid account address or index %q", account)
	}
	log.Warn("-------------------------------------------------------------------")
	log.Warn("Referring to accounts by order in the keystore folder is dangerous!")
	log.Warn("This functionality is deprecated and will be removed in the future!")
	log.Warn("Please use explicit addresses! (can search via `geth account list`)")
	log.Warn("-------------------------------------------------------------------")

	accs := ks.Accounts()
	if len(accs) <= index {
		return accounts.Account{}, fmt.Errorf("index %d higher than number of accounts %d", index, len(accs))
	}
	return accs[index], nil
}

// setCoinbase retrieves the coinbase either from the directly specified
// command line flags or from the keystore if CLI indexed.
func setCoinbase(ctx *cli.Context, ks *keystore.KeyStore, cfg *knode.Config) {
	if ctx.GlobalIsSet(CoinbaseFlag.Name) {
		account, err := MakeAddress(ks, ctx.GlobalString(CoinbaseFlag.Name))
		if err != nil {
			Fatalf("Option %q: %v", CoinbaseFlag.Name, err)
		}
		cfg.Coinbase = account.Address
		return
	}
	accounts := ks.Accounts()
	if (cfg.Coinbase == common.Address{}) {
		if len(accounts) > 0 {
			cfg.Coinbase = accounts[0].Address
		} else {
			log.Warn("No coinbase set and no accounts found as default")
		}
	}
}

func setDeposit(ctx *cli.Context, cfg *knode.Config) {
	if ctx.GlobalIsSet(ValidatorDepositFlag.Name) {
		cfg.Deposit = new(big.Int).SetUint64(ctx.GlobalUint64(ValidatorDepositFlag.Name))
	}
}

// MakePasswordList reads password lines from the file specified by the global --password flag.
func MakePasswordList(ctx *cli.Context) []string {
	path := ctx.GlobalString(PasswordFileFlag.Name)
	if path == "" {
		return nil
	}
	text, err := ioutil.ReadFile(path)
	if err != nil {
		Fatalf("Failed to read password file: %v", err)
	}
	lines := strings.Split(string(text), "\n")
	// Sanitise DOS line endings.
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], "\r")
	}
	return lines
}

func SetP2PConfig(ctx *cli.Context, cfg *p2p.Config) {
	setNodeKey(ctx, cfg)
	setNAT(ctx, cfg)
	setBootstrapNodes(ctx, cfg)
	setBootstrapNodesV5(ctx, cfg)
	setListenAddress(ctx, cfg)
	setDiscoveryV5Address(ctx, cfg)

	if ctx.GlobalIsSet(MaxPeersFlag.Name) {
		cfg.MaxPeers = ctx.GlobalInt(MaxPeersFlag.Name)
	}
	if ctx.GlobalIsSet(MaxPendingPeersFlag.Name) {
		cfg.MaxPendingPeers = ctx.GlobalInt(MaxPendingPeersFlag.Name)
	}
	if ctx.GlobalIsSet(NoDiscoverFlag.Name) || ctx.GlobalBool(LightModeFlag.Name) {
		cfg.NoDiscovery = true
	}

	// Unless --nodisover flag is set, enable V5 discovery
	if !ctx.GlobalBool(NoDiscoverFlag.Name) {
		cfg.DiscoveryV5 = true
	}

	if netrestrict := ctx.GlobalString(NetrestrictFlag.Name); netrestrict != "" {
		list, err := netutil.ParseNetlist(netrestrict)
		if err != nil {
			Fatalf("Option %q: %v", NetrestrictFlag.Name, err)
		}
		cfg.NetRestrict = list
	}
}

// SetNodeConfig applies node-related command line flags to the config.
func SetNodeConfig(ctx *cli.Context, cfg *node.Config, kowalaCfg *knode.Config) {
	SetP2PConfig(ctx, &cfg.P2P)
	setIPC(ctx, cfg)
	setHTTP(ctx, cfg)
	setWS(ctx, cfg)
	setNodeUserIdent(ctx, cfg)

	switch {
	case ctx.GlobalIsSet(DataDirFlag.Name):
		cfg.DataDir = ctx.GlobalString(DataDirFlag.Name)
	}

	if ctx.GlobalIsSet(KeyStoreDirFlag.Name) {
		cfg.KeyStoreDir = ctx.GlobalString(KeyStoreDirFlag.Name)
	}
	if ctx.GlobalIsSet(LightKDFFlag.Name) {
		cfg.UseLightweightKDF = ctx.GlobalBool(LightKDFFlag.Name)
	}
	if ctx.GlobalIsSet(NoUSBFlag.Name) {
		cfg.NoUSB = ctx.GlobalBool(NoUSBFlag.Name)
	}

	cfg.DataDir = filepath.Join(cfg.DataDir, kowalaCfg.Currency)
}

func setGPO(ctx *cli.Context, cfg *gasprice.Config) {
	if ctx.GlobalIsSet(GpoBlocksFlag.Name) {
		cfg.Blocks = ctx.GlobalInt(GpoBlocksFlag.Name)
	}
	if ctx.GlobalIsSet(GpoPercentileFlag.Name) {
		cfg.Percentile = ctx.GlobalInt(GpoPercentileFlag.Name)
	}
}

func setTxPool(ctx *cli.Context, cfg *core.TxPoolConfig) {
	if ctx.GlobalIsSet(TxPoolNoLocalsFlag.Name) {
		cfg.NoLocals = ctx.GlobalBool(TxPoolNoLocalsFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolJournalFlag.Name) {
		cfg.Journal = ctx.GlobalString(TxPoolJournalFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolRejournalFlag.Name) {
		cfg.Rejournal = ctx.GlobalDuration(TxPoolRejournalFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolPriceLimitFlag.Name) {
		cfg.PriceLimit = ctx.GlobalUint64(TxPoolPriceLimitFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolPriceBumpFlag.Name) {
		cfg.PriceBump = ctx.GlobalUint64(TxPoolPriceBumpFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolAccountSlotsFlag.Name) {
		cfg.AccountSlots = ctx.GlobalUint64(TxPoolAccountSlotsFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolGlobalSlotsFlag.Name) {
		cfg.GlobalSlots = ctx.GlobalUint64(TxPoolGlobalSlotsFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolAccountQueueFlag.Name) {
		cfg.AccountQueue = ctx.GlobalUint64(TxPoolAccountQueueFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolGlobalQueueFlag.Name) {
		cfg.GlobalQueue = ctx.GlobalUint64(TxPoolGlobalQueueFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolLifetimeFlag.Name) {
		cfg.Lifetime = ctx.GlobalDuration(TxPoolLifetimeFlag.Name)
	}
}

// checkExclusive verifies that only a single isntance of the provided flags was
// set by the user. Each flag might optionally be followed by a string type to
// specialize it further.
func checkExclusive(ctx *cli.Context, args ...interface{}) {
	set := make([]string, 0, 1)
	for i := 0; i < len(args); i++ {
		// Make sure the next argument is a flag and skip if not set
		flag, ok := args[i].(cli.Flag)
		if !ok {
			panic(fmt.Sprintf("invalid argument, not cli.Flag type: %T", args[i]))
		}
		// Check if next arg extends current and expand its name if so
		name := flag.GetName()

		if i+1 < len(args) {
			switch option := args[i+1].(type) {
			case string:
				// Extended flag, expand the name and shift the arguments
				if ctx.GlobalString(flag.GetName()) == option {
					name += "=" + option
				}
				i++

			case cli.Flag:
			default:
				panic(fmt.Sprintf("invalid argument, not cli.Flag or string extension: %T", args[i+1]))
			}
		}
		// Mark the flag if it's set
		if ctx.GlobalIsSet(flag.GetName()) {
			set = append(set, "--"+name)
		}
	}
	if len(set) > 1 {
		Fatalf("Flags %v can't be used at the same time", strings.Join(set, ", "))
	}
}

// SetKowalaConfig applies kowala-related command line flags to the config.
func SetKowalaConfig(ctx *cli.Context, stack *node.Node, cfg *knode.Config) {
	// Avoid conflicting network flags
	checkExclusive(ctx, DevModeFlag, TestnetFlag)
	checkExclusive(ctx, FastSyncFlag, LightModeFlag, SyncModeFlag)

	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
	setCoinbase(ctx, ks, cfg)
	setDeposit(ctx, cfg)
	setGPO(ctx, &cfg.GPO)
	setTxPool(ctx, &cfg.TxPool)

	switch {
	case ctx.GlobalIsSet(SyncModeFlag.Name):
		cfg.SyncMode = *GlobalTextMarshaler(ctx, SyncModeFlag.Name).(*downloader.SyncMode)
	case ctx.GlobalBool(FastSyncFlag.Name):
		cfg.SyncMode = downloader.FastSync
	case ctx.GlobalBool(LightModeFlag.Name):
		cfg.SyncMode = downloader.LightSync
	}
	if ctx.GlobalIsSet(LightServFlag.Name) {
		cfg.LightServ = ctx.GlobalInt(LightServFlag.Name)
	}
	if ctx.GlobalIsSet(LightPeersFlag.Name) {
		cfg.LightPeers = ctx.GlobalInt(LightPeersFlag.Name)
	}
	if ctx.GlobalIsSet(NetworkIdFlag.Name) {
		cfg.NetworkId = ctx.GlobalUint64(NetworkIdFlag.Name)
	}

	if ctx.GlobalIsSet(CacheFlag.Name) || ctx.GlobalIsSet(CacheDatabaseFlag.Name) {
		cfg.DatabaseCache = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheDatabaseFlag.Name) / 100
	}
	cfg.DatabaseHandles = makeDatabaseHandles()

	if gcmode := ctx.GlobalString(GCModeFlag.Name); gcmode != "full" && gcmode != "archive" {
		Fatalf("--%s must be either 'full' or 'archive'", GCModeFlag.Name)
	}
	cfg.NoPruning = ctx.GlobalString(GCModeFlag.Name) == "archive"

	if ctx.GlobalIsSet(CacheFlag.Name) || ctx.GlobalIsSet(CacheGCFlag.Name) {
		cfg.TrieCache = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheGCFlag.Name) / 100
	}
	if ctx.GlobalIsSet(DocRootFlag.Name) {
		cfg.DocRoot = ctx.GlobalString(DocRootFlag.Name)
	}
	if ctx.GlobalIsSet(ExtraDataFlag.Name) {
		cfg.ExtraData = []byte(ctx.GlobalString(ExtraDataFlag.Name))
	}
	if ctx.GlobalIsSet(GasPriceFlag.Name) {
		cfg.GasPrice = GlobalBig(ctx, GasPriceFlag.Name)
	}
	if ctx.GlobalIsSet(VMEnableDebugFlag.Name) {
		// TODO(fjl): force-enable this in --dev mode
		cfg.EnablePreimageRecording = ctx.GlobalBool(VMEnableDebugFlag.Name)
	}

	// Override any default configs for hard coded networks.
	switch {
	case ctx.GlobalBool(TestnetFlag.Name):
		cfg.NetworkId = params.TestnetChainConfig.ChainID.Uint64()
	case ctx.GlobalBool(DevModeFlag.Name):
		// Use the main net network ID. This allows us to test the p2p under realistic conditions
		cfg.NetworkId = params.MainnetChainConfig.ChainID.Uint64()
	}
	// TODO(fjl): move trie cache generations into config
	if gen := ctx.GlobalInt(TrieCacheGenFlag.Name); gen > 0 {
		state.MaxTrieCacheGen = uint16(gen)
	}
}

// RegisterKowalaService adds a Kowala client to the stack.
func RegisterKowalaService(stack *node.Node, cfg *knode.Config) {
	var err error

	err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		fullNode, err := knode.New(ctx, cfg)
		return fullNode, err
	})

	if err != nil {
		Fatalf("Failed to register the Kowala service: %v", err)
	}
}

// RegisterKowalaStatsService configures the Kowala Stats daemon and adds it to
// the given node.
func RegisterKowalaStatsService(stack *node.Node, url string) {
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		// Retrieve both eth and les services
		var kowalaServ *knode.Kowala
		ctx.Service(&kowalaServ)

		return stats.New(url, kowalaServ)
	}); err != nil {
		Fatalf("Failed to register the Kowala Stats service: %v", err)
	}
}

// SetupNetwork configures the system for either the main net or some test network.
func SetupNetwork(ctx *cli.Context) {
	// TODO(fjl): move target gas limit into config
	params.TargetGasLimit = ctx.GlobalUint64(TargetGasLimitFlag.Name)
}

func SetupMetrics(ctx *cli.Context) {
	if metrics.Enabled {
		log.Info("Enabling metrics collection")
		var (
			enableExport = ctx.GlobalBool(MetricsEnableInfluxDBFlag.Name)
			endpoint     = ctx.GlobalString(MetricsInfluxDBEndpointFlag.Name)
			database     = ctx.GlobalString(MetricsInfluxDBDatabaseFlag.Name)
			username     = ctx.GlobalString(MetricsInfluxDBUsernameFlag.Name)
			password     = ctx.GlobalString(MetricsInfluxDBPasswordFlag.Name)
			hosttag      = ctx.GlobalString(MetricsInfluxDBHostTagFlag.Name)
		)

		if enableExport {
			log.Info("Enabling metrics export to InfluxDB")
			go influxdb.InfluxDBWithTags(metrics.DefaultRegistry, 10*time.Second, endpoint, database, username, password, "geth.", map[string]string{
				"host": hosttag,
			})
		}
	}
}

// MakeChainDatabase open an LevelDB using the flags passed to the client and will hard crash if it fails.
func MakeChainDatabase(ctx *cli.Context, stack *node.Node) kcoindb.Database {
	var (
		cache   = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheDatabaseFlag.Name) / 100
		handles = makeDatabaseHandles()
	)
	name := "chaindata"
	if ctx.GlobalBool(LightModeFlag.Name) {
		name = "lightchaindata"
	}
	chainDb, err := stack.OpenDatabase(name, cache, handles)
	if err != nil {
		Fatalf("Could not open database: %v", err)
	}
	return chainDb
}

func MakeGenesis(ctx *cli.Context) *core.Genesis {
	var genesis *core.Genesis
	switch {
	case ctx.GlobalBool(TestnetFlag.Name), ctx.GlobalBool(DevModeFlag.Name):
		genesis = core.DefaultTestnetGenesisBlock()
	}
	return genesis
}

// MakeChain creates a chain manager from set command line flags.
func MakeChain(ctx *cli.Context, stack *node.Node) (chain *core.BlockChain, chainDb kcoindb.Database) {
	var err error
	chainDb = MakeChainDatabase(ctx, stack)

	engine := konsensus.New(&params.KonsensusConfig{})
	config, _, err := core.SetupGenesisBlock(chainDb, MakeGenesis(ctx))
	if err != nil {
		Fatalf("%v", err)
	}

	if gcmode := ctx.GlobalString(GCModeFlag.Name); gcmode != "full" && gcmode != "archive" {
		Fatalf("--%s must be either 'full' or 'archive'", GCModeFlag.Name)
	}
	cache := &core.CacheConfig{
		Disabled:      ctx.GlobalString(GCModeFlag.Name) == "archive",
		TrieNodeLimit: knode.DefaultConfig.TrieCache,
		TrieTimeLimit: knode.DefaultConfig.TrieTimeout,
	}
	if ctx.GlobalIsSet(CacheFlag.Name) || ctx.GlobalIsSet(CacheGCFlag.Name) {
		cache.TrieNodeLimit = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheGCFlag.Name) / 100
	}
	vmcfg := vm.Config{EnablePreimageRecording: ctx.GlobalBool(VMEnableDebugFlag.Name)}
	chain, err = core.NewBlockChain(chainDb, cache, config, engine, vmcfg)
	if err != nil {
		Fatalf("Can't create BlockChain: %v", err)
	}
	return chain, chainDb
}

// MakeConsolePreloads retrieves the absolute paths for the console JavaScript
// scripts to preload before starting.
func MakeConsolePreloads(ctx *cli.Context) []string {
	// Skip preloading if there's nothing to preload
	if ctx.GlobalString(PreloadJSFlag.Name) == "" {
		return nil
	}
	// Otherwise resolve absolute paths and return them
	preloads := []string{}

	assets := ctx.GlobalString(JSpathFlag.Name)
	for _, file := range strings.Split(ctx.GlobalString(PreloadJSFlag.Name), ",") {
		preloads = append(preloads, common.AbsolutePath(assets, strings.TrimSpace(file)))
	}
	return preloads
}

// MigrateFlags sets the global flag from a local flag when it's set.
// This is a temporary function used for migrating old command/flags to the
// new format.
//
// e.g. kcoin account new --keystore /tmp/mykeystore --lightkdf
//
// is equivalent after calling this method with:
//
// kcoin --keystore /tmp/mykeystore --lightkdf account new
//
// This allows the use of the existing configuration functionality.
// When all flags are migrated this function can be removed and the existing
// configuration functionality must be changed that is uses local flags
func MigrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		for _, name := range ctx.FlagNames() {
			if ctx.IsSet(name) {
				ctx.GlobalSet(name, ctx.String(name))
			}
		}
		return action(ctx)
	}
}
