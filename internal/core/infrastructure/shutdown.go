package infrastructure

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

// OnShutdown registers a function to be called when the application receives a
// termination signal. The function is executed in a separate goroutine.
func OnShutdown(
	logger *logrus.Logger,
	callback func() error,
) {
	go (func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		s := <-signals
		logger.Printf("Received signal '%v'. Shutting down...", s.String())

		if err := callback(); err != nil {
			logger.Fatalf("Error during graceful shutdown: %v", err)
		}
	})()
}
