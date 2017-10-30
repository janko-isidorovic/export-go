//
// Copyright (c) 2017
// Cavium
// Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"github.com/drasko/edgex-export"
	"github.com/drasko/edgex-export/mongo"

	"go.uber.org/zap"
)

var repo *mongo.Repository

// InitMongoRepository - Init Mongo DB
func InitMongoRepository(r *mongo.Repository) {
	repo = r
	return
}

func getRegistrations(repo *mongo.Repository) []export.Registration {

	s := repo.Session.Copy()
	defer s.Close()

	c := s.DB(mongo.DBName).C(mongo.CollectionName)

	results := []export.Registration{}
	err := c.Find(nil).All(&results)
	if err != nil {
		logger.Error("Failed to query", zap.Error(err))
		return nil
	}

	return results
}
