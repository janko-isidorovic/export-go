//
// Copyright (c) 2017 Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package export

// Message - Encapsulating / wrapper message object that contains Event
// to be exported and the client export registration details
type Message struct {
	Registration Registration
	Evt          Event
}

// Event - packet of Readings
type Event struct {
	Pushed   int64
	Device   string
	Readings []Reading
}

// Reading - Sensor measurement
type Reading struct {
	Pushed int64
	Name   string
	Value  string
	Device string
}
