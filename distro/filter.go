//
// Copyright (c) 2017
// Cavium
//
// SPDX-License-Identifier: Apache-2.0

package distro

import (
	//"encoding/json"
	"fmt"

	"github.com/drasko/edgex-export"
)

type divIdFilterer struct {
	deviceIDs []string
}

func NewDeviceIDFilter(deviceIDsExt []string) Filterer {
	devIdFilt := divIdFilterer{
		deviceIDs: deviceIDsExt,
	}
	return devIdFilt
}

func (devIdFilt divIdFilterer) Filter(event *export.Event) bool {

	for i := range devIdFilt.deviceIDs {
		if event.Device == devIdFilt.deviceIDs[i] {
			fmt.Println("Filtering by Device id: ", devIdFilt)
			return true
		}
	}

	return false
}

type valueDescFilterer struct {
	valueDesc []string
}

func NewValueDescFilter(valueDescExt []string) Filterer {
	valueDescFilt := valueDescFilterer{
		valueDesc: valueDescExt,
	}
	return valueDescFilt
}

func (valueDescFilt valueDescFilterer) Filter(event *export.Event) bool {

	for i := range valueDescFilt.valueDesc {
		if event.Device == valueDescFilt.valueDesc[i] {
			fmt.Println("Filtering by value descriptor id: ", valueDescFilt)
			return true
		}
	}

	return false
}
