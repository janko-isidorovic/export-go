//
// Copyright (c) 2017 Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"bytes"
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

	buf := bytes.Buffer{}
	for i := 0; i < 1000; i++ {
		buf.WriteString(fmt.Sprintf("hola %d", i))

		sender.Send(buf)
	}
	logger.Info("Test ok")
}
