//Package grpcstorage implements client-side for GRPC service
package grpcstorage

import (
	"MailService/grpc2"
	"MailService/grpc2/proto"
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"
)

//Storage ...
type Storage struct {
	storage proto.ServiceStorageClient
}

//NewStorage ...
func NewStorage() *Storage {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}

	cl := proto.NewServiceStorageClient(conn)
	return &Storage{
		storage: cl,
	}
}

// Get ...
func (g Storage) Get(id uint) (*grpc2.Message, error) {
	req := &proto.Request{ID: uint32(id)}
	res, err := g.storage.Get(context.TODO(), req)

	if err != nil {
		return nil, errors.New("Error")
	}
	mes := res.Elements[0]

	return &grpc2.Message{
		ID:        uint(mes.ID),
		Content:   mes.Content,
		Firstname: mes.Firstname,
		To:        mes.To,
	}, nil
}

// Add ...
func (g Storage) Add(mes grpc2.Message) *grpc2.Message {
	req := &proto.Request{Element: &proto.Message{
		ID:        uint32(mes.ID),
		Content:   mes.Content,
		Firstname: mes.Firstname,
		To:        mes.To,
	}}

	res, err := g.storage.Write(context.TODO(), req)
	if err != nil {
		return nil
	}

	final := res.Elements[0]

	return &grpc2.Message{
		ID:        uint(final.ID),
		Firstname: final.Firstname,
		Content:   final.Content,
		To:        final.To,
	}
}

//Update ...
func (g Storage) Update(mes grpc2.Message, id uint) (*grpc2.Message, error) {
	req := &proto.Request{ID: uint32(id), Element: &proto.Message{
		ID:        uint32(mes.ID),
		Content:   mes.Content,
		Firstname: mes.Firstname,
		To:        mes.To,
	}}
	res, err := g.storage.Update(context.TODO(), req)
	if err != nil {
		return nil, errors.New("Error occured")
	}

	resms := res.Elements[0]

	return &grpc2.Message{
		ID:        uint(resms.ID),
		Content:   resms.Content,
		Firstname: resms.Firstname,
		To:        resms.To,
	}, nil
}

func (g Storage) Delete(id uint) (*grpc2.Message, error) {
	req := &proto.Request{ID: uint32(id)}
	res, err := g.storage.Delete(context.TODO(), req)
	if err != nil {
		return nil, errors.New("Error occured")
	}

	resms := res.Elements[0]

	return &grpc2.Message{
		ID:        uint(resms.ID),
		Content:   resms.Content,
		Firstname: resms.Firstname,
		To:        resms.To,
	}, nil
}

func (g Storage) GetUserMessage(mes grpc2.Message) ([]grpc2.Message, error) {
	messages := []grpc2.Message{}
	req := &proto.Request{Element: &proto.Message{
		ID:        uint32(mes.ID),
		Content:   mes.Content,
		Firstname: mes.Firstname,
		To:        mes.To,
	}}

	res, err := g.storage.GetUserMes(context.TODO(), req)
	if err != nil {
		return nil, errors.New("Error occured")
	}

	result := res.Elements
	for _, v := range result {
		messages = append(messages, grpc2.Message{
			ID:        uint(v.ID),
			Content:   v.Content,
			Firstname: v.Firstname,
			To:        v.To,
		})
	}
	return messages, nil
}

func (g Storage) DeleteUserMessage(mes grpc2.Message, msgs []grpc2.Message) ([]grpc2.Message, error) {
	messages := []grpc2.Message{}
	req := &proto.Request{Element: &proto.Message{
		ID:        uint32(mes.ID),
		Content:   mes.Content,
		Firstname: mes.Firstname,
		To:        mes.To,
	}}

	res, err := g.storage.DeleteUserMes(context.TODO(), req)
	if err != nil {
		return nil, errors.New("Error occured")
	}

	resms := res.Elements
	for _, v := range resms {
		messages = append(messages, grpc2.Message{
			ID:        uint(v.ID),
			Content:   v.Content,
			Firstname: v.Firstname,
			To:        v.To,
		})
	}
	return messages, nil
}

func (g Storage) GetAll() []*grpc2.Message {
	return nil
}

func (g Storage) Len() uint {
	return 0
}
