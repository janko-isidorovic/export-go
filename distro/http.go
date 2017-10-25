//
// Copyright (c) 2017 Cavium 
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"github.com/drasko/edgex-export"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

type httpSender struct {
	url    string
	method string
}

const mimeTypeJSON = "application/json"

func NewHttpSender(addr export.Addressable) Sender {
	// CHN: Should be added protocol from Addressable instead of include it the address param.
	// CHN: We will maintain this behaviour for compatibility with Java
	sender := httpSender{
		url:    addr.Address + ":" + strconv.Itoa(addr.Port) + addr.Path,
		method: addr.Method,
	}
	return sender
}

func (sender httpSender) Send(data string) {
	switch sender.method {

	case export.MethodGet:
		response, err := http.Get(sender.url)
		if err != nil {
			//FIXME
			logger.Error("Error: ", zap.Error(err))
			return
		}
		defer response.Body.Close()
		logger.Info("Response: ", zap.String("status", response.Status))

	case export.MethodPost:
		var buf string
		response, err := http.Post(sender.url, mimeTypeJSON, nil)
		if err != nil {
			logger.Error("Error: ", zap.Error(err))
			return
		}
		defer response.Body.Close()
		logger.Info("Response: ", zap.String("status", response.Status))
		logger.Info("Buf: ", zap.String("buf", buf))

 zap.String("method", sender.method))
	default:
		logger.Info("Unsupported method: ", zap.String("method", sender.method))
	}

	logger.Info("Sent data: ", zap.String("data", data))
}
