package main

import (
	"MailService/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//init middleware
	//temp := auth.NewReqVerification(auth.GetUserFromToken)
	// Init router

	//inner.LocalStotage(storage.NewMemStore())
	r := mux.NewRouter()

	// Hardcoded data for 'Messages'
	//service.Messages = append(service.Messages, service.Message{ID: "1", In: "438227", Content: "One", Author: &service.Author{Firstname: "John", Lastname: "Doe"}})
	//service.Messages = append(service.Messages, service.Message{ID: "2", In: "454555", Content: "Two", Author: &service.Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route handles & endpoints
	//r.HandleFunc("/message", service.GetUserMessage).Methods("GET")
	//r.HandleFunc("/messages", service.GetAllMessages).Methods("GET")
	r.HandleFunc("/messages/{id}", service.GetMessage).Methods("GET")
	r.HandleFunc("/messages", service.WriteMessage).Methods("POST")
	//r.HandleFunc("/messages", temp.RequestVerification(service.WriteMessage)).Methods("POST")
	r.HandleFunc("/messages/{id}", service.UpdateMessage).Methods("PUT")
	r.HandleFunc("/messages/{id}", service.DeleteMessage).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
