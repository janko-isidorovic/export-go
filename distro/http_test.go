//
// Copyright (c) 2017 Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"fmt"
	"github.com/drasko/edgex-export"
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
	"testing"
)

var log *zap.Logger

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == export.MethodGet {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Error("err", zap.Error(err))
		}
		fmt.Println(string(requestDump))
	}

	if r.Method == export.MethodPost {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Fatal("err", zap.Error(err))
		}
		fmt.Println(string(requestDump))
	}

	w.WriteHeader(http.StatusOK)
}

func RunServer() {
	http.HandleFunc("/", handler)            // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Error("ListenAndServe: ", zap.Error(err))
	}
}

// Probably not a good test as it requires external infrastucture
func TestHttpNew(t *testing.T) {
	log, _ = zap.NewProduction()
	defer log.Sync()

	InitLogger(log)

	go RunServer()

	sender := NewHttpSender(export.Addressable{
		Name:     "test",
		Method:   export.MethodGet,
		Protocol: export.ProtoHTTP,
		Address:  "http://127.0.0.1",
		Port:     9090,
		Path:     "/"})

	for i := 0; i < 10; i++ {
		sender.Send(fmt.Sprintf("hola %d", i))
	}

	log.Info("Test ok")

	senderPost := NewHttpSender(export.Addressable{
		Name:     "test",
		Method:   export.MethodPost,
		Protocol: export.ProtoHTTP,
		Address:  "http://127.0.0.1",
		Port:     9090,
		Path:     "/"})

	for i := 0; i < 10; i++ {
		senderPost.Send(fmt.Sprintf("hola %d", i))
	}

	log.Info("Test ok")
}
