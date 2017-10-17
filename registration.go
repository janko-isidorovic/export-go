package export

// Compression algorithm types
const (
	CompNone = iota
	CompGzip
	CompZip
)

// Data format types
const (
	FormatJSON = iota
	FormatXML
	FormatSerialized
	FormatIoTCoreJSON
	FormatAzureJSON
	FormatCSV
)

// Export destination types
const (
	DestMQTT = iota
	DestZMQ
	DestIotCoreMQTT
	DestAzureMQTT
	DestRest
)

// Registration - Defines the registration details
// on the part of north side export clients
type Registration struct {
	ID          string            `json:"id,omitempty"`
	Created     int64             `json:"created,omitempty"`
	Modified    int64             `json:"modified,omitempty"`
	Origin      int64             `json:"origin,omitempty"`
	Name        string            `json:"name,omitempty"`
	Addr        Addressable       `json:"addressable,omitempty"`
	Format      int               `json:"format,omitempty"`
	Filter      Filter            `json:"filter,omitempty"`
	Encryption  EncryptionDetails `json:"encryption,omitempty"`
	Compression int               `json:"compression,omitempty"`
	Enable      bool              `json:"enable,omitempty"`
	Destination int               `json:"destination,omitempty"`
}
