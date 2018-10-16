package main

import (
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/knode/genesis"
)

var (
	packageFlag  = flag.String("package", "genesis", "Package name to include in generated output file")
	currencyFlag = flag.String("currency", "kusd", "Currency to generate genesis for")
)

var template = "// Auto-generated with genesisgen, do not edit!\n\npackage %s\n\nvar Generated%s = map[string][]byte { \n\"%s\": []byte(`%s`), \n\"%s\": []byte(`%s`),\n}"

func main() {

	flag.Parse()

	mainnetJson := mustGetGenesisJson(*currencyFlag, genesis.MainNetwork, mustFindGenesis(*currencyFlag, genesis.MainNetwork))
	testnetJson := mustGetGenesisJson(*currencyFlag, genesis.TestNetwork, mustFindGenesis(*currencyFlag, genesis.TestNetwork))

	outputFile := fmt.Sprintf("%s_generated.go", *currencyFlag)

	content := fmt.Sprintf(template,
		*packageFlag,
		strings.ToUpper(*currencyFlag),
		genesis.MainNetwork,
		string(mainnetJson),
		genesis.TestNetwork,
		string(testnetJson),
	)

	formatted, err := format.Source([]byte(content))

	if err != nil {
		fmt.Printf("Genesis code format for %s failed: %s", *currencyFlag, err)
		os.Exit(-1)
	}

	if err := ioutil.WriteFile(outputFile, formatted, 0600); err != nil {
		fmt.Printf("Genesis code file write for %s failed: %s", *currencyFlag, err)
		os.Exit(-1)
	}
}

func mustFindGenesis(currency, network string) *core.Genesis {

	gen, err := genesis.NetworkGenesisBlock("", currency, network)

	if err != nil {
		fmt.Printf("Genesis generation for %s (%s) failed: %s", currency, network, err)
		os.Exit(-1)
	}

	return gen
}

func mustGetGenesisJson(currency, network string, gen *core.Genesis) []byte {

	json, err := gen.MarshalJSON()

	if err != nil {
		fmt.Printf("Genesis marshal for %s (%s) failed: %s", currency, network, err)
		os.Exit(-1)
	}

	return json
}
