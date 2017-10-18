package distro

import (
	// "io/ioutil"
	// "net/http"
	"fmt"
	"github.com/drasko/edgex-export"
	"testing"
)

func TestMqttNew(t *testing.T) {
	sender := NewMqttSender(export.Addressable{
		Address: "tcp://127.0.0.1",
		Port:    1883,
	})

	for i := 0; i < 1000; i++ {

		sender.Send(fmt.Sprintf("hola %d", i))
	}
	fmt.Println("Test ok")
}
