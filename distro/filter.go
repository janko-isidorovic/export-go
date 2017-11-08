//
// Copyright (c) 2017
// Cavium
//
// SPDX-License-Identifier: Apache-2.0

package distro

import (
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
					auxEvent.Readings = append(auxEvent.Readings, event.Readings[j])
				}
			}
		}
		return true, auxEvent
	}
	// Return the event as is
	return false, event
}
