package console

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/kowala-tech/kcoin/internal/jsre"
	"github.com/kowala-tech/kcoin/kcoin"
	"github.com/kowala-tech/kcoin/node"
	"github.com/kowala-tech/kcoin/accounts/keystore"
	"math/big"
	"github.com/kowala-tech/kcoin/params"
	"github.com/kowala-tech/kcoin/cluster"
	"github.com/kowala-tech/kcoin/kcoinclient"
	"github.com/kowala-tech/kcoin/accounts"
	"github.com/kowala-tech/kcoin/core/types"
	"context"
	"github.com/kowala-tech/kcoin"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/kcoin/downloader"
)

const (
	testInstance = "console-tester"
	testAddress  = "0x8605cdbbdb6d264aa742e77020dcbc58fcdce182"
)

// hookedPrompter implements UserPrompter to simulate use input via channels.
type hookedPrompter struct {
	scheduler chan string
}

func (p *hookedPrompter) PromptInput(prompt string) (string, error) {
	// Send the prompt to the tester
	select {
	case p.scheduler <- prompt:
	case <-time.After(time.Second):
		return "", errors.New("prompt timeout")
	}
	// Retrieve the response and feed to the console
	select {
	case input := <-p.scheduler:
		return input, nil
	case <-time.After(time.Second):
		return "", errors.New("input timeout")
	}
}

func (p *hookedPrompter) PromptPassword(prompt string) (string, error) {
	return "", errors.New("not implemented")
}
func (p *hookedPrompter) PromptConfirm(prompt string) (bool, error) {
	return false, errors.New("not implemented")
}
func (p *hookedPrompter) SetHistory(history []string)              {}
func (p *hookedPrompter) AppendHistory(command string)             {}
func (p *hookedPrompter) SetWordCompleter(completer WordCompleter) {}

// tester is a console test environment for the console tests to operate on.
type tester struct {
	workspace string
	stack     *node.Node
	kowala    *kcoin.Kowala
	console   *Console
	input     *hookedPrompter
	output    *bytes.Buffer
}

// newTester creates a test environment based on which the console can operate.
// Please ensure you call Close() on the returned tester to avoid leaks.
func newTester(t *testing.T, confOverride func(*kcoin.Config)) *tester {
	// Create a temporary storage for the node keys and initialize it
	workspace := "/home/eugene/kcoin/test/"
	//workspace, err := ioutil.TempDir("", "console-tester-")
	//if err != nil {
	//	t.Fatalf("failed to create temporary keystore: %v", err)
	//}

	// Create a networkless protocol stack and start an Kowala service within
	stack, err := node.New(&node.Config{DataDir: workspace, UseLightweightKDF: true, Name: testInstance})
	if err != nil {
		t.Fatalf("failed to create node: %v", err)
	}

	accountManager := stack.AccountManager()
	keyStore := accountManager.Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
	const accountPassword = "test"

	fmt.Println("qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq", keyStore.Accounts(), stack.DataDir())

	seederAccount, err := keyStore.NewAccount(accountPassword)
	if err != nil {
		t.Fatal("failed creating account", err)
	}
	keyStore.Unlock(seederAccount, accountPassword)

	newGenesis, err := cluster.GetGenesis(seederAccount.Address)
	if err != nil {
		t.Fatal("cant generate genesis", err)
	}

	kcoinConf := kcoin.DefaultConfig
	kcoinConf.Genesis = newGenesis

	kcoinConf.SyncMode = downloader.FullSync
	kcoinConf.LightPeers = 20
	kcoinConf.DatabaseCache = 20
	kcoinConf.GasPrice = big.NewInt(1)
	kcoinConf.TxPool.Journal = "transactions.rlp"
	kcoinConf.TxPool.Rejournal = time.Hour
	kcoinConf.TxPool.PriceLimit = 1
	kcoinConf.TxPool.PriceBump = 1
	kcoinConf.TxPool.AccountSlots = 16
	kcoinConf.TxPool.GlobalSlots = 4096
	kcoinConf.TxPool.GlobalQueue = 1024
	kcoinConf.TxPool.AccountQueue = 1024
	kcoinConf.TxPool.Lifetime = 3*time.Hour
	kcoinConf.GPO.Blocks = 10
	kcoinConf.GPO.Percentile = 50
	kcoinConf.MaxPeers = 25

	if confOverride != nil {
		confOverride(&kcoinConf)
	}

	fmt.Println("******************************** Alloc ACCOUNT", seederAccount.Address.String())

	var kowala *kcoin.Kowala
	if err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		var err error
		kowala, err = kcoin.New(ctx, &kcoinConf)
		if err != nil {
			return kowala, err
		}

		kowala.SetCoinbase(seederAccount.Address)

		return kowala, err
	}); err != nil {
		t.Fatalf("failed to register Kowala protocol: %v", err)
	}
	fmt.Println("Test BEFORE start", kowala)

	// Start the node and assemble the JavaScript console around it
	if err = stack.Start(); err != nil {
		t.Fatalf("failed to start test stack: %v", err)
	}
	client, err := stack.Attach()
	if err != nil {
		t.Fatalf("failed to attach to node: %v", err)
	}
	prompter := &hookedPrompter{scheduler: make(chan string)}
	printer := new(bytes.Buffer)

	console, err := New(Config{
		DataDir:  stack.DataDir(),
		DocRoot:  "testdata",
		Client:   client,
		Prompter: prompter,
		Printer:  printer,
		Preload:  []string{"preload.js"},
	})
	if err != nil {
		t.Fatalf("failed to create JavaScript console: %v", err)
	}

	//eth.sendTransaction({from:eth.coinbase,to: "0x259be75d96876f2ada3d202722523e9cd4dd917d",value: 1})
	kclient := kcoinclient.NewClient(client)
	_ = kclient

	for i:=0; i<10; i++ {
		tx, err := sendFunds(kclient, keyStore, kowala.BlockChain().Config().ChainID, seederAccount, common.HexToAddress("0x259be75d96876f2ada3d202722523e9cd4dd917d"), 1+int64(i))
		fmt.Println("TRANSACTION RESULT", err, tx.String())
		fmt.Println("BLOCK", kowala.ApiBackend.CurrentBlock().String())
		time.Sleep(5*time.Second)
		fmt.Println("\n\n\n----------------------------------------------------------------------------------------")
	}

	fmt.Println("Test start", kowala)
	if err := kowala.StartValidating(); err != nil {
		t.Fatalf("Failed to start validation: %v", err)
	}

	return &tester{
		workspace: workspace,
		stack:     stack,
		kowala:    kowala,
		console:   console,
		input:     prompter,
		output:    printer,
	}
}

