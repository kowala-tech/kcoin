package wallet

// const EndPointBackend = "docker"

// type TestContext struct {
// 	Testnet                  testnet.Testnet
// 	ContainerIDs             []string
// 	RedisContainerID         string
// 	TxPersistanceContainerID string
// 	WalletBackendContainerID string
// }

// func TestGetBlockHeight(t *testing.T) {
// 	dockerEngine, err := testnet.NewDockerEngineWithDefaultClient()
// 	if err != nil {
// 		t.Fatalf("Error getting docker engine. %s", err)
// 	}
// 	ctx := launchTestNetAndWalletEnvironment(t, dockerEngine)

// 	t.Run("It shows current blockHeight", func(t *testing.T) {
// 		blockHeight1, err := getBlockHeight()
// 		if err != nil {
// 			t.Fatalf("%s", err)
// 		}

// 		time.Sleep(2 * time.Second)

// 		blockHeight2, err := getBlockHeight()
// 		if err != nil {
// 			t.Fatalf("%s", err)
// 		}

// 		assert.True(t, blockHeight1.Cmp(blockHeight2) < 0)
// 	})

// 	t.Run("We can send a transaction and then get balance and txs from account.", func(t *testing.T) {
// 		keyStore := ctx.Testnet.GetKeyStore()
// 		account := ctx.Testnet.GetGenesisValidatorAccount()

// 		toAddr := common.HexToAddress("0x1f2284b6e214eceef5a97985c033a842017efefa")

// 		tx := types.NewTransaction(1, toAddr, big.NewInt(100), big.NewInt(50000), big.NewInt(10000), nil)

// 		signedTx, err := keyStore.SignTx(account, tx, big.NewInt(2))
// 		if err != nil {
// 			t.Fatalf("Error signing transaction. %s", err)
// 		}

// 		resp, err := sendTransaction(signedTx)
// 		if err != nil {
// 			t.Fatalf("Error sending signed transaction. %s", err)
// 		}

// 		assert.Equal(t, "ok", resp.Status)

// 		time.Sleep(10 * time.Second)

// 		respTxs, err := getTransactions(toAddr)
// 		if err != nil {
// 			t.Fatalf("Error sending getting transactions. %s", err)
// 		}

// 		assert.Len(t, respTxs.Transactions, 1)
// 		assert.Equal(t, respTxs.Transactions[0].To, toAddr.String())

// 		//Same but FROM point of view.
// 		respTxs, err = getTransactions(account.Address)
// 		if err != nil {
// 			t.Fatalf("Error sending getting transactions. %s", err)
// 		}

// 		assert.Len(t, respTxs.Transactions, 2) // Including validation first tx
// 		assert.Equal(t, respTxs.Transactions[1].From, account.Address.String())

// 		//Get balance of account
// 		balance, err := getBalance(toAddr)
// 		assert.Equal(t, big.NewInt(100), balance)
// 	})

// 	endEnvironment(t, ctx, dockerEngine)
// }

// func launchTestNetAndWalletEnvironment(t *testing.T, dockerEngine testnet.DockerEngine) *TestContext {
// 	testnet := testnet.NewTestnet(dockerEngine)
// 	err := testnet.Start()

// 	if err != nil {
// 		t.Fatalf("Error starting testnet. %s", err)
// 	}

// 	if !testnet.IsValidating() {
// 		t.Fatalf("Error with testnet, it is not validating.")
// 	}

// 	runRedis(t, dockerEngine, testnet.GetNetworkID())
// 	runTxPersistance(t, dockerEngine, testnet)
// 	runAPIRPC(t, dockerEngine, testnet)
// 	buildWalletBackend(t, dockerEngine, testnet)

// 	return &TestContext{
// 		Testnet: testnet,
// 		ContainerIDs: []string{
// 			"thebackend",
// 			"dbsync",
// 			"redis",
// 			"api",
// 		},
// 	}
// }

// func endEnvironment(t *testing.T, context *TestContext, dockerEngine testnet.DockerEngine) {
// 	for _, id := range context.ContainerIDs {
// 		err := dockerEngine.StopAndRemoveContainer(id)
// 		if err != nil {
// 			t.Fatalf("Error closing container. %s", err)
// 		}
// 	}

// 	err := context.Testnet.Stop()
// 	if err != nil {
// 		t.Fatalf("Error stoping testnet. %s", err)
// 	}
// }

// func runAPIRPC(t *testing.T, dockerEngine testnet.DockerEngine, testnet testnet.Testnet) {
// 	imageName := "kowalatech/backend_api"
// 	err := dockerEngine.PullImage(imageName)
// 	if err != nil {
// 		t.Fatalf("Error downloading sync image. %s", err)
// 	}

// 	err = dockerEngine.CreateContainer(
// 		imageName,
// 		"api",
// 		testnet.GetNetworkID(),
// 		nil,
// 		[]string{
// 			fmt.Sprintf("PORT=3000"),
// 			"REDIS_ADDR=redis:6379",
// 		},
// 		nil,
// 	)
// 	if err != nil {
// 		t.Fatalf("Error creating api image. %s", err)
// 	}

// 	err = dockerEngine.StartContainer("api")
// 	if err != nil {
// 		t.Fatalf("Error starting sync container. %s", err)
// 	}
// }

