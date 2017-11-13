package distro

import (
	"github.com/drasko/edgex-export"
	zmq "github.com/pebbe/zmq4"
	"go.uber.org/zap"
)

func ZeroMQReceiver(eventCh chan *export.Event) {
	go initZmq(eventCh)
}

func initZmq(eventCh chan *export.Event) {
	q, _ := zmq.NewSocket(zmq.SUB)
	defer q.Close()

	logger.Info("Connecting to zmq...")
	q.Connect("tcp://localhost:32768")
	logger.Info("Connected to zmq")
	q.SetSubscribe("")

	for {
		msg, err := q.RecvMessage(0)
		if err != nil {
			id, _ := q.GetIdentity()
			logger.Error("Error getting mesage", zap.String("id", id))
		} else {
			for _, str := range msg {
				// Why the offset of 7?? zmq v3 vs v4 ?
				event := parseEvent(str[7:])
				logger.Info("Event received", zap.Any("event", event))
				eventCh <- event
			}
		}
	}
}
