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

/*
type Filter struct {
	DeviceIDs          []string `json:"deviceIdentifiers,omitempty"`
	ValueDescriptorIDs []string `json:"valueDescriptorIdentifiers,omitempty"`
}


*/

func FilterbyDeviceID(event *export.Event) bool {

	fmt.Println("Filterin by Device id")

	return true
}
