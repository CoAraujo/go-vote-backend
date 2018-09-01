package repository

import (
	"fmt"
	"log"

	mongo "github.com/coaraujo/go-vote-backend/config/mongo"
	domain "github.com/coaraujo/go-vote-backend/domain"
	"gopkg.in/mgo.v2/bson"
)

const (
	voteDatabase       = "vote"
	voteCollectionName = "vote"
)

type VoteRepository struct {
	MongoDB *mongo.MongoDB
}

func newVoteRepository() *VoteRepository {
	voteRepository := VoteRepository{MongoDB: mongo.NewConnection()}
	return &voteRepository
}

func (v *VoteRepository) GetVoteById(id string) {
	fmt.Println("[MONGODB] Getting vote by id:", id)

	result := domain.Vote{}
	err := v.MongoDB.GetCollection(voteDatabase, voteCollectionName).Find(bson.M{"id": id}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[MONGODB] Get vote:", result)
}
