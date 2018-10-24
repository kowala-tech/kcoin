package main

import (
	"flag"
	"os"
	"strings"

	"github.com/DATA-DOG/godog"
)

type stringArr []string

func (i *stringArr) String() string {
	return strings.Join([]string(*i), ", ")
}

func (i *stringArr) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var (
	featuresFlag   stringArr
	stdErrLogsFlag bool
)

func main() {
	flag.Var(&featuresFlag, "features", "Specify the path to the features files. Supports multiple paths.")
	flag.BoolVar(&stdErrLogsFlag, "stdout-logs", false, "Send logs to the standard output instead to the logs/ directory.")
	flag.Parse()

	if len(featuresFlag) == 0 {
		featuresFlag = stringArr{"/features"}
	}

	status := godog.RunWithOptions("features", func(s *godog.Suite) {
		FeatureContext(&FeatureContextOpts{
			suite:        s,
			logsToStdout: stdErrLogsFlag,
		})
	}, godog.Options{
		Format:        "progress",
		Concurrency:   1,
		Paths:         []string(featuresFlag),
		Randomize:     -1,
		StopOnFailure: true,
		Tags:          "bootes",
	})

	os.Exit(status)
}
