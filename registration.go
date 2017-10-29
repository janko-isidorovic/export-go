//
// Copyright 2017 Mainflux.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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
