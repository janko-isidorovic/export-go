package export

type ExportRegistration struct {
	Name        string
	Addr        Addressable
	Format      ExportFormat
	Filter      ExportFilter
	Encryption  EncryptionDetails
	Compression ExportCompression
	Enable      bool
	Destination ExportDestination
}
