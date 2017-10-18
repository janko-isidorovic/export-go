package distro

import (
	"fmt"

	"github.com/drasko/edgex-export"
	"github.com/drasko/edgex-export/mongo"
)

var registrations []RegistrationInfo

// To be removed when any other formater is implemented
type dummyFormat struct {
}

func (dummy dummyFormat) Format( /*event*/ ) []byte {
	return []byte("dummy")
}

var dummy dummyFormat

func (reg *RegistrationInfo) update(newReg export.Registration) bool {
	reg.registration = newReg

	reg.format = nil
	switch newReg.Format {
	case export.FormatJSON:
		// reg.format = distro.NewJsonFormat()
		reg.format = dummy
	case export.FormatXML:
		// reg.format = distro.NewXmlFormat()
	case export.FormatSerialized:
		// reg.format = distro.NewSerializedFormat()
	case export.FormatIoTCoreJSON:
		// reg.format = distro.NewIotCoreFormat()
	case export.FormatAzureJSON:
		// reg.format = distro.NewAzureFormat()
	case export.FormatCSV:
		// reg.format = distro.NewCsvFormat()
	default:
		fmt.Println("Format not supported: ", newReg.Compression)
	}

	reg.compression = nil
	switch newReg.Compression {
	case export.CompNone:
		reg.compression = nil
	case export.CompGzip:
		// reg.compression = distro.NewGzipComppression()
	case export.CompZip:
		// reg.compression = distro.NewZipComppression()
	default:
		fmt.Println("Compression not supported: ", newReg.Compression)
	}

	reg.sender = nil
	switch newReg.Destination {
	case export.DestMQTT:
		reg.sender = NewMqttSender("tcp://127.0.0.1:1883", "", "")
	case export.DestZMQ:
		fmt.Print("Destination ZMQ is not supported")
		//reg.sender = distro.NewHttpSender("TODO URL")
	case export.DestIotCoreMQTT:
		//reg.sender = distro.NewIotCoreSender("TODO URL")
	case export.DestAzureMQTT:
		//reg.sender = distro.NewAzureSender("TODO URL")
	case export.DestRest:
		reg.sender = NewHttpSender("http://127.0.0.1")
	default:
		fmt.Println("Destination not supported: ", newReg.Destination)
	}
	if reg.format == nil || reg.sender == nil {
		fmt.Println("Registration not supported")
		return false
	}
	return true
}

func (reg RegistrationInfo) processEvent( /*, event*/ ) {
	// Valid Event Filter, needed?

	// Device filtering TODO

	// Value filtering TODO

	//formated := reg.format.Format( /* event*/ )
	formated := []byte("just an example")
	compressed := formated
	if reg.compression != nil {
		compressed = reg.compression.Transform(formated)
	}

	encrypted := compressed
	if reg.encrypt != nil {
		encrypted = reg.encrypt.Transform(compressed)
	}

	reg.sender.Send(string(encrypted))
}

func TestDistro(repo *mongo.MongoRepository) {
	sourceReg := getRegistrations(repo)

	for i := range sourceReg {
		var reg RegistrationInfo
		if reg.update(sourceReg[i]) {
			registrations = append(registrations, reg)
		}
	}

	for _, r := range registrations {
		fmt.Println("a registration: ", r)
		r.processEvent()
	}
}
