package distro

import (
	"bytes"
	"testing"
	"time"

	"github.com/drasko/edgex-export"
	"github.com/surgemq/message"
	"github.com/surgemq/surgemq/service"
	"go.uber.org/zap"
)

const addressAndPort = "tcp://127.0.0.1:12883"
const stringCompare = "Hello, World!"
const topic = "EdGeX"

func runClient() {
	// Instantiates a new Client
	c := &service.Client{}
	// Creates a new MQTT CONNECT message and sets the proper parameters
	msg := message.NewConnectMessage()
	msg.SetVersion(4)
	msg.SetCleanSession(true)
	msg.SetClientId([]byte("edgex"))
	msg.SetKeepAlive(10)
	msg.SetUsername([]byte("user"))
	msg.SetPassword([]byte("password"))

	err := c.Connect(addressAndPort, msg)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}
	submsg := message.NewSubscribeMessage()
	submsg.AddTopic([]byte(topic), 0)
	c.Subscribe(submsg, nil, publishFunc)
}

func TestMqttNew(t *testing.T) {
	log, _ := zap.NewProduction()
	defer log.Sync()

	InitLogger(log)

	svr := &service.Server{
		KeepAlive:        300,           // seconds
		ConnectTimeout:   2,             // seconds
		SessionsProvider: "mem",         // keeps sessions in memory
		Authenticator:    "mockSuccess", // always succeed
		TopicsProvider:   "mem",         // keeps topic subscriptions in memory
	}

	go svr.ListenAndServe(addressAndPort)
	defer svr.Close()

	sender := NewMqttSender(export.Addressable{
		Address:  "tcp://127.0.0.1",
		Port:     12883,
		User:     "user",
		Password: "password",
		Topic:    topic,
	})

	time.Sleep(2 * time.Second)
	go runClient()
	time.Sleep(2 * time.Second)
	sender.Send([]byte("Hello, World!"))
	time.Sleep(1 * time.Second)
}

func publishFunc(msg *message.PublishMessage) error {

	log.Info("Message received: ", zap.ByteString("Payload:", msg.Payload()))

	if !bytes.Equal([]byte(stringCompare), msg.Payload()) {
		log.Info("Test failed")
		return nil
	}

	log.Info("Test OK")

	return nil
}
