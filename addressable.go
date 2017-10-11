package export

// Protocols
const (
   ProtoHTTP = iota
   ProtoTCP
   ProtoMAC
   ProtoZMQ
   ProtoOther
)

// Methods
const (
   MethodGet = iota
   MethodPost
   MethodPut
   MethodPatch
   MethodDelete
)

// Addressable - address for reaching the service
type Addressable struct {
   Name       string
   Method     int
   Protocol   int
   Address    string
   Port       int
   Path       int
   Publisher  string
   User       string
   Password   string
   Topic      string
}
