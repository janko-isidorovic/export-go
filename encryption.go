package export

// Encryption types
const (
	EncNone = iota
	EncAes
)

type EncryptionDetails struct {
	Algo       int
	Key        string
	InitVector string
}
