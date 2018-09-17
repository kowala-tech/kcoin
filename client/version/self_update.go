package version

import (
	"github.com/kowala-tech/kcoin/client/internal/debug"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/node"
	"math/rand"
	"time"
)

type SelfUpdater struct {
	repository string
	stack      *node.Node
	logger     log.Logger
}

func NewSelfUpdater(repository string, stack *node.Node, logger log.Logger) *SelfUpdater {
	return &SelfUpdater{
		repository: repository,
		stack:      stack,
		logger:     logger,
	}
}

func (su *SelfUpdater) Run() {
	updater, err := NewUpdater(su.repository, su.logger)
	if err != nil {
		su.logger.Warn("Error starting update for selfupdate, selfupdate disabled")
		return
	}

	randomTo60SecDelay()

	for range time.Tick(time.Minute) {
		su.logger.Debug("Checking if newer version is available")
		isLatest, err := updater.isCurrentLatestForMajor()
		if err != nil {
			continue
		}

		if !isLatest {
			su.logger.Info("Binary is not latest, updating")
			updater.Update()
			su.logger.Info("Exiting node")
			su.exit()
		}
	}
}

func (su *SelfUpdater) exit() {
	go su.stack.Stop()

	// wait 10 seconds for graceful shutdown
	time.Sleep(time.Second * 10)

	debug.Exit()
}

func randomTo60SecDelay() {
	randomTo60 := time.Duration(randomIntInRange(1, 60) * int(time.Second))
	time.Sleep(randomTo60)
}

func randomIntInRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
