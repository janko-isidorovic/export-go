package api

import "github.com/drasko/edgex-exportclient/mongo"

var repo *mongo.MongoRepository

func InitMongoRepository(r *mongo.MongoRepository) {
	repo = r
	return
}
