//
// Copyright (c) 2017
// Cavium
//
// SPDX-License-Identifier: Apache-2.0

package distro

import (
	"github.com/drasko/edgex-export"
	"go.uber.org/zap"
)

type devIdFilterDetails struct {
	deviceIDs []string
}

func newDevIdFilter(filter export.Filter) Filterer {

	filterer := devIdFilterDetails{
		deviceIDs: filter.DeviceIDs,
	}
	return filterer
}

func (filter devIdFilterDetails) Filter(event *export.Event) (bool, *export.Event) {

	if filter.deviceIDs != nil {
		for i := range filter.deviceIDs {
			if event.Device == filter.deviceIDs[i] {
				logger.Debug("Event filtered", zap.Any("Event", event))
				return true, event
			}
		}
		return false, event
	}

	return false, event
}

type valueDescFilterDetails struct {
	valueDescIDs []string
}

func newValueDescFilter(filter export.Filter) Filterer {
	filterer := valueDescFilterDetails{
		valueDescIDs: filter.ValueDescriptorIDs,
	}
	return filterer
}

func (filter valueDescFilterDetails) Filter(event *export.Event) (bool, *export.Event) {

	if filter.valueDescIDs != nil {

		auxEvent := &export.Event{
			Pushed:   event.Pushed,
			Device:   event.Device,
			Created:  event.Created,
			Modified: event.Modified,
			Origin:   event.Origin,
			Readings: []export.Reading{},
		}

		for i := range filter.valueDescIDs {
			for j := range event.Readings {
				if event.Readings[j].Name == filter.valueDescIDs[i] {
					logger.Debug("Reading filtered", zap.Any("Reading", event.Readings[j]))
					auxEvent.Readings = append(auxEvent.Readings, event.Readings[j])
				}
			}
		}
		return true, auxEvent
	}
	// Return the event as is
	return false, event
}
