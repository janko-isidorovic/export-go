package distro

import (
	"bytes"
	"fmt"
	"github.com/drasko/edgex-export"
	"testing"
)

// Probably not a good test as it requires external infrastucture
func TestHttpNew(t *testing.T) {

	sender := NewHttpSender(export.Addressable{
		Name:     "test",
		Method:   export.MethodGet,
		Protocol: export.ProtoHTTP,
		Address:  "http://127.0.0.1",
		Port:     80,
		Path:     "/"})

	buf := bytes.Buffer{}
	for i := 0; i < 1000; i++ {
		buf.WriteString(fmt.Sprintf("hola %d", i))
		sender.Send(buf)
		buf.Reset()
	}

	logger.Info("Test ok")
}
