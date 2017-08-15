package tiko

import (
	"github.com/Sirupsen/logrus"
	"github.com/looztra/tiko/logger"
	"net/http"
)

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
