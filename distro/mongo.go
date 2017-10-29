package distro

import (
	"github.com/drasko/edgex-export"
	"github.com/drasko/edgex-export/mongo"

	"go.uber.org/zap"
)

var repo *mongo.MongoRepository

func InitMongoRepository(r *mongo.MongoRepository) {
	repo = r
	return
}

func getRegistrations(repo *mongo.MongoRepository) []export.Registration {

	s := repo.Session.Copy()
	defer s.Close()

	c := s.DB(mongo.DbName).C(mongo.CollectionName)

	results := []export.Registration{}
	err := c.Find(nil).All(&results)
	if err != nil {
		logger.Error("Failed to query", zap.Error(err))
		return nil
	}

	return results
}