func sendFunds(client *kcoinclient.Client, keyStore *keystore.KeyStore, chainID *big.Int, from accounts.Account, to common.Address, kcoin int64) (*types.Transaction, error) {
	nonce, err := client.NonceAt(context.Background(), from.Address, nil)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	gp, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	gas, err := client.EstimateGas(ctx, kowala.CallMsg{
		From:     from.Address,
		To:       &to,
		Value:    toWei(kcoin),
		GasPrice: gp,
	})
	if err != nil {
		return nil, err
	}

	tx := types.NewTransaction(nonce, to, toWei(kcoin), gas, gp, nil)

	tx, err = keyStore.SignTx(from, tx, chainID)
	if err != nil {
		return nil, err
	}

	return tx, client.SendTransaction(ctx, tx)
}

func toWei(kcoin int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(kcoin), big.NewInt(params.Ether))
}

// Close cleans up any temporary data folders and held resources.
func (env *tester) Close(t *testing.T) {
	fmt.Println("^^^^^^^^^^^ tester CLOSE()")
	if err := env.console.Stop(false); err != nil {
		t.Errorf("failed to stop embedded console: %v", err)
	}
	if err := env.stack.Stop(); err != nil {
		t.Errorf("failed to stop embedded node: %v", err)
	}
	//os.RemoveAll(env.workspace)
}

// Tests that the node lists the correct welcome message, notably that it contains
// the instance name, coinbase account, block number, data directory and supported
// console modules.
func TestWelcome(t *testing.T) {
	tester := newTester(t, nil)
	defer tester.Close(t)

	tester.console.Welcome()

	output := string(tester.output.Bytes())
	if want := "Welcome"; !strings.Contains(output, want) {
		t.Fatalf("console output missing welcome message: have\n%s\nwant also %s", output, want)
	}
	if want := fmt.Sprintf("instance: %s", testInstance); !strings.Contains(output, want) {
		t.Fatalf("console output missing instance: have\n%s\nwant also %s", output, want)
	}
	if want := fmt.Sprintf("coinbase: %s", testAddress); !strings.Contains(output, want) {
		t.Fatalf("console output missing coinbase: have\n%s\nwant also %s", output, want)
	}
	if want := "at block: 0"; !strings.Contains(output, want) {
		t.Fatalf("console output missing sync status: have\n%s\nwant also %s", output, want)
	}
	if want := fmt.Sprintf("datadir: %s", tester.workspace); !strings.Contains(output, want) {
		t.Fatalf("console output missing coinbase: have\n%s\nwant also %s", output, want)
	}
}

// Tests that JavaScript statement evaluation works as intended.
func TestEvaluate(t *testing.T) {
	tester := newTester(t, nil)
	defer tester.Close(t)

	tester.console.Evaluate("2 + 2")
	if output := string(tester.output.Bytes()); !strings.Contains(output, "4") {
		t.Fatalf("statement evaluation failed: have %s, want %s", output, "4")
	}

	time.Sleep(20 * time.Second)
}

