package logger

import "github.com/Sirupsen/logrus"

var Log = logrus.New()
var LogLevelAsString = "Info"

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
