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
	"strconv"
	"syscall"
	"time"

	"github.com/drasko/edgex-export"
	"github.com/drasko/edgex-export/distro"
	"github.com/drasko/edgex-export/mongo"

	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
)

const (
	port                   int    = 48070
	defMongoURL            string = "0.0.0.0"
	defMongoUsername       string = "core"
	defMongoPassword       string = "password"
	defMongoDatabase       string = "coredata"
	defMongoPort           int    = 27017
	defMongoConnectTimeout int    = 5000
	defMongoSocketTimeout  int    = 60000

	envMongoURL string = "EXPORT_DISTRO_MONGO_URL"
)

type config struct {
	Port                int
	MongoURL            string
	MongoUser           string
	MongoPass           string
	MongoDatabase       string
	MongoPort           int
	MongoConnectTimeout int
	MongoSocketTimeout  int
}

type distroFormater interface {
	Format( /*event*/ ) []byte
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

var registrations []registrationInfo

type dummyFormat struct {
}

func (dummy dummyFormat) Format( /*event*/ ) []byte {
	return []byte("dummy")
}

var dummy dummyFormat

func (reg *registrationInfo) update(newReg export.Registration) bool {
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
		reg.sender = distro.NewMqttSender("tcp://127.0.0.1:1883", "", "")
	case export.DestZMQ:
		fmt.Print("Destination ZMQ is not supported")
		//reg.sender = distro.NewHttpSender("TODO URL")
	case export.DestIotCoreMQTT:
		//reg.sender = distro.NewIotCoreSender("TODO URL")
	case export.DestAzureMQTT:
		//reg.sender = distro.NewAzureSender("TODO URL")
	case export.DestRest:
		reg.sender = distro.NewHttpSender("http://127.0.0.1")
	default:
		fmt.Println("Destination not supported: ", newReg.Destination)
	}
	if reg.format == nil || reg.sender == nil {
		fmt.Println("Registration not supported")
		return false
	}
	return true
}

func getRegistrations(repo *mongo.MongoRepository) []export.Registration {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	s := repo.Session.Copy()
	defer s.Close()

	c := s.DB(mongo.DbName).C(mongo.CollectionName)

	results := []export.Registration{}
	err := c.Find(nil).All(&results)
	if err != nil {
		logger.Error("Failed to query", zap.Error(err))
		return nil
	}

	return results
}

func (reg registrationInfo) processEvent( /*, event*/ ) {
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

func sample(repo *mongo.MongoRepository) {
	sourceReg := getRegistrations(repo)

	for i := range sourceReg {
		var reg registrationInfo
		if reg.update(sourceReg[i]) {
			registrations = append(registrations, reg)
		}
	}

	for _, r := range registrations {
		fmt.Println("a registration: ", r)
		r.processEvent()
	}
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
	fmt.Println("Starting distro")
	cfg := loadConfig()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	distro.InitLogger(logger)

	ms, err := connectToMongo(cfg)
	if err != nil {
		logger.Error("Failed to connect to Mongo.", zap.Error(err))
		return
	}
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

	sample(repo)

	c := <-errs
	logger.Info("terminated", zap.String("error", c.Error()))
}

func loadConfig() *config {
	return &config{

		Port:                port,
		MongoURL:            env(envMongoURL, defMongoURL),
		MongoUser:           defMongoUsername,
		MongoPass:           defMongoPassword,
		MongoDatabase:       defMongoDatabase,
		MongoPort:           defMongoPort,
		MongoConnectTimeout: defMongoConnectTimeout,
		MongoSocketTimeout:  defMongoSocketTimeout,
	}
}

func env(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func connectToMongo(cfg *config) (*mgo.Session, error) {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{cfg.MongoURL + ":" + strconv.Itoa(cfg.MongoPort)},
		Timeout:  time.Duration(cfg.MongoConnectTimeout) * time.Millisecond,
		Database: cfg.MongoDatabase,
		Username: cfg.MongoUser,
		Password: cfg.MongoPass,
	}

	ms, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return nil, err
	}
	//logger, _ := zap.NewProduction()

	//logger.Info("--", zap.String("url", mongoDBDialInfo.Addrs[0]))

	ms.SetSocketTimeout(time.Duration(cfg.MongoSocketTimeout) * time.Millisecond)
	ms.SetMode(mgo.Monotonic, true)

	return ms, nil
}
