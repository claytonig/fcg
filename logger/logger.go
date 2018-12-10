package logger

import "github.com/sirupsen/logrus"

// log - logger instance
var log *logrus.Logger

// Log - Function to get logger instance
func Log() *logrus.Logger {
	if log != nil {
		return log
	}
	log = logrus.New()
	return log
}
