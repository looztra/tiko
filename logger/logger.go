package logger

import "github.com/Sirupsen/logrus"

// Log is the main log entrypoint
var Log = logrus.New()

// LogLevelAsString is the string representation of the current log level
var LogLevelAsString = "Info"

// InitLogger inits the log engine
func InitLogger() {

	level, err := logrus.ParseLevel(LogLevelAsString)
	if err != nil {
		Log.Info("Invalid value <" + LogLevelAsString + "> for the log level, falling back to 'Info'")
		Log.Level = logrus.InfoLevel
	} else {
		Log.Level = level
	}
	Log.Info("Log Level set to " + Log.Level.String())
}