// func buildWalletBackend(t *testing.T, engine testnet.DockerEngine, testnet testnet.Testnet) {
// 	err := engine.BuildImage("../../Dockerfile", "backend")
// 	if err != nil {
// 		t.Fatalf("Error build backend image. %s", err)
// 	}

// 	err = engine.CreateContainer(
// 		"backend",
// 		"thebackend",
// 		testnet.GetNetworkID(),
// 		nil,
// 		[]string{
// 			fmt.Sprintf("NODE_ENDPOINT=http://%s:11223", testnet.GetValidatorID()),
// 		},
// 		map[int32]int32{
// 			8080: 8080,
// 		},
// 	)
// 	if err != nil {
// 		t.Fatalf("Error creating backend container. %s", err)
// 	}

// 	err = engine.StartContainer("thebackend")
// 	if err != nil {
// 		t.Fatalf("Error starting backend container. %s", err)
// 	}

// 	time.Sleep(10 * time.Second)
// }

// func runTxPersistance(t *testing.T, dockerEngine testnet.DockerEngine, testnet testnet.Testnet) {
// 	imageName := "kowalatech/transactions_persistance"
// 	err := dockerEngine.PullImage(imageName)
// 	if err != nil {
// 		t.Fatalf("Error downloading sync image. %s", err)
// 	}

// 	err = dockerEngine.CreateContainer(
// 		imageName,
// 		"dbsync",
// 		testnet.GetNetworkID(),
// 		nil,
// 		[]string{
// 			"REDIS_ADDR=redis:6379",
// 			fmt.Sprintf("TESTNET_RPC_ADDR=http://%s:11223", testnet.GetValidatorID()),
// 		},
// 		nil,
// 	)
// 	if err != nil {
// 		t.Fatalf("Error creating sync image. %s", err)
// 	}

// 	err = dockerEngine.StartContainer("dbsync")
// 	if err != nil {
// 		t.Fatalf("Error starting sync container. %s", err)
// 	}
// }

// func runRedis(t *testing.T, dockerEngine testnet.DockerEngine, networkName string) error {
// 	imageName := "redis:alpine"
// 	err := dockerEngine.PullImage(imageName)
// 	if err != nil {
// 		t.Fatalf("Error downloading redis image. %s", err)
// 	}

// 	err = dockerEngine.CreateContainer(imageName, "redis", networkName, nil, nil, nil)
// 	if err != nil {
// 		t.Fatalf("Error creating redis container. %s", err)
// 	}

// 	err = dockerEngine.StartContainer("redis")
// 	if err != nil {
// 		t.Fatalf("Error starting sync container. %s", err)
// 	}

// 	return err
// }

// func getBalance(addresses common.Address) (*big.Int, error) {
// 	res, err := http.Get(
// 		fmt.Sprintf("http://%s:8080/api/balance/%s", EndPointBackend, addresses.String()),
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("error connecting to wallet backend to get balance. %s", err)
// 	}

// 	rawResp, err := ioutil.ReadAll(res.Body)
// 	defer res.Body.Close()
// 	if err != nil {
// 		return nil, fmt.Errorf("error parsing response from wallet backend to get block height. %s", err)
// 	}

// 	var balanceResponse *command.BalanceResponse

// 	err = json.Unmarshal(rawResp, &balanceResponse)
// 	if err != nil {
// 		return nil, fmt.Errorf("error unmarshalling from json response. %s", err)
// 	}

// 	return balanceResponse.Balance, nil
// }

// func getBlockHeight() (*big.Int, error) {
// 	res, err := http.Get(
// 		fmt.Sprintf("http://%s:8080/api/blockheight", EndPointBackend),
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("error connecting to wallet backend to get block height. %s", err)
// 	}

// 	rawResp, err := ioutil.ReadAll(res.Body)
// 	defer res.Body.Close()
// 	if err != nil {
// 		return nil, fmt.Errorf("error parsing response from wallet backend to get block height. %s", err)
// 	}

// 	var blockHeightResponse *command.BlockHeightResponse

// 	err = json.Unmarshal(rawResp, &blockHeightResponse)
// 	if err != nil {
// 		return nil, fmt.Errorf("error unmarshalling from json response. %s", err)
// 	}

// 	return blockHeightResponse.BlockHeight, nil
// }

// func sendTransaction(tx *types.Transaction) (*command.BroadcastTransactionResponse, error) {
// 	rawTx, err := rlp.EncodeToBytes(tx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp, err := http.Get(
// 		fmt.Sprintf("http://%s:8080/api/broadcasttx/%x", EndPointBackend, rawTx),
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("error sending signed transaction. %s", err)
// 	}

// 	rawResp, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var appResp command.BroadcastTransactionResponse

// 	err = json.Unmarshal(rawResp, &appResp)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &appResp, nil
// }

// func getTransactions(addr common.Address) (*command.TransactionsResponse, error) {
// 	resp, err := http.Get(
// 		fmt.Sprintf("http://%s:8080/api/transactions/%s", EndPointBackend, addr.String()),
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("error getting transactions. %s", err)
// 	}

// 	rawResp, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var txResp command.TransactionsResponse
// 	err = json.Unmarshal(rawResp, &txResp)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &txResp, nil
// }
