package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Vote struct {
	option    int
	paredaoID string
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("vote").C("testing")
	err = c.Insert(
		&Vote{2, "1"},
		&Vote{1, "1"})
	if err != nil {
		log.Fatal(err)
	}

	result := Vote{}
	err = c.Find(bson.M{"name": "Rafael"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Option:", result.option, " ParedaoId:", result.paredaoID)
}
