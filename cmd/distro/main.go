/**
 * Copyright (c) 2017 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/drasko/edgex-export"
	"github.com/drasko/edgex-export/distro"
	"github.com/drasko/edgex-export/mongo"

	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
)

const (
	port        int    = 48070
	defMongoURL string = "mongodb://0.0.0.0:27017"
	envMongoURL string = "EXPORT_DISTRO_MONGO_URL"
)

type config struct {
	Port     int
	MongoURL string
}

var registrations []export.Registration

type distroFormater interface {
	Format( /*event*/ ) string
}

type distroTransformer interface {
	Transform(data []byte) []byte
}

type registrationInfo struct {
	registration export.Registration
	format       distroFormater
	compression  distroTransformer
	encrypt      distroTransformer
	sender       distro.Sender
}

func (reg registrationInfo) update(newReg export.Registration) bool {
	reg.registration = newReg

	reg.format = nil
	switch newReg.Format {
	case export.FormatJSON:
		// reg.format = distro.NewJsonFormat()
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
		reg.sender = distro.NewMqttSender("broker", "user", "password")
	case export.DestZMQ:
		fmt.Print("Destination ZMQ is not supported")
		//reg.sender = distro.NewHttpSender("TODO URL")
	case export.DestIotCoreMQTT:
		//reg.sender = distro.NewIotCoreSender("TODO URL")
	case export.DestAzureMQTT:
		//reg.sender = distro.NewAzureSender("TODO URL")
	case export.DestRest:
		reg.sender = distro.NewHttpSender("TODO URL")
	default:
		fmt.Println("Destination not supported: ", newReg.Destination)
	}
	if reg.format == nil || reg.sender == nil {
		fmt.Println("Registration not supported")
		return false
	}
	return true
}

func sample() {
	var sourceReg export.Registration
	sourceReg.ID = "1"
	sourceReg.Name = "test export"
	// sourceReg.Addr
	sourceReg.Format = export.FormatJSON
	//sourceReg.Filter
	sourceReg.Compression = export.CompNone
	sourceReg.Encryption.Algo = export.EncNone
	sourceReg.Enable = true
	sourceReg.Destination = export.DestMQTT

	var reg registrationInfo
	reg.update(sourceReg)
}

// {
// 	"_id" : ObjectId("59e4bdb9e4b091c2db2054c8"),
// 	"_class" : "org.edgexfoundry.domain.export.ExportRegistration",
// 	"name" : "MQTTClient2",
// 	"addressable" : {
// 		"_id" : null,
// 		"name" : "FuseTestMQTTBroker2",
// 		"protocol" : "TCP",
// 		"address" : "tcp://127.0.0.1",
// 		"port" : 1883,
// 		"publisher" : "FuseExportPublisher",
// 		"user" : "dummy",
// 		"password" : "dummy",
// 		"topic" : "FuseDataTopic",
// 		"created" : NumberLong(0),
// 		"modified" : NumberLong(0),
// 		"origin" : NumberLong("1471806386919")
// 	},
// 	"format" : "JSON",
// 	"compression" : "NONE",
// 	"enable" : true,
// 	"destination" : "MQTT_TOPIC",
// 	"created" : NumberLong("1508163001830"),
// 	"modified" : NumberLong("1508163001830"),
// 	"origin" : NumberLong("1471806386919")
// }

func main() {
	cfg := loadConfig()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	distro.InitLogger(logger)

	ms := connectToMongo(cfg, logger)
	defer ms.Close()

	repo := mongo.NewMongoRepository(ms)
	distro.InitMongoRepository(repo)

	errs := make(chan error, 2)

	go func() {
		p := fmt.Sprintf(":%d", cfg.Port)
		logger.Info("Starting Export Distro", zap.String("url", p))
		errs <- http.ListenAndServe(p, distro.HTTPServer())
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	sample()

	c := <-errs
	logger.Info("terminated", zap.String("error", c.Error()))
}

func loadConfig() *config {
	return &config{
		Port:     port,
		MongoURL: env(envMongoURL, defMongoURL),
	}
}

func env(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func connectToMongo(cfg *config, logger *zap.Logger) *mgo.Session {
	ms, err := mgo.Dial(cfg.MongoURL)
	if err != nil {
		logger.Error("Failed to connect to Mongo.", zap.Error(err))
	}

	ms.SetMode(mgo.Monotonic, true)

	return ms
}
