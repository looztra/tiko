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

package tiko

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/looztra/tiko/logger"
)

// Resolve is the core url resolution engine
func Resolve(url string) (string, error) {
	logger.Log.Debug("in 'Resolver.Resolve', received " + url)

	resp, err := http.Head(url)
	if err != nil {
		return "", err
	}
	if logger.Log.Level == logrus.DebugLevel {
		logger.Log.Debug("Status code : " + resp.Status)
		logger.Log.Debug("Request URL : " + resp.Request.URL.String())
		for name, value := range resp.Header {
			logger.Log.Debug("resp.Header >> name: " + name + ", value:" + value[0])
		}
		for name, value := range resp.Request.Header {
			logger.Log.Debug("req.Header >> name: " + name + ", value:" + value[0])
		}
	}
	return resp.Request.URL.String(), nil
}
