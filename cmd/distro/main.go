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

	distro.TestDistro(repo)

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

	ms.SetSocketTimeout(time.Duration(cfg.MongoSocketTimeout) * time.Millisecond)
	ms.SetMode(mgo.Monotonic, true)

	return ms, nil
}
