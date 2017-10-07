package export

type EncryptionDetails struct {
	Algo       ExportEncryption
	Key        string
	InitVector string
}
