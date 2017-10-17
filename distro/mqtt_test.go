package distro

import (
	// "io/ioutil"
	// "net/http"
	"fmt"
	"testing"
)

func TestMqttNew(t *testing.T) {
	sender := NewMqttSender("tcp://127.0.0.1:1883", "", "")
	for i := 0; i < 1000; i++ {

		sender.Send(fmt.Sprintf("hola %d", i))
	}
	fmt.Println("Test ok")
}
