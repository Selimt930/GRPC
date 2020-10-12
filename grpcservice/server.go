package grpcservice

import (
	"MailService/grpc2/grpcstorage"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Service ...

//InitService ...
func (s *Service) InitService() *Service {
	s.strg.SetStorage(grpcstorage.NewStorage())
	return s
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Run() {
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
	r.HandleFunc("/inbox", s.GetUserMessage).Methods("GET")
	r.HandleFunc("/messages/{id}", s.GetMessage).Methods("GET")
	r.HandleFunc("/messages", s.WriteMessage).Methods("POST")
	//r.HandleFunc("/messages", temp.RequestVerification(s.WriteMessage)).Methods("POST")
	r.HandleFunc("/messages/{id}", s.UpdateMessage).Methods("PUT")
	r.HandleFunc("/messages/{id}", s.DeleteMessage).Methods("DELETE")
	r.HandleFunc("/delete", s.DeleteUserMessage).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))

}
