//
// Copyright (c) 2017 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//
package distro

import (
	"encoding/json"

	"github.com/drasko/edgex-export"
	"go.uber.org/zap"
)

//const sampleEvent string = `{"pushed":0,"device":"GS1-AC-Drive01","readings":[{"pushed":0,"name":"HoldingRegister_8455","value":"287.27","device":"GS1-AC-Drive01","id":"59f70666e4b0e3fab1d4232e","created":1509361254069,"modified":1509361254069,"origin":1509361254001}],"id":"59f70666e4b0e3fab1d4232f","created":1509361254072,"modified":1509361254072,"origin":1509361254001}`

const sampleEvent string = `{"pushed":0,"device":"livingroomthermostat","readings":[{"pushed":0,"name":"temperature","value":"72","id":"57ed24f0502fdf73bb637915","created":1475159280744,"modified":1475159280744,"origin":1471806386919},{"pushed":0,"name":"humidity","value":"58","id":"57ed24f0502fdf73bb637916","created":1475159280756,"modified":1475159280756,"origin":1471806386919}],"id":"57ed24f0502fdf73bb637917","created":1475159280762,"modified":1475159280762,"origin":1471806386919}`

func getNextEvent() *export.Event {
	event := export.Event{}
	if err := json.Unmarshal([]byte(sampleEvent), &event); err != nil {
		logger.Error("Failed to query add registration", zap.Error(err))
		return nil
	}

	return &event
}
