package grpcservice

import (
	"MailService/grpc2"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Service ...
type Service struct {
	strg grpc2.Service
}

//Message struct defining key features for massage obj
// type Message struct {
// 	ID      string  `json:"id"`
// 	In      string  `json:"in"`
// 	Content string  `json:"content"`
// 	Author  *Author `json:"author"`
// 	To      string  `json:"name"`
// }

// Author struct
// type Author struct {
// 	ID        string `json:"idAU"`
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// }

// Messages var as a slice Message struct
//var Messages []Message

// Get all massages ever sent

// func GetAllMessages(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(Messages)
// }

// GetMessage single massage by its id
func (s Service) GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idS, ok := mux.Vars(r)["id"] // Gets params
	if !ok {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(idS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mes, err := s.strg.GetByID(uint(id))
	if err == errors.New("Deleted message") {
		http.Error(w, err.Error(), http.StatusGone)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonResponse, err := json.Marshal(mes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(jsonResponse))
	// Loop through massages and find one with the id from params
	// for _, mas := range Messages {
	// 	if mas.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(mas)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(&Message{})
}

func (s Service) GetUserMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//user := r.Context().Value(auth.KeyContext)

	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := s.strg.GetUserMes(bodyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res)
	fmt.Fprint(w, "")
}

func (s Service) DeleteUserMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := s.strg.DeleteUserMes(bodyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
	fmt.Fprint(w, "")
}

//Get massage by users id
// func GetUserMessage(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var mas Message
// 	_ = json.NewDecoder(r.Body).Decode(&mas)

// 	for _, i := range Messages {
// 		if mas.Author.Firstname == i.To {
// 			json.NewEncoder(w).Encode(i)
// 			return
// 		}
// 	}
//json.NewEncoder(w).Encode(&Message{}) //do not uncomment this line
//}

// WriteMessage adds new massage
func (s Service) WriteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	/*user := r.Context().Value(auth.KeyContext)*/ //edited

	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//usr, err := strconv.ParseUint(user.(string), 2, 36)
	s.strg.Write(bodyReq /*, uint(usr)*/) //edited
	fmt.Fprint(w, "")
}

// UpdateMessage updates your sent massage
func (s Service) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idS, ok := mux.Vars(r)["id"] // Gets params
	if !ok {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(idS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mes, err := s.strg.Update(bodyReq, uint(id))
	if err == errors.New("Deleted message1") {
		http.Error(w, err.Error(), http.StatusGone)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonResponse, err := json.Marshal(mes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(jsonResponse))

	// params := mux.Vars(r)
	// for index, mas := range Messages {
	// 	if mas.ID == params["id"] {
	// 		Messages = append(Messages[:index], Messages[index+1:]...)
	// 		var mas Message
	// 		_ = json.NewDecoder(r.Body).Decode(&mas)
	// 		mas.ID = params["id"]
	// 		Messages = append(Messages, mas)
	// 		json.NewEncoder(w).Encode(mas)
	// 		return
	// 	}
	// }
}

// DeleteMessage  deletes your massage
func (s Service) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idS, ok := mux.Vars(r)["id"] // Gets params
	if !ok {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(idS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mes, err := s.strg.DeleteByID(uint(id))
	if err == errors.New("Deleted message") {
		http.Error(w, err.Error(), http.StatusGone)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonResponse, err := json.Marshal(mes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(jsonResponse))
	// params := mux.Vars(r)
	// for index, mas := range Messages {
	// 	if mas.ID == params["id"] {
	// 		Messages = append(Messages[:index], Messages[index+1:]...)
	// 		break
	// 	}
	// }
	// json.NewEncoder(w).Encode(Messages)
}
