package voteService

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/coaraujo/go-mongo-rabbitmq/mongodb"

	"github.com/coaraujo/go-mongo-rabbitmq/rabbitmq"

	"github.com/coaraujo/go-mongo-rabbitmq/domain"
)

type VoteService struct {
	rabbitmq *rabbitmq.RabbitMQ
	mongodb  *mongodb.MongoDB
}

func Init(r *rabbitmq.RabbitMQ, m *mongodb.MongoDB) *VoteService {
	voteService := VoteService{rabbitmq: r, mongodb: m}
	return &voteService
}

func (v *VoteService) SendVote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[VOTESERVICE] Send vote was invoked...")

	var vote domain.Vote
	_ = json.NewDecoder(r.Body).Decode(&vote)
	v.rabbitmq.SendVote(vote)

	// json.NewEncoder(w).Encode(vote)
	// fmt.Println("Meu voto foi: ", vote)
}

func (v *VoteService) GetVote(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// for _, item := range people {
	// 	if item.ID == params["id"] {
	// 	}
	// }
}
