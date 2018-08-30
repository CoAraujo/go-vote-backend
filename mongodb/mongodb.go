package mongodb

import (
	"fmt"
	"log"

	"github.com/coaraujo/go-mongo-rabbitmq/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	host           = "localhost"
	database       = "vote"
	collectionName = "vote"
)

type MongoDB struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func NewConnection() *MongoDB {
	fmt.Println("[MONGODB] Connecting...")
	mongodb := MongoDB{}

	s, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	mongodb.session = s

	c := s.DB(database).C(collectionName)
	mongodb.collection = c
	fmt.Println("[MONGODB] Connection started sucessfully.")

	return &mongodb
}

func (m *MongoDB) CloseConnection() {
	fmt.Println("[MONGODB] Closing connection...")

	m.session.Close()

	fmt.Println("[MONGODB] Connection closed.")
}

func (m *MongoDB) InsertVote(vote domain.Vote) {
	fmt.Println("[MONGODB] Inserting value: ", vote)

	err := m.collection.Insert(vote)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[MONGODB] Value inserted: ", vote)
}

func (m *MongoDB) GetVote(paredaoId string) {
	fmt.Println("[MONGODB] Getting vote for paredaoId:", paredaoId)

	result := domain.Vote{}
	err := m.collection.Find(bson.M{"name": paredaoId}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[MONGODB] Get vote:", result)
}
