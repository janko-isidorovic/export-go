//
// Copyright (c) 2017 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"encoding/json"
	"fmt"
	"github.com/drasko/edgex-export"
	"go.uber.org/zap"
	"net"
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
		log.Info("Dump", zap.ByteString("Dump", requestDump))
		w.WriteHeader(http.StatusOK)
	}

	if r.Method == export.MethodPost {
		w.Header().Set("Content-Type", "application/json")

		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Fatal("err", zap.Error(err))
		}
		log.Info("Dump", zap.ByteString("Dump", requestDump))

		var jsonT string
		err = json.NewDecoder(r.Body).Decode(&jsonT)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		log.Info("JSON", zap.String("JSON", fmt.Sprintf("%#v", jsonT)))
	}

}

// Probably not a good test as it requires external infrastucture
func TestHttpNew(t *testing.T) {
	log, _ = zap.NewProduction()
	defer log.Sync()

	InitLogger(log)

	http.HandleFunc("/", handler)

	ln, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Error("Can't listen: %s", zap.Error(err))
	}

	go http.Serve(ln, nil)

	//	defer ln.Close()

	senderHttp := NewHttpSender(export.Addressable{
		Name:     "test",
		Method:   export.MethodGet,
		Protocol: export.ProtoHTTP,
		Address:  "http://127.0.0.1",
		Port:     9090,
		Path:     "/"})
	senderHttp.Send("dummy")

	log.Info("Test ok")

	senderPost := NewHttpSender(export.Addressable{
		Name:     "test",
		Method:   export.MethodPost,
		Protocol: export.ProtoHTTP,
		Address:  "http://127.0.0.1",
		Port:     9090,
		Path:     "/"})

	senderPost.Send("{\"key\": \"Hello, \", \"value\": \"World!\"}")

	log.Info("Test ok")
}
