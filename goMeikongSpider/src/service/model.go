package service

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"models"
)

func QueryPage(dbUri string, pageNo int, pageSize int) ([]models.Model, models.Page) {
	session, err := mgo.Dial(dbUri)

	defer session.Close()

	m := session.DB("meikong").C("model")

	var results []models.Model

	start := (pageNo - 1) * pageSize
	end := pageNo * pageSize

	err = m.Find(bson.M{"number": bson.M{"$gt": start, "$lte": end}}).Sort("number").All(&results)

	if err != nil {
		panic(err)
	}

	totalRecord, err := m.Find(bson.M{}).Count()

	if err != nil {
		panic(err)
	}

	totalPage := 0
	if totalRecord%pageSize == 0 {
		totalPage = totalRecord / pageSize
	} else {
		totalPage = totalRecord/pageSize + 1
	}

	var next, previous = pageNo + 1, pageNo - 1

	if pageNo == 1 {
		previous = 1
	}

	if pageNo == totalPage {
		next = totalPage
	}

	page := models.Page{PageNo: pageNo, PageSize: pageSize, TotalRecord: totalRecord, TotalPage: totalPage, Next: next, Previous: previous}

	return results, page
}
