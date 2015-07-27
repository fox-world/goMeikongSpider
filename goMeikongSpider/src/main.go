package main

import (

	"labix.org/v2/mgo"
    "service"
)

func main() {
	session, err := mgo.Dial("127.0.0.1")
	defer session.Close()

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	modelCollection := session.DB("meikong").C("model")
	//先删除所有的记录
	modelCollection.RemoveAll(nil)

	service.MeikongSpider("http://www.moko.cc/channels/post/23/1.html",0,modelCollection)
}

