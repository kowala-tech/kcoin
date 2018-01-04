package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// @Note (rgeraldes) - backup keys

// AccountCmd contains the account related commands
var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage accounts",
	Long: `
	Manage accounts, list all existing accounts, import a private key into a new
	account, create a new account or update an existing account.
	It supports interactive mode, when you are prompted for password as well as
	non-interactive mode where passwords are supplied via a given password file.
	Non-interactive mode is only meant for scripted use on test networks or known
	safe environments.
	Make sure you remember the password you gave when creating a new account (with
	either new or import). Without it you are not able to unlock your account.
	Note that exporting your key in unencrypted format is NOT supported.
	Keys are stored under <DATADIR>/keystore.
	It is safe to transfer the entire directory or the individual keys therein
	between ethereum nodes by simply copying.
	Make sure you backup your keys regularly.`
	Args:  cobra.NoArgs,
}

func init() {
	// new
	AccountCmd.AddCommand(&cobra.Command{
		Use: "new",
		Short: "Create a new account",
		Args: cobra.NoArgs,
		Run: newAccount,
	})
	
	// list
	AccountCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "Print summary of existing accounts",
		Args:  cobra.NoArgs,
		Run:   listAccounts,
	})

	// import 
	AccountCmd.AddCommand(&cobra.Command{
		Use: "import",
		Short: "Import a private key into a new account",
		Args: cobra.ExactArgs(1),
		Run: importAccount,
	})
}

// listAccounts lists all the existing accounts
func listAccounts(cmd *cobra.Command, args []string) {
	node, _ := new.Node(config.Node)
	var index int
	for _, wallet := range node.AccountManager().Wallets() {
		for _, account := range wallet.Accounts() {
			fmt.Printf("Account #%d: {%x} %s\n", index, account.Address, &account.URL)
			index++
		}
	}
	return nil
}

// newAccount creates a new account into the keystore defined by the CLI flags.
func newAccount(cmd *cobra.Command, args []string) {
	scryptN, scryptP, keydir, err := config.Node.AccountConfig()
	if err != nil {
		utils.Fatalf("Failed to read configuration: %v", err)
	}

	// generate key
	password := getPassPhrase("Your new account is locked with a password. Please give a password. Do not forget this password.", true, 0, utils.MakePasswordList(ctx))

	// store key
	address, err := keystore.StoreKey(keydir, password, scryptN, scryptP)
	if err != nil {
		utils.Fatalf("Failed to create account: %v", err)
	}
	
	fmt.Printf("Address: {%x}\n", address)
	return nil
}

// importAccount imports a private key into a new account
func importAccount(cmd *cobra.Command, args []string) {
	keyfile := args[0]
	key, err := crypto.LoadECDSA(keyfile)
	if err != nil {
		utils.Fatalf("Failed to load the private key: %v", err)
	}
	stack, _ := makeConfigNode(ctx)
	passphrase := getPassPhrase("Your new account is locked with a password. Please give a password. Do not forget this password.", true, 0, utils.MakePasswordList(ctx))

	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
	acct, err := ks.ImportECDSA(key, passphrase)
	if err != nil {
		utils.Fatalf("Could not create the account: %v", err)
	}
	fmt.Printf("Address: {%x}\n", acct.Address)
	return nil
}

// getPassPhrase retrieves the password associated with an account, either fetched
// from a list of preloaded passphrases, or requested interactively from the user.
func getPassPhrase(prompt string, confirmation bool, i int, passwords []string) string {
	// If a list of passwords was supplied, retrieve from them
	if len(passwords) > 0 {
		if i < len(passwords) {
			return passwords[i]
		}
		return passwords[len(passwords)-1]
	}
	// Otherwise prompt the user for the password
	if prompt != "" {
		fmt.Println(prompt)
	}
	password, err := console.Stdin.PromptPassword("Passphrase: ")
	if err != nil {
		utils.Fatalf("Failed to read passphrase: %v", err)
	}
	if confirmation {
		confirm, err := console.Stdin.PromptPassword("Repeat passphrase: ")
		if err != nil {
			utils.Fatalf("Failed to read passphrase confirmation: %v", err)
		}
		if password != confirm {
			utils.Fatalf("Passphrases do not match")
		}
	}
	return password
}