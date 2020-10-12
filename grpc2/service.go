package grpc2

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Service struct {
	MessageStorage Storage
}

type Message struct {
	ID        uint   `json:"id"`
	In        string `json:"in"`
	Content   string `json:"content"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	//Author    *Author `json:"author"`
	To string `json:"name"`
}

func (s *Service) SetStorage(storage Storage) {
	s.MessageStorage = storage
}

func (s *Service) GetStorage() Storage {
	return s.MessageStorage
}

func (s *Service) Write(jsonString []byte /*, user uint*/) { //edited
	mes := &Message{}
	err := json.Unmarshal(jsonString, mes)
	//mes.ID = uint(rand.Intn(100))

	if err == nil {
		fmt.Print("Invalid json format")
	}

	s.GetStorage().Add(Message{ID: mes.ID /*user*/, Content: mes.Content, Firstname: mes.Firstname, To: mes.To})

}

func (s *Service) Update(jsonString []byte, id uint) (Message, error) {
	mes := &Message{}
	err := json.Unmarshal(jsonString, mes)

	mes.ID = id
	el, err := s.GetStorage().Update(*mes, id)

	if err != nil {
		return Message{}, errors.New("Not found")
	}
	return *el, nil
}

func (s *Service) DeleteByID(id uint) (Message, error) {
	mes, err := s.GetStorage().Delete(id)

	if err != nil {
		return Message{}, errors.New("Deleted message")
	}
	return *mes, nil
}

func (s *Service) GetByID(id uint) (Message, error) {
	mes, err := s.GetStorage().Get(id)

	if err != nil {
		return Message{}, errors.New("Deleted message")
	}
	return *mes, nil
}

func (s *Service) GetUserMes(jsonString []byte) ([]Message, error) {
	//res := []*Message{}
	mes := &Message{}
	err := json.Unmarshal(jsonString, mes)

	res, err := s.GetStorage().GetUserMessage(*mes)
	if err != nil {
		return []Message{}, errors.New("Inbox is clear")
	}
	return res, nil
}

//DeleteUserMes provides input message to storage service layer
func (s *Service) DeleteUserMes(jsonString []byte) ([]Message, error) {
	mes := &Message{}
	err := json.Unmarshal(jsonString, mes)

	getmsgs, err := s.GetStorage().GetUserMessage(*mes)
	if err != nil {
		return []Message{}, errors.New("Error with fetching data")
	}

	res, err := s.GetStorage().DeleteUserMessage(*mes, getmsgs)
	if err != nil {
		return []Message{}, errors.New("Message does not exist")
	}

	return res, nil
}
