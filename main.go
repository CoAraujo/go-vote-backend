package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coaraujo/go-mongo-rabbitmq/rabbitmq"
	"github.com/coaraujo/go-mongo-rabbitmq/service"

	"github.com/coaraujo/go-mongo-rabbitmq/mongodb"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("[MAIN] Starting project...")

	mongoCon := *mongodb.NewConnection()
	rabbitCon := *rabbitmq.NewConnection()

	voteServ := voteService.Init(&rabbitCon, &mongoCon)

	router := mux.NewRouter()
	router.HandleFunc("/vote", voteServ.SendVote).Methods("POST")
	router.HandleFunc("/vote/{id}", voteServ.GetVote).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

	defer mongoCon.CloseConnection()
	defer rabbitCon.CloseConnection()
}
