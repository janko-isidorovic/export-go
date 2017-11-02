//
// Copyright (c) 2017
// Cavium
//
// SPDX-License-Identifier: Apache-2.0

package distro

import (
	"encoding/json"

	"github.com/drasko/edgex-export"
	"go.uber.org/zap"
)

type jsonFormater struct {
}

func (jsonTr jsonFormater) Format(event *export.Event) []byte {

	b, err := json.Marshal(event)
	if err != nil {
		logger.Error("Error parsing JSON", zap.Error(err))
		return nil
	}
	return b
}

type xmlFormater struct {
	xml []string
}

func (xmlTr xmlFormater) Format(event *export.Event) []byte {
	return []byte("dummy")
}
