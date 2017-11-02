//
// Copyright (c) 2017 Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package export

// Encryption types
const (
	EncNone = "NONE"
	EncAes  = "AES"
)

// EncryptionDetails - Provides details for encryption
// of export data per client request
type EncryptionDetails struct {
	Algo       string
	Key        string
	InitVector string
}
