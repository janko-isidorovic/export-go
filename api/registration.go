package api

import (
	"encoding/json"
	"net/http"

	"github.com/drasko/edgex-exportclient/mongo"
	"go.uber.org/zap"
)

type Registration struct{}

func getAllReg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	s := repo.Session.Copy()
	defer s.Close()

	c := s.DB(mongo.DbName).C(mongo.CollectionName)

	results := []Registration{}
	err := c.Find(nil).All(&results)
	if err != nil {
		logger.Error("Failed to query", zap.Error(err))
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}
