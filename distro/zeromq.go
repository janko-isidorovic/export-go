package distro

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func initZmq() {
	q, _ := zmq.NewSocket(zmq.SUB)
	defer q.Close()

	logger.Info("Connecting to zmq...")
	q.Connect("tcp://localhost:5563")
	logger.Info("Connected to zmq")

	for {
		msg, err := q.RecvMessage(0)
		logger.Info("Received zmq msg")
		if err == nil {
			id, _ := q.GetIdentity()
			fmt.Println("ERROR:", msg[0], id)
		} else {
			fmt.Println("MSG  :", msg)
		}
	}
}
