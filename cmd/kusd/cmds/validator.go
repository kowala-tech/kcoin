package cmd

import (
	"runtime"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/kowala-tech/kUSD/node"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	cli "gopkg.in/urfave/cli.v1"
)


var (
	// validatorCmd contains the validator related commands
	ValidatorCmd = &cobra.Command{
		Use:  "validator",
		Args: cobra.NoArgs,
	}

	// @NOTE (rgeraldes) - graceful exit is done by catching a specific signal
	// @NOTE (rgeraldes) - confirm if it's on application exit
	// post is executed after the node process termination
	startValidator = &cobra.Command{
		Use:     "start",
		Short:   "Start the validator process",
		Args:    cobra.NoArgs,
		PreRun:  func(cmd *cobra.Command, args []string) {
			runtime.GOMAXPROCS(runtime.NumCPU())
			/*
				if err := debug.Setup(ctx); err != nil {
					return err
				}
		
				// start system runtime metrics collection
				go metrics.CollectProcessMetrics(3 * time.Second)
		
				// @TODO (rgeraldes) - missing statement
			*/
		},
		Run:     startValidator,
		PostRun: func (cmd *cobra.Command, args []string) {
			//debug.Exit()
			//console.Stdin.Close()
			return nil
		},
	}

	newValidator = &cobra.Command{
		Use: "new",
		Short: "Generate a new validator keypair and account"
		Args: cobra.NoArgs,
		Run: newValidator,
	}
)

func init() {
	// set parentship
	ValidatorCmd.AddCommand(startValidator, newValidator)	
	
	// startValidatorCmd flags
	startValidator.Flags().AddFlag(utils.DataDirFlag)
}

// newValidator generates a new identity and account 
func newValidator(cmd *cobra.Command, args []string) {
	// generate a new key file - validator identity
	newKeyFile(cmd, args)
	// generate a new account - validator account
	newAccount(cmd, args)
}

// startValidator starts the validator node
func startValidator(cmd *cobra.Command, args []string) {
	// init node
	node := node.New(config.Node)

	// register consensus service
	if err := consensus.RegisterService(node, cfg.Consensus); err != nil {
		utils.Fatalf("Failed to register the KUSD service: %v", err)
	}

	// start node
	startNode(node)
}

// startNode boots up the system node and all registered protocols
func startNode(ctx *cli.Context, node *node.Node) {
	// Start up the node itself
	utils.StartNode(node)

	/*
			// Unlock any account specifically requested
			ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)

			passwords := utils.MakePasswordList(ctx)
			unlocks := strings.Split(ctx.GlobalString(utils.UnlockedAccountFlag.Name), ",")
			for i, account := range unlocks {
				if trimmed := strings.TrimSpace(account); trimmed != "" {
					unlockAccount(ctx, ks, trimmed, i, passwords)
				}
			}
			// Register wallet event handlers to open and auto-derive wallets
			events := make(chan accounts.WalletEvent, 16)
			stack.AccountManager().Subscribe(events)

			go func() {
				// Create an chain state reader for self-derivation
				rpcClient, err := stack.Attach()
				if err != nil {
					utils.Fatalf("Failed to attach to self: %v", err)
				}
				stateReader := ethclient.NewClient(rpcClient)

				// Open any wallets already attached
				for _, wallet := range stack.AccountManager().Wallets() {
					if err := wallet.Open(""); err != nil {
						log.Warn("Failed to open wallet", "url", wallet.URL(), "err", err)
					}
				}
				// Listen for wallet event till termination
				for event := range events {
					switch event.Kind {
					case accounts.WalletArrived:
						if err := event.Wallet.Open(""); err != nil {
							log.Warn("New wallet appeared, failed to open", "url", event.Wallet.URL(), "err", err)
						}
					case accounts.WalletOpened:
						status, _ := event.Wallet.Status()
						log.Info("New wallet appeared", "url", event.Wallet.URL(), "status", status)

						if event.Wallet.URL().Scheme == "ledger" {
							event.Wallet.SelfDerive(accounts.DefaultLedgerBaseDerivationPath, stateReader)
						} else {
							event.Wallet.SelfDerive(accounts.DefaultBaseDerivationPath, stateReader)
						}

					case accounts.WalletDropped:
						log.Info("Old wallet dropped", "url", event.Wallet.URL())
						event.Wallet.Close()
					}
				}
			}()


		// Start auxiliary services if enabled
		if ctx.GlobalBool(utils.MiningEnabledFlag.Name) || ctx.GlobalBool(utils.DeveloperFlag.Name) {
			// Mining only makes sense if a full Ethereum node is running
			var ethereum *eth.Ethereum
			if err := stack.Service(&ethereum); err != nil {
				utils.Fatalf("ethereum service not running: %v", err)
			}
			// Use a reduced number of threads if requested
			if threads := ctx.GlobalInt(utils.MinerThreadsFlag.Name); threads > 0 {
				type threaded interface {
					SetThreads(threads int)
				}
				if th, ok := ethereum.Engine().(threaded); ok {
					th.SetThreads(threads)
				}
			}
			// Set the gas price to the limits from the CLI and start mining
			ethereum.TxPool().SetGasPrice(utils.GlobalBig(ctx, utils.GasPriceFlag.Name))
			if err := ethereum.StartMining(true); err != nil {
				utils.Fatalf("Failed to start mining: %v", err)
			}
		}

	*/

}

