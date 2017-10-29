//
// Copyright (c) 2017 Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package client

import "github.com/drasko/edgex-export/mongo"

var repo *mongo.MongoRepository

func InitMongoRepository(r *mongo.MongoRepository) {
	repo = r
	return
}
