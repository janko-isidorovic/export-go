package distro

import (
	"fmt"
	"github.com/drasko/edgex-export"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"strconv"
	"strings"
)

type mqttSender struct {
	mqttClient MQTT.Client
}

// Change parameters to Addressable?
func NewMqttSender(addr export.Addressable) Sender {

	opts := MQTT.NewClientOptions()
	broker := strings.ToLower(addr.Protocol) + "://" + addr.Address + ":" + strconv.Itoa(addr.Port)
	opts.AddBroker(broker)
	opts.SetClientID("edgex")
	opts.SetUsername(addr.User)
	opts.SetPassword(addr.Password)
	// opts.SetCleanSession(cleansess)

	var sender mqttSender

	sender.mqttClient = MQTT.NewClient(opts)
	if token := sender.mqttClient.Connect(); token.Wait() && token.Error() != nil {
		// FIXME
		panic(token.Error())
	}
	fmt.Println("Sample Publisher Started")

	return sender
}

func (sender mqttSender) Send(data string) {
	token := sender.mqttClient.Publish("FCR", 0, false, data)
	// FCR could be removed? set of tokens?
	token.Wait()
	fmt.Println("Sent data: " + data)
}