// Tests that the console can be used in interactive mode.
func TestInteractive(t *testing.T) {
	// Create a tester and run an interactive console in the background
	tester := newTester(t, nil)
	defer tester.Close(t)

	go tester.console.Interactive()

	// Wait for a promt and send a statement back
	select {
	case <-tester.input.scheduler:
	case <-time.After(time.Second):
		t.Fatalf("initial prompt timeout")
	}
	select {
	case tester.input.scheduler <- "2+2":
	case <-time.After(time.Second):
		t.Fatalf("input feedback timeout")
	}
	// Wait for the second promt and ensure first statement was evaluated
	select {
	case <-tester.input.scheduler:
	case <-time.After(time.Second):
		t.Fatalf("secondary prompt timeout")
	}
	if output := string(tester.output.Bytes()); !strings.Contains(output, "4") {
		t.Fatalf("statement evaluation failed: have %s, want %s", output, "4")
	}
}

// Tests that preloaded JavaScript files have been executed before user is given
// input.
func TestPreload(t *testing.T) {
	tester := newTester(t, nil)
	defer tester.Close(t)

	tester.console.Evaluate("preloaded")
	if output := string(tester.output.Bytes()); !strings.Contains(output, "some-preloaded-string") {
		t.Fatalf("preloaded variable missing: have %s, want %s", output, "some-preloaded-string")
	}
}

// Tests that JavaScript scripts can be executes from the configured asset path.
func TestExecute(t *testing.T) {
	tester := newTester(t, nil)
	defer tester.Close(t)

	tester.console.Execute("exec.js")

	tester.console.Evaluate("execed")
	if output := string(tester.output.Bytes()); !strings.Contains(output, "some-executed-string") {
		t.Fatalf("execed variable missing: have %s, want %s", output, "some-executed-string")
	}
}

// Tests that the JavaScript objects returned by statement executions are properly
// pretty printed instead of just displaing "[object]".
func TestPrettyPrint(t *testing.T) {
	tester := newTester(t, nil)
	defer tester.Close(t)

	tester.console.Evaluate("obj = {int: 1, string: 'two', list: [3, 3, 3], obj: {null: null, func: function(){}}}")

	// Define some specially formatted fields
	var (
		one   = jsre.NumberColor("1")
		two   = jsre.StringColor("\"two\"")
		three = jsre.NumberColor("3")
		null  = jsre.SpecialColor("null")
		fun   = jsre.FunctionColor("function()")
	)
	// Assemble the actual output we're after and verify
	want := `{
  int: ` + one + `,
  list: [` + three + `, ` + three + `, ` + three + `],
  obj: {
    null: ` + null + `,
    func: ` + fun + `
  },
  string: ` + two + `
}
`
	if output := string(tester.output.Bytes()); output != want {
		t.Fatalf("pretty print mismatch: have %s, want %s", output, want)
	}
}

// Tests that the JavaScript exceptions are properly formatted and colored.
func TestPrettyError(t *testing.T) {
	tester := newTester(t, nil)
	defer tester.Close(t)
	tester.console.Evaluate("throw 'hello'")

	want := jsre.ErrorColor("hello") + "\n"
	if output := string(tester.output.Bytes()); output != want {
		t.Fatalf("pretty error mismatch: have %s, want %s", output, want)
	}
}

// Tests that tests if the number of indents for JS input is calculated correct.
func TestIndenting(t *testing.T) {
	testCases := []struct {
		input               string
		expectedIndentCount int
	}{
		{`var a = 1;`, 0},
		{`"some string"`, 0},
		{`"some string with (parentesis`, 0},
		{`"some string with newline
		("`, 0},
		{`function v(a,b) {}`, 0},
		{`function f(a,b) { var str = "asd("; };`, 0},
		{`function f(a) {`, 1},
		{`function f(a, function(b) {`, 2},
		{`function f(a, function(b) {
		     var str = "a)}";
		  });`, 0},
		{`function f(a,b) {
		   var str = "a{b(" + a, ", " + b;
		   }`, 0},
		{`var str = "\"{"`, 0},
		{`var str = "'("`, 0},
		{`var str = "\\{"`, 0},
		{`var str = "\\\\{"`, 0},
		{`var str = 'a"{`, 0},
		{`var obj = {`, 1},
		{`var obj = { {a:1`, 2},
		{`var obj = { {a:1}`, 1},
		{`var obj = { {a:1}, b:2}`, 0},
		{`var obj = {}`, 0},
		{`var obj = {
			a: 1, b: 2
		}`, 0},
		{`var test = }`, -1},
		{`var str = "a\""; var obj = {`, 1},
	}

	for i, tt := range testCases {
		counted := countIndents(tt.input)
		if counted != tt.expectedIndentCount {
			t.Errorf("test %d: invalid indenting: have %d, want %d", i, counted, tt.expectedIndentCount)
		}
	}
}
