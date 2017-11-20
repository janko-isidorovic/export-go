//
// Copyright (c) 2017 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

// +build zeromq

package distro

import (
	"encoding/json"

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
	q.Connect("tcp://localhost:5563")
	logger.Info("Connected to zmq")
	q.SetSubscribe("")

	for {
		msg, err := q.RecvMessage(0)
		if err != nil {
			id, _ := q.GetIdentity()
			logger.Error("Error getting mesage", zap.String("id", id))
		} else {
			for _, str := range msg {
				event := parseEvent(str)
				logger.Info("Event received", zap.Any("event", event))
				eventCh <- event
			}
		}
	}
}

func parseEvent(str string) *export.Event {
	event := export.Event{}

	if err := json.Unmarshal([]byte(str), &event); err == nil {
		return &event
	}

	// Why the offset of 7?? zmq v3 vs v4 ?
	if err := json.Unmarshal([]byte(str[7:]), &event); err != nil {
		logger.Error("Failed to parse event", zap.Error(err))
		return nil
	}

	return &event
}
