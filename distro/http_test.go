//
// Copyright (c) 2017 Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
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

	for i := 0; i < 1000; i++ {
		sender.Send(fmt.Sprintf("hola %d", i))
	}
  
	logger.Info("Test ok")
}