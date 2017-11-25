//
// Copyright (c) 2017
// Cavium
// Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/drasko/edgex-export"
	"github.com/drasko/edgex-export/distro"

	"go.uber.org/zap"
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

var logger *zap.Logger

func main() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	distro.InitLogger(logger)

	logger.Info("Starting distro")
	cfg := loadConfig()

	errs := make(chan error, 2)
	eventCh := make(chan *export.Event, 10)

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

	// There can be another receivers that can be initialiced here
	distro.ZeroMQReceiver(eventCh)

	distro.Loop(errs, eventCh)

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
