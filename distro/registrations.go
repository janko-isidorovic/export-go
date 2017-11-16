//
// Copyright (c) 2017
// Cavium
// Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

// TODO:
// - Filtering by id and value
// - Receive events from 0mq until a new message broker/rpc is chosen
// - Event buffer management per sender(do not block distro.Loop on full
//   registration channel)

import (
	"time"

	"github.com/drasko/edgex-export"
	"go.uber.org/zap"
)

var registrationChanges chan export.NotifyUpdate = make(chan export.NotifyUpdate, 2)

func RefreshRegistrations(update export.NotifyUpdate) {
	// TODO make it not blocking, return bool?
	registrationChanges <- update
}

func newRegistrationInfo() *RegistrationInfo {
	reg := &RegistrationInfo{}

	reg.chRegistration = make(chan *export.Registration)
	reg.chEvent = make(chan *export.Event)
	return reg
}

func (reg *RegistrationInfo) update(newReg export.Registration) bool {
	reg.registration = newReg

	reg.format = nil
	switch newReg.Format {
	case export.FormatJSON:
		reg.format = jsonFormater{}
	case export.FormatXML:
		reg.format = xmlFormater{}
	case export.FormatSerialized:
		// TODO reg.format = distro.NewSerializedFormat()
	case export.FormatIoTCoreJSON:
		// TODO reg.format = distro.NewIotCoreFormat()
	case export.FormatAzureJSON:
		// TODO reg.format = distro.NewAzureFormat()
	case export.FormatCSV:
		// TODO reg.format = distro.NewCsvFormat()
	default:
		logger.Info("Format not supported: ", zap.String("format", newReg.Format))
	}

	reg.compression = nil
	switch newReg.Compression {
	case export.CompNone:
		reg.compression = nil
	case export.CompGzip:
		reg.compression = gzipTransformer{}
	case export.CompZip:
		reg.compression = zlibTransformer{}
	default:
		logger.Info("Compression not supported: ", zap.String("compression", newReg.Compression))
	}

	reg.sender = nil
	switch newReg.Destination {
	case export.DestMQTT:
		reg.sender = NewMqttSender(newReg.Addressable)
	case export.DestZMQ:
		logger.Info("Destination ZMQ is not supported")
	case export.DestIotCoreMQTT:
		// TODO reg.sender = distro.NewIotCoreSender("TODO URL")
	case export.DestAzureMQTT:
		// TODO reg.sender = distro.NewAzureSender("TODO URL")
	case export.DestRest:
		reg.sender = NewHTTPSender(newReg.Addressable)
	default:
		logger.Info("Destination not supported: ", zap.String("destination", newReg.Destination))
	}
	if reg.format == nil || reg.sender == nil {
		logger.Error("Registration not supported")
		return false
	}

	reg.encrypt = nil
	switch newReg.Encryption.Algo {
	case export.EncNone:
		reg.encrypt = nil
	case export.EncAes:
		reg.encrypt = NewAESEncryption(newReg.Encryption)
	default:
		logger.Info("Encryption not supported: ", zap.String("Algorithm", newReg.Encryption.Algo))

	}

	reg.filter = nil

	if len(newReg.Filter.DeviceIDs) > 0 {
		reg.filter = append(reg.filter, newDevIdFilter(newReg.Filter))
		logger.Info("Device ID filter added: ", zap.Any("filters", newReg.Filter.DeviceIDs))
	}

	if len(newReg.Filter.ValueDescriptorIDs) > 0 {
		reg.filter = append(reg.filter, newValueDescFilter(newReg.Filter))
		logger.Info("Value descriptor filter added: ", zap.Any("filters", newReg.Filter.ValueDescriptorIDs))
	}

	return true
}

func (reg RegistrationInfo) processEvent(event *export.Event) {
	// Valid Event Filter, needed?

	var filtered bool
	for i := range reg.filter {
		filtered, event = reg.filter[i].Filter(event)
		if !filtered {
			logger.Info("Event filtered")
			return
		}
	}

	formated := reg.format.Format(event)

	compressed := formated
	if reg.compression != nil {
		compressed = reg.compression.Transform(formated)
	}

	encrypted := compressed
	if reg.encrypt != nil {
		encrypted = reg.encrypt.Transform(compressed)
	}

	reg.sender.Send(encrypted)
	logger.Debug("Sent event with registration:",
		zap.Any("Event", event),
		zap.String("Name", reg.registration.Name))
}

func registrationLoop(reg *RegistrationInfo) {
	logger.Info("registration loop started",
		zap.String("Name", reg.registration.Name))
	for {
		select {
		case event := <-reg.chEvent:
			reg.processEvent(event)

		case newReg := <-reg.chRegistration:
			if newReg == nil {
				logger.Info("Terminating registration goroutine")
				return
			} else {
				if reg.update(*newReg) {
					logger.Info("Registration updated: OK",
						zap.String("Name", reg.registration.Name))
				} else {
					logger.Info("Registration updated: KO, terminating goroutine",
						zap.String("Name", reg.registration.Name))
					reg.deleteMe = true
					return
				}
			}
		}
	}
}

func updateRunningRegistrations(running map[string]*RegistrationInfo,
	update export.NotifyUpdate) {

	switch update.Operation {
	case export.NotifyUpdateDelete:
		for k, v := range running {
			if k == update.Name {
				v.chRegistration <- nil
				delete(running, k)
				return
			}
		}
		logger.Warn("delete update not processed")
	case export.NotifyUpdateUpdate:
		reg := getRegistrationByName(update.Name)
		if reg == nil {
			logger.Error("Could not find registration", zap.String("name", update.Name))
			return
		}
		for k, v := range running {
			if k == update.Name {
				v.chRegistration <- reg
				return
			}
		}
		logger.Error("Could not find running registration", zap.String("name", update.Name))
	case export.NotifyUpdateAdd:
		reg := getRegistrationByName(update.Name)
		if reg == nil {
			logger.Error("Could not find registration", zap.String("name", update.Name))
			return
		}
		regInfo := newRegistrationInfo()
		if regInfo.update(*reg) {
			running[reg.Name] = regInfo
			go registrationLoop(regInfo)
		}
	default:
		logger.Error("Invalid update operation", zap.String("operation", update.Operation))
	}
}

// Loop - registration loop
func Loop(errChan chan error) {

	registrations := make(map[string]*RegistrationInfo)

	// Create new goroutines for each registration
	for _, reg := range getRegistrations() {
		regInfo := newRegistrationInfo()
		if regInfo.update(reg) {
			registrations[reg.Name] = regInfo
			go registrationLoop(regInfo)
		}
	}

	logger.Info("Starting registration loop")
	for {
		select {
		case e := <-errChan:
			// kill all registration goroutines
			for k, reg := range registrations {
				if !reg.deleteMe {
					// Do not write in channel that will not be read
					reg.chRegistration <- nil
				}
				delete(registrations, k)
			}
			logger.Info("exit msg", zap.Error(e))
			return

		case update := <-registrationChanges:
			logger.Info("Registration changes")
			updateRunningRegistrations(registrations, update)

		case <-time.After(time.Second):
			// Simulate receiving events
			event := getNextEvent()

			for k, reg := range registrations {
				if reg.deleteMe {
					delete(registrations, k)
				} else {
					// TODO only sent event if it is not blocking
					reg.chEvent <- event
				}
			}
		}
	}
}
