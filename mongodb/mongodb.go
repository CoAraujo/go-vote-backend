package mongodb

import (
	"fmt"
	"log"

	"github.com/coaraujo/go-mongo-rabbitmq/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session
var collection *mgo.Collection

func Connect() {
	fmt.Println("[MONGODB] Connecting...")

	s, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	session = s

	c := s.DB("vote").C("vote")
	collection = c
	fmt.Println("[MONGODB] Connection started sucessfully.")
}

func CloseConnection() {
	fmt.Println("[MONGODB] Closing connection...")

	session.Close()

	fmt.Println("[MONGODB] Connection closed.")
}

func InsertVote(vote domain.Vote) {
	fmt.Println("[MONGODB] Inserting value: ", vote)

	err := collection.Insert(vote)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[MONGODB] Value inserted: ", vote)
}

func GetVote(paredaoId string) {
	fmt.Println("[MONGODB] Getting vote for paredaoId:", paredaoId)

	result := domain.Vote{}
	err := collection.Find(bson.M{"name": paredaoId}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[MONGODB] Get vote:", result)
}
