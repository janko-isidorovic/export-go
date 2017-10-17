package distro

import (
	"fmt"
	// "os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type mqttSender struct {
	mqttClient MQTT.Client
}

func NewMqttSender(broker string, user string, password string) Sender {

	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("edgex")
	opts.SetUsername(user)
	opts.SetPassword(password)
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
