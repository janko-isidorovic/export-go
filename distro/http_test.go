package distro

import (
	"net"
	"net/http"
	"net/http/httputil"
	"testing"

	"github.com/drasko/edgex-export"
	"go.uber.org/zap"
)

var log *zap.Logger

func handlerGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == export.MethodGet {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Error("err", zap.Error(err))
		}
		log.Info("Dump", zap.ByteString("Dump", requestDump))
	}
	w.WriteHeader(http.StatusOK)
}

func handlerPost(w http.ResponseWriter, r *http.Request) {

	if r.Method == export.MethodPost {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Fatal("err", zap.Error(err))
		}
		log.Info("Dump", zap.ByteString("Dump", requestDump))

	}

	w.WriteHeader(http.StatusOK)

}

func TestHttpNew(t *testing.T) {
	log, _ = zap.NewProduction()
	defer log.Sync()

	InitLogger(log)

	http.HandleFunc("/GetTest", handlerGet)
	http.HandleFunc("/PostTest", handlerPost)

	ln, err := net.Listen("tcp", "127.0.0.1:9090")

	if err != nil {
		log.Error("Can't listen: %s", zap.Error(err))
	}

	go http.Serve(ln, nil)

	defer ln.Close()

	senderHttp := NewHttpSender(export.Addressable{
		Name:     "test",
		Method:   export.MethodGet,
		Protocol: export.ProtoHTTP,
		Address:  "http://127.0.0.1",
		Port:     9090,
		Path:     "/GetTest"})

	dataToSend := []byte("dummy")
	senderHttp.Send(dataToSend)

	log.Info("Test ok")

	senderPost := NewHttpSender(export.Addressable{
		Name:     "test",
		Method:   export.MethodPost,
		Protocol: export.ProtoHTTP,
		Address:  "http://127.0.0.1",
		Port:     9090,
		Path:     "/PostTest"})

	senderPost.Send([]byte("{\"key\": \"Hello, \", \"value\": \"World!\"}"))

	log.Info("Test ok")
}
