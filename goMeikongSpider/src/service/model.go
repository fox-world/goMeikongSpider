package service

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"models"
)

func QueryPage() []models.Model {
	session, err := mgo.Dial("127.0.0.1")

	defer session.Close()

	m := session.DB("meikong").C("model")

	var results []models.Model

	err = m.Find(bson.M{"number": bson.M{"$gte": 1, "$lte": 50}}).Sort("number").All(&results)

	if err != nil {
		panic(err)
	}

	return results
}
