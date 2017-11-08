//
// Copyright (c) 2017
// Cavium
//
// SPDX-License-Identifier: Apache-2.0

package distro

import (
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

				_, auxEvent = filterByValueDescriptor(filter, event)
				return true, auxEvent
			}
		}
		return false, auxEvent
	}

	return filterByValueDescriptor(filter, event)
}

type valueDescFilterer struct {
	valueDesc []string
}

func filterByValueDescriptor(filter filterDetails, event *export.Event) (bool, *export.Event) {
	auxEvent := &export.Event{}

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

		for i := range filter.valueDesc {
			for j := range event.Readings {
				if event.Readings[j].Name == filter.valueDesc[i] {
					fmt.Println("Filtering by value descriptor id: ", filter.valueDesc[i])
					fmt.Println("Reading ", j)
					auxEvent.Readings = append(auxEvent.Readings, event.Readings[j])
				}
			}
		}
		return true, auxEvent
	}
	// Return the event as is
	return false, event
}
