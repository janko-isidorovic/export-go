//
// Copyright (c) Mainflux
//
// Mainflux server is licensed under an Apache license, version 2.0.
// All rights not explicitly granted in the Apache license, version 2.0 are reserved.
// See the included LICENSE file for more details.
//

package export

// Encryption types
const (
	EncNone = iota
	EncAes
)

// EncryptionDetails - Provides details for encryption
// of export data per client request
type EncryptionDetails struct {
	Algo       int
	Key        string
	InitVector string
}
