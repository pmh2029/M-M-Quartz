package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogger creates a new instance of a logger with the given configuration
func NewLogger() *logrus.Logger {
	// create a new instance of a logger
	logger := logrus.New()

	if os.Getenv("STAGE_MODE") != "production" {
		logger.SetOutput(os.Stdout)
	} else {
		// open the log file for writing
		logFilePath := "/tmp/logs/" + os.Getenv("SERVER_VERSION") + "_log.txt"
		logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			// log an error if unable to open the log file
			logger.WithError(err).Info("Failed to log to file, using default stderr")
		}
		// set the output of the logger to the file
		logger.SetOutput(logFile)
	}

	// enable caller reporting
	logger.SetReportCaller(true)

	// set the log formatter to JSON with pretty printing
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "15:04:05 02/01/2006",
		PrettyPrint:     true,
	})

	// set the log level to info
	logger.SetLevel(logrus.InfoLevel)

	return logger
}
