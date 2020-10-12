package grpc2

import (
	"errors"
)

type Store struct {
	messages []*Message
}

func NewStore() *Store {
	return &Store{
		messages: make([]*Message, 0, 20),
	}
}

//GetUserMesaage gets all messages which were sent to concrete user

func (m Store) Get(id uint) (*Message, error) {

	if m.messages[id] == nil {
		return nil, errors.New("message was deleted")
	}
	return m.messages[id], nil
}

func (m *Store) Add(c Message) *Message {
	m.messages = append(m.messages, &c)
	mm := *m.messages[len(m.messages)-1]

	return &mm
}

func (m Store) Delete(id uint) (*Message, error) {
	if m.messages[id] == nil {
		return nil, errors.New("message was deleted")
	}
	el := m.messages[id]

	m.messages[id] = nil
	return el, nil
}

// DeleteUserMessage deletes concrete user message
func (m Store) DeleteUserMessage(msg Message, msgs []Message) ([]Message, error) {
	//var id uint
	var res []Message
	for i, v := range msgs {
		if msg.To == v.Firstname {
			//id = v.ID
			//res = append(res, *m.messages[id])
			m.messages = append(m.messages[:i], m.messages[i+1:]...)
			res = append(res, *m.messages[i])
			break
		}
		//m.messages[id] = nil
	}
	return res, nil
}

func (m Store) Update(c Message, id uint) (*Message, error) {
	// if id >= m.Len() {
	// 	return nil, errors.New("invalid id")
	// }
	if m.messages[id] == nil {
		return nil, errors.New("message was deleted")
	}
	m.messages[id] = &c
	return m.messages[id], nil
}

func (m Store) GetAll() []*Message {
	result := make([]*Message, m.Len())
	copy(result, m.messages)
	return result
}

func (m Store) Len() uint {
	return uint(len(m.messages))
}

//GetUserMessage returns users Messages
func (m Store) GetUserMessage(msg Message) ([]Message, error) {
	var res []Message
	for _, v := range m.messages {
		if msg.Firstname == v.To {
			res = append(res, *v)
		}
	}
	return res, nil
}
