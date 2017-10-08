package export

// Compression algorithm types
const (
	CompNone = iota
	CompGzip
	CompZip
)

// Data format types
const (
	FormatJson = iota
	FormatXml
	FormatSerialized
	FormatIoTCoreJson
	FormatAzureJson
	FormatCsv
)

// Export destination types
const (
	DestMqtt = iota
	DestZmqt
	DestIoTCoreMqtt
	DestAzureMqtt
	DestRest
)

type ExportRegistration struct {
	Name        string
	Addr        Addressable
	Format      int
	Filter      ExportFilter
	Encryption  EncryptionDetails
	Compression int
	Enable      bool
	Destination int
}
