//
// Copyright (c) Mainflux
//
// Mainflux server is licensed under an Apache license, version 2.0.
// All rights not explicitly granted in the Apache license, version 2.0 are reserved.
// See the included LICENSE file for more details.
//

package export

// Protocols
const (
	ProtoHTTP  = "HTTP"
	ProtoTCP   = "TCP"
	ProtoMAC   = "MAC"
	ProtoZMQ   = "ZMQ"
	ProtoOther = "OTHER"
)

// Methods
const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodPatch  = "PATCH"
	MethodDelete = "DELETE"
)

// Addressable - address for reaching the service
type Addressable struct {
	Name      string
	Method    string
	Protocol  string
	Address   string
	Port      int
	Path      string
	Publisher string
	User      string
	Password  string
	Topic     string
}
