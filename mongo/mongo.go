//
// Copyright (c) Mainflux
//
// Mainflux server is licensed under an Apache license, version 2.0.
// All rights not explicitly granted in the Apache license, version 2.0 are reserved.
// See the included LICENSE file for more details.
//

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
