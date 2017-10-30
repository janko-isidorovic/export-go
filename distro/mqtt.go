//
// Copyright (c) 2017 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"strconv"

	"github.com/drasko/edgex-export"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

type mqttSender struct {
	client MQTT.Client
	topic  string
}

const clientID = "edgex"
const topic = "EdgeX"

func NewMqttSender(addr export.Addressable) Sender {
	opts := MQTT.NewClientOptions()
	// CHN: Should be added protocol from Addressable instead of include it the address param.
	// CHN: We will maintain this behaviour for compatibility with Java
	broker := addr.Address + ":" + strconv.Itoa(addr.Port)
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername(addr.User)
	opts.SetPassword(addr.Password)

	sender := mqttSender{
		client: MQTT.NewClient(opts),
		topic:  addr.Topic,
	}

	if token := sender.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	logger.Info("Sample Publisher Started")

	return sender
}

func (sender mqttSender) Send(data []byte) {
	token := sender.client.Publish(sender.topic, 0, false, data)
	// FIXME: could be removed? set of tokens?
	token.Wait()
	logger.Debug("Sent data: ", zap.ByteString("data", data))
}
