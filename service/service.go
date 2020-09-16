package service

import (
	"MailService/auth"
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

//Message struct defining key features for massage obj
type Message struct {
	ID      string  `json:"id"`
	In      string  `json:"in"`
	Content string  `json:"content"`
	Author  *Author `json:"author"`
	To      string  `json:"name"`
}

// Author struct
type Author struct {
	ID        string `json:"idAU"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var Users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Init Messages var as a slice Message struct
var Messages []Message

// Get all massages ever sent
func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Messages)
}

// Get single massage by its id
func GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params

	// Loop through massages and find one with the id from params
	for _, mas := range Messages {
		if mas.ID == params["id"] {
			json.NewEncoder(w).Encode(mas)
			return
		}
	}
	json.NewEncoder(w).Encode(&Message{})
}

//Get massage by users id
func GetUserMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mas Message
	_ = json.NewDecoder(r.Body).Decode(&mas)

	for _, i := range Messages {
		if mas.Author.Firstname == i.To {
			json.NewEncoder(w).Encode(i)
			return
		}
	}
	//json.NewEncoder(w).Encode(&Message{})
}

// Add new massage
func WriteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := r.Context().Value(auth.KeyContext)
	var mas Message

	_ = json.NewDecoder(r.Body).Decode(&mas)
	mas.ID = strconv.Itoa(rand.Intn(100000000))
	mas.Author.ID = user.(string)
	Messages = append(Messages, mas)
	json.NewEncoder(w).Encode(mas)
}

// Update sent massage
func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, mas := range Messages {
		if mas.ID == params["id"] {
			Messages = append(Messages[:index], Messages[index+1:]...)
			var mas Message
			_ = json.NewDecoder(r.Body).Decode(&mas)
			mas.ID = params["id"]
			Messages = append(Messages, mas)
			json.NewEncoder(w).Encode(mas)
			return
		}
	}
}

// Delete massage
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, mas := range Messages {
		if mas.ID == params["id"] {
			Messages = append(Messages[:index], Messages[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Messages)
}
