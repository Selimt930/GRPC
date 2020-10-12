package main

import (
	"MailService/grpc2"
	"MailService/grpc2/grpcsrv"
	"MailService/grpc2/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	proto.RegisterServiceStorageServer(s, grpcsrv.InitGRPCStore(grpc2.NewStore()))
	log.Fatal(s.Serve(listen))
}
