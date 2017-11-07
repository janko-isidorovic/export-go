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

type filterDetails struct {
	deviceIDs []string
	valueDesc []string
}

func applyFilters(filter export.Filter) Filterer {
	filterer := filterDetails{}

	if len(filter.DeviceIDs) > 0 {
		filterer.deviceIDs = filter.DeviceIDs
	}

	if len(filter.ValueDescriptorIDs) > 0 {
		filterer.valueDesc = filter.ValueDescriptorIDs
	}

	return filterer
}

func (filter filterDetails) Filter(event *export.Event) (bool, *export.Event) {

	auxEvent := &export.Event{}

	if filter.deviceIDs != nil {

		for i := range filter.deviceIDs {
			if event.Device == filter.deviceIDs[i] {
				fmt.Println("Filtering by Device id: ", filter)
				return true, auxEvent
			}
		}

		return false, auxEvent
	}

	if filter.valueDesc != nil {
		fmt.Println("lens: ", len(filter.valueDesc), " lens ", len(event.Readings))
		auxEvent = &export.Event{
			Pushed:   event.Pushed,
			Device:   event.Device,
			Created:  event.Created,
			Modified: event.Modified,
			Origin:   event.Origin,
			Readings: []export.Reading{},
		}
		/*
			fmt.Println("---------------------------------------------------")
			fmt.Println(auxEvent)
			fmt.Println("---------------------------------------------------")
		*/
		for i := range filter.valueDesc {
			for j := range event.Readings {
				if event.Readings[j].Name == filter.valueDesc[i] {
					fmt.Println("Filtering by value descriptor id: ", filter.valueDesc[i])
					fmt.Println("Reading ", j)
					auxEvent.Readings = append(auxEvent.Readings, event.Readings[j])
				}
			}
		}

		/*	fmt.Println("---------------------------------------------------")
			fmt.Println(auxEvent)
			fmt.Println("---------------------------------------------------")
		*/
		return true, auxEvent
	}
	return false, auxEvent
}
