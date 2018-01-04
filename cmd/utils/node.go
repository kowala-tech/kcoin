package utils

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/kowala-tech/kUSD/internal/debug"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/node"
)

// StartNode launches the node processes
func StartNode(stack *node.Node) error {
	if err := stack.Start(); err != nil {
		return fmt.Errorf("Error starting protocol stack: %v", err)
	}
	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, os.Interrupt)
		defer signal.Stop(sigc)
		<-sigc
		log.Info("Got interrupt, shutting down...")
		go stack.Stop()
		for i := 10; i > 0; i-- {
			<-sigc
			if i > 1 {
				log.Warn("Already shutting down, interrupt more to panic.", "times", i-1)
			}
		}
		debug.Exit() // ensure trace and CPU profile data is flushed.
		debug.LoudPanic("boom")
	}()
	return nil
}
