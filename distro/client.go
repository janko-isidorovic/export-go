//
// Copyright (c) 2017
// Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/drasko/edgex-export"
	"go.uber.org/zap"
)

const (
	// TODO this consts need to be configurable somehow
	clientHost     = "127.0.0.1"
	clientPort int = 48071
)

func getRegistrations() []export.Registration {
	url := "http://" + clientHost + ":" + strconv.Itoa(clientPort) +
		"/api/v1/registration"

	response, err := http.Get(url)
	if err != nil {
		logger.Warn("Error getting all registrations", zap.String("url", url))
		return nil
	}
	defer response.Body.Close()

	results := []export.Registration{}
	if err := json.NewDecoder(response.Body).Decode(&results); err != nil {
		logger.Warn("Could not parse json", zap.Error(err))
	}

	return results
}

func getRegistrationByName(name string) *export.Registration {
	url := "http://" + clientHost + ":" + strconv.Itoa(clientPort) +
		"/api/v1/registration/name/" + name

	response, err := http.Get(url)
	if err != nil {
		logger.Error("Error getting all registrations", zap.String("url", url))
		return nil
	}
	defer response.Body.Close()

	reg := export.Registration{}
	if err := json.NewDecoder(response.Body).Decode(&reg); err != nil {
		logger.Error("Could not parse json", zap.Error(err))
	}

	return &reg
}
