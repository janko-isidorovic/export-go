//
// Copyright (c) Mainflux
//
// Mainflux server is licensed under an Apache license, version 2.0.
// All rights not explicitly granted in the Apache license, version 2.0 are reserved.
// See the included LICENSE file for more details.
//

package client

import "github.com/drasko/edgex-export/mongo"

var repo *mongo.MongoRepository

func InitMongoRepository(r *mongo.MongoRepository) {
	repo = r
	return
}
