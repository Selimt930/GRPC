package grpcsrv

import (
	"MailService/grpc2"
	"MailService/grpc2/proto"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPCStore ...
type GRPCStore struct {
	proto.UnimplementedServiceStorageServer
	storage grpc2.Storage
}

// InitGRPCStore ...
func InitGRPCStore(storage grpc2.Storage) *GRPCStore {
	return &GRPCStore{
		storage: storage,
	}
}

//Get ...
func (g *GRPCStore) Get(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	res, err := g.storage.Get(uint(req.ID))
	if err != nil {
		return nil, status.New(codes.Unknown, "Error occured").Err()
	}

	return &proto.Response{
		Elements: []*proto.Message{
			{
				ID:        uint32(res.ID),
				Content:   res.Content,
				Firstname: res.Firstname,
				To:        res.To,
			},
		},
	}, nil
}

//Write ...
func (g *GRPCStore) Write(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	mes := req.Element
	res := g.storage.Add(grpc2.Message{ID: uint(mes.ID), Firstname: mes.Firstname, Content: mes.Content, To: mes.To})

	return &proto.Response{
		Elements: []*proto.Message{
			{
				ID:        uint32(res.ID),
				Content:   res.Content,
				Firstname: res.Firstname,
				To:        res.To,
			},
		},
	}, nil
}

//Update ...
func (g *GRPCStore) Update(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	mes := req.Element
	res, err := g.storage.Update(grpc2.Message{
		ID:        uint(mes.ID),
		Content:   mes.Content,
		Firstname: mes.Firstname,
		To:        mes.To,
	}, uint(req.ID))
	if err != nil {
		return nil, status.New(codes.Unknown, err.Error()).Err()
	}

	return &proto.Response{
		Elements: []*proto.Message{
			{
				ID:        uint32(res.ID),
				Content:   res.Content,
				Firstname: res.Firstname,
				To:        res.To,
			},
		},
	}, nil
}

//Delete ...
func (g *GRPCStore) Delete(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	res, err := g.storage.Delete(uint(req.ID))
	if err != nil {
		return nil, status.New(codes.Unknown, err.Error()).Err()
	}
	return &proto.Response{
		Elements: []*proto.Message{
			{
				ID:        uint32(res.ID),
				Content:   res.Content,
				Firstname: res.Firstname,
				To:        res.To,
			},
		},
	}, nil
}

//DeleteUserMes provides data to storage service layer
func (g *GRPCStore) DeleteUserMes(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	mes := req.Element
	result := make([]*proto.Message, 0, 20)
	resmsgs, err := g.storage.GetUserMessage(grpc2.Message{
		ID:        uint(mes.ID),
		Content:   mes.Content,
		Firstname: mes.Firstname,
		To:        mes.To,
	})
	if err != nil {
		return nil, status.New(codes.Unknown, err.Error()).Err()
	}

	res, err := g.storage.DeleteUserMessage(grpc2.Message{
		ID:        uint(mes.ID),
		Content:   mes.Content,
		Firstname: mes.Firstname,
		To:        mes.To,
	}, resmsgs)
	if err != nil {
		return nil, status.New(codes.Unknown, err.Error()).Err()
	}

	for _, v := range res {
		result = append(result, &proto.Message{
			ID:        uint32(v.ID),
			Content:   v.Content,
			Firstname: v.Firstname,
			To:        v.To,
		})
	}
	return &proto.Response{
		Elements: result,
	}, nil
}

//GetAll ...
func (g *GRPCStore) GetAll(ctx context.Context, empty *empty.Empty) (*proto.Response, error) {
	messages := g.storage.GetAll()
	res := make([]*proto.Message, 0, len(messages))

	for _, v := range messages {
		res = append(res, &proto.Message{
			ID:        uint32(v.ID),
			Content:   v.Content,
			Firstname: v.Firstname,
			To:        v.To,
		})
	}
	return &proto.Response{
		Elements: res,
	}, nil
}

//GetUserMes returns concrete users messages
func (g *GRPCStore) GetUserMes(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	mes := req.Element
	result := make([]*proto.Message, 0, 20)
	res, err := g.storage.GetUserMessage(grpc2.Message{
		ID:        uint(mes.ID),
		Content:   mes.Content,
		Firstname: mes.Firstname,
		To:        mes.To,
	})
	if err != nil {
		return nil, status.New(codes.Unknown, err.Error()).Err()
	}

	for _, v := range res {
		result = append(result, &proto.Message{
			ID:        uint32(v.ID),
			Content:   v.Content,
			Firstname: v.Firstname,
			To:        v.To,
		})
	}
	return &proto.Response{
		Elements: result,
	}, nil

}
