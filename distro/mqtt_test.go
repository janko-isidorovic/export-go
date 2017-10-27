//
// Copyright (c) Mainflux
//
// Mainflux server is licensed under an Apache license, version 2.0.
// All rights not explicitly granted in the Apache license, version 2.0 are reserved.
// See the included LICENSE file for more details.
//

package distro

import (
	"fmt"
	"github.com/drasko/edgex-export"
	"testing"
)

// Probably not a good test as it requires external infrastucture
func TestMqttNew(t *testing.T) {
	sender := NewMqttSender(export.Addressable{
		Address: "tcp://127.0.0.1",
		Port:    1883,
	})

	for i := 0; i < 1000; i++ {

		sender.Send(fmt.Sprintf("hola %d", i))
	}
	logger.Info("Test ok")
}
