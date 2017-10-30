//
// Copyright (c) 2017 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"github.com/drasko/edgex-export"
)

type Sender interface {
	Send(data []byte)
}

type Formater interface {
	Format( /* FIXME event*/ ) []byte
}

type Transformer interface {
	Transform(data []byte) []byte
}

type RegistrationInfo struct {
	registration export.Registration
	format       Formater
	compression  Transformer
	encrypt      Transformer
	sender       Sender

	chRegistration chan *RegistrationInfo

	// TODO To be changed to event
	chEvent chan bool
}
