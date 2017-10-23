package distro

import (
	"bytes"
	"github.com/drasko/edgex-export"
)

type Sender interface {
	Send(data bytes.Buffer)
}

type Formater interface {
	Format( /*event*/ ) bytes.Buffer
}

type Transformer interface {
	Transform(data bytes.Buffer) bytes.Buffer
}

type RegistrationInfo struct {
	registration export.Registration
	format       Formater
	compression  Transformer
	encrypt      Transformer
	sender       Sender
}
