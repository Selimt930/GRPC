package main

import (
	"MailService/auth"
	"MailService/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//init middleware
	temp := auth.NewReqVerification(auth.GetUserFromToken)
	// Init router
	r := mux.NewRouter()

	// Hardcoded data for 'Massages'
	service.Massages = append(service.Massages, service.Massage{ID: "1", In: "438227", Content: "One", Author: &service.Author{Firstname: "John", Lastname: "Doe"}})
	service.Massages = append(service.Massages, service.Massage{ID: "2", In: "454555", Content: "Two", Author: &service.Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route handles & endpoints
	r.HandleFunc("/massage", service.GetUserMassage).Methods("GET")
	r.HandleFunc("/massages", service.GetAllMassages).Methods("GET")
	r.HandleFunc("/massages/{id}", service.GetMassage).Methods("GET")
	r.HandleFunc("/massages", temp.RequestVerification(service.WriteMassage)).Methods("POST")
	r.HandleFunc("/massages/{id}", service.UpdateMassage).Methods("PUT")
	r.HandleFunc("/massages/{id}", service.DeleteMassage).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
