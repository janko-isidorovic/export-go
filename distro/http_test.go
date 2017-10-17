package distro

import (
	"fmt"
	"testing"
)

func TestHttpNew(t *testing.T) {
	sender := NewHttpSender("http://127.0.0.1")
	for i := 0; i < 1000; i++ {
		sender.Send(fmt.Sprintf("hola %d", i))
	}
	fmt.Println("Test ok")
}
