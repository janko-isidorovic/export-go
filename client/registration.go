package client

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/drasko/edgex-export"
	"github.com/drasko/edgex-export/mongo"
	"github.com/go-zoo/bone"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
)

func getRegByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id := bone.GetValue(r, "id")

	s := repo.Session.Copy()
	defer s.Close()

	c := s.DB(mongo.DbName).C(mongo.CollectionName)
	result := export.Registration{}

	if err := c.Find(bson.M{"id": id}).One(&result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	res, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(res))
}

func getRegList(w http.ResponseWriter, r *http.Request) {
}

func getAllReg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	s := repo.Session.Copy()
	defer s.Close()

	c := s.DB(mongo.DbName).C(mongo.CollectionName)

	results := []export.Registration{}
	err := c.Find(nil).All(&results)
	if err != nil {
		logger.Error("Failed to query", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func getRegByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	name := bone.GetValue(r, "name")

	s := repo.Session.Copy()
	defer s.Close()

	c := s.DB(mongo.DbName).C(mongo.CollectionName)
	result := export.Registration{}

	if err := c.Find(bson.M{"name": name}).One(&result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	res, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(res))
}

func addReg(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reg := export.Registration{}
	if err := json.Unmarshal(data, &reg); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	s := repo.Session.Copy()
	defer s.Close()
	c := s.DB(mongo.DbName).C(mongo.CollectionName)

	count, _ := c.Find(bson.M{"name": reg.Name}).Count()
	if count != 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.Insert(reg); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func updateReg(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var body map[string]interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
	}

	s := repo.Session.Copy()
	defer s.Close()
	c := s.DB(mongo.DbName).C(mongo.CollectionName)

	name := body["name"]
	query := bson.M{"name": name}
	update := bson.M{"$set": body}

	if err := c.Update(query, update); err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func delRegByID(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")

	s := repo.Session.Copy()
	defer s.Close()
	c := s.DB(mongo.DbName).C(mongo.CollectionName)

	if err := c.Remove(bson.M{"id": id}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func delRegByName(w http.ResponseWriter, r *http.Request) {
	name := bone.GetValue(r, "name")

	s := repo.Session.Copy()
	defer s.Close()
	c := s.DB(mongo.DbName).C(mongo.CollectionName)

	if err := c.Remove(bson.M{"name": name}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
