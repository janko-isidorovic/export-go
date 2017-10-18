package distro

import (
	"fmt"
	"github.com/drasko/edgex-export"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type httpSender struct {
	url string
}

// Change parameter to Addressable?
func NewHttpSender(addr export.Addressable) Sender {
	var sender httpSender
	sender.url = strings.ToLower(addr.Protocol) + "://" + addr.Address + ":" + strconv.Itoa(addr.Port) + addr.Path
	return sender
}

func (sender httpSender) Send(data string) {
	response, err := http.Get(sender.url)
	if err != nil {
		//FIXME
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		// fmt.Println("Response: ", response.Status)
		// //_, err := io.Copy(os.Stdout, response.Body)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}
	fmt.Println("Sent data: " + data)
}
