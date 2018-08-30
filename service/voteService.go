package voteService

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/coaraujo/go-mongo-rabbitmq/client/recaptcha"

	"github.com/coaraujo/go-mongo-rabbitmq/mongodb"

	"github.com/coaraujo/go-mongo-rabbitmq/rabbitmq"

	"github.com/coaraujo/go-mongo-rabbitmq/domain"
)

type VoteService struct {
	rabbitmq *rabbitmq.RabbitMQ
	mongodb  *mongodb.MongoDB
}

func NewVoteService(r *rabbitmq.RabbitMQ, m *mongodb.MongoDB) *VoteService {
	voteService := VoteService{rabbitmq: r, mongodb: m}
	return &voteService
}

func (v *VoteService) SendVote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[VOTESERVICE] Send vote was invoked...")

	var vote domain.Vote
	_ = json.NewDecoder(r.Body).Decode(&vote)

	if client.Verify(&vote, r.Header.Get("x-auth")) == false {
		fmt.Println("[VOTESERVICE] Error 401 - Unauthorized.")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - Check the recaptcha!"))
		return
	}

	v.rabbitmq.SendVote(vote)
}

func (v *VoteService) GetVote(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// for _, item := range people {
	// 	if item.ID == params["id"] {
	// 	}
	// }
}
