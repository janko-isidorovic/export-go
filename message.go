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
	Pushed   int64     `json:"pushed,omitempty"`
	Device   string    `json:"device,omitempty"`
	Readings []Reading `json:"readings,omitempty"`
	Created  int64     `json:"created,omitempty"`
	Modified int64     `json:"modified,omitempty"`
	Origin   int64     `json:"origin,omitempty"`
}

// Reading - Sensor measurement
type Reading struct {
	Pushed   int64  `json:"pushed,omitempty"`
	Name     string `json:"name,omitempty"`
	Value    string `json:"value,omitempty"`
	Device   string `json:"device,omitempty"`
	Created  int64  `json:"created,omitempty"`
	Modified int64  `json:"modified,omitempty"`
	Origin   int64  `json:"origin,omitempty"`
}
