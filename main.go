package main

import (
	"fmt"
	"log"
	"net/http"

	mongo "github.com/coaraujo/go-vote-backend/config/mongo"
	rabbit "github.com/coaraujo/go-vote-backend/config/rabbit"
	service "github.com/coaraujo/go-vote-backend/service"
	stream "github.com/coaraujo/go-vote-backend/stream"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("[MAIN] Starting project...")

	mongoCon := *mongo.NewConnection()
	rabbitCon := *rabbit.NewConnection()
	rabbitStream := *stream.NewRabbitStream(&rabbitCon)
	voteServ := service.NewVoteService(&mongoCon, &rabbitStream)

	router := mux.NewRouter()
	router.HandleFunc("/vote", voteServ.SendVote).Methods("POST")
	router.HandleFunc("/vote/{id}", voteServ.GetVote).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

	defer mongoCon.CloseConnection()
	defer rabbitCon.CloseConnection()
}
