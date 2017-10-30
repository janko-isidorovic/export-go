//
// Copyright (c) 2017 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

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

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	distro.InitLogger(logger)

	logger.Info("Starting distro")
	cfg := loadConfig()

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

	distro.Loop(repo, errs)

	logger.Info("terminated")
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
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	distro.InitLogger(logger)

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{cfg.MongoURL + ":" + strconv.Itoa(cfg.MongoPort)},
		Timeout:  time.Duration(cfg.MongoConnectTimeout) * time.Millisecond,
		Database: cfg.MongoDatabase,
		Username: cfg.MongoUser,
		Password: cfg.MongoPass,
	}

	ms, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		logger.Error("Failed to connect to Mongo.", zap.Error(err))
		return nil, err
	}

	ms.SetSocketTimeout(time.Duration(cfg.MongoSocketTimeout) * time.Millisecond)
	ms.SetMode(mgo.Monotonic, true)

	return ms, nil
}
