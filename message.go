//
// Copyright (c) Mainflux
//
// Mainflux server is licensed under an Apache license, version 2.0.
// All rights not explicitly granted in the Apache license, version 2.0 are reserved.
// See the included LICENSE file for more details.
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
