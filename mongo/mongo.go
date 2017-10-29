package mongo

import (
	"gopkg.in/mgo.v2"
)

const (
	DbName         string = "coredata"
	CollectionName string = "exportConfiguration"
)

type MongoRepository struct {
	Session *mgo.Session
}

func NewMongoRepository(ms *mgo.Session) *MongoRepository {
	return &MongoRepository{Session: ms}
}
