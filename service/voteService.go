package voteService

import (
	"encoding/json"
	"fmt"
	"net/http"

	client "github.com/coaraujo/go-vote-backend/client/recaptcha"
	mongo "github.com/coaraujo/go-vote-backend/config/mongo"
	domain "github.com/coaraujo/go-vote-backend/domain"
	stream "github.com/coaraujo/go-vote-backend/stream"
)

type VoteService struct {
	MongoDB      *mongo.MongoDB
	RabbitStream *stream.RabbitStream
}

func NewVoteService(m *mongo.MongoDB, s *stream.RabbitStream) *VoteService {
	voteService := VoteService{MongoDB: m, RabbitStream: s}
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

	v.RabbitStream.SendVote(vote)
}

func (v *VoteService) GetVote(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// for _, item := range people {
	// 	if item.ID == params["id"] {
	// 	}
	// }
}
