package voteService

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/coaraujo/go-mongo-rabbitmq/domain"
	"github.com/coaraujo/go-mongo-rabbitmq/rabbitmq"
)

func SendVote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[VOTESERVICE] Send vote was invoked...")

	var vote domain.Vote
	_ = json.NewDecoder(r.Body).Decode(&vote)
	rabbitmq.SendVote(vote)

	// json.NewEncoder(w).Encode(vote)
	// fmt.Println("Meu voto foi: ", vote)
}

func GetVote(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// for _, item := range people {
	// 	if item.ID == params["id"] {
	// 	}
	// }
}
