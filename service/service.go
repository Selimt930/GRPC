package service

import (
	"MailService/auth"
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

//Massage struct defining key features for massage obj
type Massage struct {
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

// Init Massages var as a slice Massage struct
var Massages []Massage

// Get all massages ever sent
func GetAllMassages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Massages)
}

// Get single massage by its id
func GetMassage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params

	// Loop through massages and find one with the id from params
	for _, mas := range Massages {
		if mas.ID == params["id"] {
			json.NewEncoder(w).Encode(mas)
			return
		}
	}
	json.NewEncoder(w).Encode(&Massage{})
}

//Get massage by users id
func GetUserMassage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mas Massage
	_ = json.NewDecoder(r.Body).Decode(&mas)

	for _, i := range Massages {
		if mas.Author.Firstname == i.To {
			json.NewEncoder(w).Encode(i)
			return
		}
	}
	//json.NewEncoder(w).Encode(&Massage{})
}

// Add new massage
func WriteMassage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := r.Context().Value(auth.KeyContext)
	var mas Massage

	_ = json.NewDecoder(r.Body).Decode(&mas)
	mas.ID = strconv.Itoa(rand.Intn(100000000))
	mas.Author.ID = user.(string)
	Massages = append(Massages, mas)
	json.NewEncoder(w).Encode(mas)
}

// Update sent massage
func UpdateMassage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, mas := range Massages {
		if mas.ID == params["id"] {
			Massages = append(Massages[:index], Massages[index+1:]...)
			var mas Massage
			_ = json.NewDecoder(r.Body).Decode(&mas)
			mas.ID = params["id"]
			Massages = append(Massages, mas)
			json.NewEncoder(w).Encode(mas)
			return
		}
	}
}

// Delete massage
func DeleteMassage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, mas := range Massages {
		if mas.ID == params["id"] {
			Massages = append(Massages[:index], Massages[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Massages)
}
