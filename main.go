package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coaraujo/go-mongo-rabbitmq/service"

	"github.com/coaraujo/go-mongo-rabbitmq/mongodb"

	"github.com/coaraujo/go-mongo-rabbitmq/rabbitmq"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("[MAIN] Starting project...")

	mongodb.Connect()
	rabbitmq.Initializer()

	router := mux.NewRouter()
	router.HandleFunc("/vote", voteService.SendVote).Methods("POST")
	router.HandleFunc("/vote/{id}", voteService.GetVote).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

	defer mongodb.CloseConnection()
	defer rabbitmq.CloseConnection()
}
