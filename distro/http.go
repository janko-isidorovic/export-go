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
			logger.Error("Error: ", zap.Error(err))
			return
		}
		defer response.Body.Close()
		logger.Info("Response: ", zap.String("status", response.Status))

	case export.MethodPost:
		response, err := http.Post(sender.url, mimeTypeJSON, strings.NewReader(data))
		if err != nil {
			logger.Error("Error: ", zap.Error(err))
			return
		}
		defer response.Body.Close()
		logger.Info("Response: ", zap.String("status", response.Status))
	default:
		logger.Info("Unsupported method: ", zap.String("method", sender.method))
	}

	logger.Info("Sent data: ", zap.String("data", data))
}
