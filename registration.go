//
// Copyright (c) 2017 Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package export

// Compression algorithm types
const (
	CompNone = "NONE"
	CompGzip = "GZIP"
	CompZip  = "ZIP"
)

// Data format types
const (
	FormatJSON        = "JSON"
	FormatXML         = "XML"
	FormatSerialized  = "SERIALIZED"
	FormatIoTCoreJSON = "IOTCORE_JSON"
	FormatAzureJSON   = "AZURE_JSON"
	FormatCSV         = "CSV"
)

// Export destination types
const (
	DestMQTT        = "MQTT_TOPIC"
	DestZMQ         = "ZMQ_TOPIC"
	DestIotCoreMQTT = "IOTCORE_TOPIC"
	DestAzureMQTT   = "AZURE_TOPIC"
	DestRest        = "REST_ENDPOINT"
)

// Registration - Defines the registration details
// on the part of north side export clients
type Registration struct {
	ID          string            `json:"_id,omitempty"`
	Created     int64             `json:"created,omitempty"`
	Modified    int64             `json:"modified,omitempty"`
	Origin      int64             `json:"origin,omitempty"`
	Name        string            `json:"name,omitempty"`
	Addressable Addressable       `json:"addressable,omitempty"`
	Format      string            `json:"format,omitempty"`
	Filter      Filter            `json:"filter,omitempty"`
	Encryption  EncryptionDetails `json:"encryption,omitempty"`
	Compression string            `json:"compression,omitempty"`
	Enable      bool              `json:"enable,omitempty"`
	Destination string            `json:"destination,omitempty"`
}
