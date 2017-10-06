package api

import (
	"gopkg.in/mgo.v2"
)

var ms *mgo.Session

func SetMongoSession(s *mgo.Session) {
	ms = s
}
