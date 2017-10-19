package distro

import (
	"fmt"
	"github.com/drasko/edgex-export"
	"testing"
)

func TestHttpNew(t *testing.T) {
	addr := export.Addressable{
		Name:     "test",
		Method:   "GET",
		Protocol: "HTTP",
		Address:  "http://127.0.0.1",
		Port:     80,
		Path:     "/"}

	sender := NewHttpSender(addr)
	for i := 0; i < 1000; i++ {
		sender.Send(fmt.Sprintf("hola %d", i))
	}
	logger.Info("Test ok")
}
