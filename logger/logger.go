// Copyright © 2017 Christophe Furmaniak <christophe point furmaniak chez gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
