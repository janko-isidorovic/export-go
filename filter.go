//
// Copyright (c) Mainflux
//
// Mainflux server is licensed under an Apache license, version 2.0.
// All rights not explicitly granted in the Apache license, version 2.0 are reserved.
// See the included LICENSE file for more details.
//

package export

// Filter - Specifies the client filters on reading data
type Filter struct {
	DeviceIDs          []string
	ValueDescriptorIDs []string
}
