package distro

import (
	"bytes"
	"fmt"

	"github.com/drasko/edgex-export"
	"github.com/drasko/edgex-export/mongo"
	"go.uber.org/zap"
)

var registrations []RegistrationInfo

// To be removed when any other formater is implemented
type dummyFormat struct {
}

func (dummy dummyFormat) Format( /*event*/ ) bytes.Buffer {
	buf := bytes.Buffer{}
	buf.WriteString("dummy")
	return buf
}

var dummy dummyFormat

func (reg *RegistrationInfo) update(newReg export.Registration) bool {
	reg.registration = newReg

	reg.format = nil
	switch newReg.Format {
	case export.FormatJSON:
		// TODO reg.format = distro.NewJsonFormat()
		reg.format = dummy
	case export.FormatXML:
		// TODO reg.format = distro.NewXmlFormat()
		reg.format = dummy
	case export.FormatSerialized:
		// TODO reg.format = distro.NewSerializedFormat()
	case export.FormatIoTCoreJSON:
		// TODO reg.format = distro.NewIotCoreFormat()
	case export.FormatAzureJSON:
		// TODO reg.format = distro.NewAzureFormat()
	case export.FormatCSV:
		// TODO reg.format = distro.NewCsvFormat()
	default:
		logger.Info("Format not supported: ", zap.String("format", newReg.Format))
	}

	reg.compression = nil
	switch newReg.Compression {
	case export.CompNone:
		reg.compression = nil
	case export.CompGzip:
		reg.compression = gzipTransformer{}
	case export.CompZip:
		// TODO reg.compression = distro.NewZipComppression()
	default:
		logger.Info("Compression not supported: ", zap.String("compression", newReg.Compression))
	}

	reg.sender = nil
	switch newReg.Destination {
	case export.DestMQTT:
		reg.sender = NewMqttSender(newReg.Addressable)
	case export.DestZMQ:
		logger.Info("Destination ZMQ is not supported")
	case export.DestIotCoreMQTT:
		// TODO reg.sender = distro.NewIotCoreSender("TODO URL")
	case export.DestAzureMQTT:
		// TODO reg.sender = distro.NewAzureSender("TODO URL")
	case export.DestRest:
		reg.sender = NewHttpSender(newReg.Addressable)
	default:
		logger.Info("Destination not supported: ", zap.String("destination", newReg.Destination))
	}
	if reg.format == nil || reg.sender == nil {
		logger.Error("Registration not supported")
		return false
	}
	return true
}

func (reg RegistrationInfo) processEvent( /*, event*/ ) {
	// Valid Event Filter, needed?

	// TODO Device filtering

	// TODO Value filtering
	formated := reg.format.Format( /* event*/ )
	compressed := formated
	if reg.compression != nil {
		compressed = reg.compression.Transform(formated)
	}

	encrypted := compressed
	if reg.encrypt != nil {
		encrypted = reg.encrypt.Transform(compressed)
	}

	reg.sender.Send(encrypted)
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
		logger.Info("a registration:", zap.String("reg", fmt.Sprintf("%#v", r)))
		r.processEvent()
	}
}
