package main

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"poll/ent"
	l "poll/handler"
)

func main() {
	client, errOpenDB := ent.Open("sqlite3", "./entc.db?_fk=1")
	if errOpenDB != nil {
		panic(errOpenDB)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	h := &l.Handler{Client: client}

	// Create new Router instance
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api").Subrouter()

	// Create and Fetch Polls
	sub.Methods("GET").Path("/polls").HandlerFunc(h.GetAllPolls)
	sub.Methods("POST").Path("/polls").HandlerFunc(h.CreateNewPoll)
	sub.Methods("GET").Path("/polls/{id}").HandlerFunc(h.GetAllPollsByUser)

	// Cast Vote
	sub.Methods("POST").Path("/polls/vote").HandlerFunc(h.CastVote)

	//Fetch Votes
	sub.Methods("GET").Path("/votes").HandlerFunc(h.GetVotes)
	sub.Methods("GET").Path("/votes/{id}").HandlerFunc(h.GetVotesByUser)

	// User Signup and Login
	sub.Methods("POST").Path("/auth/signup").HandlerFunc(h.Register)
	sub.Methods("POST").Path("/auth/login").HandlerFunc(h.Login)
	sub.Methods("GET").Path("/user").HandlerFunc(h.GetUser)

	// Fetch list of answers
	sub.Methods("GET").Path("/polls/answers").HandlerFunc(h.GetAnswers)

	log.Print("Server Running on 8080")

	// To enable CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	err := http.ListenAndServe(":8080", handlers.CORS(headersOk, originsOk, methodsOk)(router))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
