package distro

import (
	"fmt"
	"github.com/drasko/edgex-export"
	"log"
	"net/http"
	"strconv"
)

type httpSender struct {
	url    string
	method string
}

// Change parameter to Addressable?
func NewHttpSender(addr export.Addressable) Sender {
	var sender httpSender
	// Should be added protocol from Addressable instead of include it the address param. We will maintain this behaviour for compatibility with Java
	//sender.url = strings.ToLower(addr.Protocol) + "://" + addr.Address + ":" + strconv.Itoa(addr.Port) + addr.Path
	sender.url = addr.Address + ":" + strconv.Itoa(addr.Port) + addr.Path
	sender.method = addr.Method
	return sender
}

func (sender httpSender) Send(data string) {
	switch sender.method {

	case export.MethodGet:
		response, err := http.Get(sender.url)
		if err != nil {
			//FIXME
			log.Fatal(err)
		} else {
			defer response.Body.Close()
			fmt.Println("Response: ", response.Status)
			// //_, err := io.Copy(os.Stdout, response.Body)
			// if err != nil {
			// 	log.Fatal(err)
			// }
		}

	case export.MethodPost:
		var buf string
		response, err := http.Post(sender.url, "application/json", nil)
		if err != nil {
			//FIXME
			log.Fatal(err)
		} else {
			defer response.Body.Close()
			fmt.Println("Response: ", response.Status)
			fmt.Println("Buf: ", buf)
			// //_, err := io.Copy(os.Stdout, response.Body)
			// if err != nil {
			// 	log.Fatal(err)
			// }
		}

	case export.MethodPut:
		fmt.Println("TBD method: ", sender.method)
	case export.MethodPatch:
		fmt.Println("TBD method: ", sender.method)
	case export.MethodDelete:
		fmt.Println("TBD method: ", sender.method)
	default:
		fmt.Println("Unsupported method: ", sender.method)
	}

	fmt.Println("Sent data: " + data)
}
