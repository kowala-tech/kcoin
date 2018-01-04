package cmd

import "github.com/spf13/cobra"

var KeyCmd = &cobra.Command{
	Use: "key"
}

func init() {
	// new key file
	KeyCmd.AddCommand(&cobra.Command{
		Use: "new",
		Short: "Generate a new keyfile",
		Args: cobra.NoArgs,
		Run: newKeyFile,
	})

	// inspect key file
	KeyCmd.AddCommand(&cobra.Command{
		Use: "inspect",
		Short: "Print various information about the keyfile",
		Args: cobra.ExactArgs(1),
		Use: inspectKeyFile,
	}) 
}

// newKeyFile generates a new key file
func newKeyFile(cmd *cobra.Command, args []string) {
	keyfilepath = defaultKeyfileName

	if _, err := os.Stat(keyfilepath); err == nil {
		utils.Fatalf("Keyfile already exists at %s.", keyfilepath)
	} else if !os.IsNotExist(err) {
		utils.Fatalf("Error checking if keyfile exists: %v", err)
	}

	var privateKey *ecdsa.PrivateKey

	// First check if a private key file is provided.
	privateKeyFile := ctx.String("privatekey")
	if privateKeyFile != "" {
		privateKeyBytes, err := ioutil.ReadFile(privateKeyFile)
		if err != nil {
			utils.Fatalf("Failed to read the private key file '%s': %v",
				privateKeyFile, err)
		}

		pk, err := crypto.HexToECDSA(string(privateKeyBytes))
		if err != nil {
			utils.Fatalf(
				"Could not construct ECDSA private key from file content: %v",
				err)
		}
		privateKey = pk
	}

	// If not loaded, generate random.
	if privateKey == nil {
		pk, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
		if err != nil {
			utils.Fatalf("Failed to generate random private key: %v", err)
		}
		privateKey = pk
	}

	// Create the keyfile object with a random UUID.
	id := uuid.NewRandom()
	key := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(privateKey.PublicKey),
		PrivateKey: privateKey,
	}

	// Encrypt key with passphrase.
	passphrase := getPassPhrase(ctx, true)
	keyjson, err := keystore.EncryptKey(key, passphrase,
		keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		utils.Fatalf("Error encrypting key: %v", err)
	}

	// Store the file to disk.
	if err := os.MkdirAll(filepath.Dir(keyfilepath), 0700); err != nil {
		utils.Fatalf("Could not create directory %s", filepath.Dir(keyfilepath))
	}
	if err := ioutil.WriteFile(keyfilepath, keyjson, 0600); err != nil {
		utils.Fatalf("Failed to write keyfile to %s: %v", keyfilepath, err)
	}

	// Output some information.
	out := outputGenerate{
		Address: key.Address.Hex(),
	}
	if ctx.Bool(jsonFlag.Name) {
		mustPrintJSON(out)
	} else {
		fmt.Println("Address:       ", out.Address)
	}
	return nil
}

// inspectKeyFile prints various information about the keyfile
func inspectKeyFile(cmd *cobra.Command, args []string) {
	keyfilepath := args[0]

	// Read key from file.
	keyjson, err := ioutil.ReadFile(keyfilepath)
	if err != nil {
		utils.Fatalf("Failed to read the keyfile at '%s': %v", keyfilepath, err)
	}

	// Decrypt key with passphrase.
	passphrase := getPassPhrase(ctx, false)
	key, err := keystore.DecryptKey(keyjson, passphrase)
	if err != nil {
		utils.Fatalf("Error decrypting key: %v", err)
	}

	// Output all relevant information we can retrieve.
	showPrivate := ctx.Bool("private")
	out := outputInspect{
		Address: key.Address.Hex(),
		PublicKey: hex.EncodeToString(
			crypto.FromECDSAPub(&key.PrivateKey.PublicKey)),
	}
	if showPrivate {
		out.PrivateKey = hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
	}

	if ctx.Bool(jsonFlag.Name) {
		mustPrintJSON(out)
	} else {
		fmt.Println("Address:       ", out.Address)
		fmt.Println("Public key:    ", out.PublicKey)
		if showPrivate {
			fmt.Println("Private key:   ", out.PrivateKey)
		}
	}
	return nil
}