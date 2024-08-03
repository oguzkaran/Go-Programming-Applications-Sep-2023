package main

import (
	"Server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GreetingServer struct {
	proto.UnimplementedGreetingServiceServer
}

func (s *GreetingServer) HelloTR(context context.Context, req *proto.GreetingRequest) (*proto.GreetingResponse, error) {
	ms := "Bekar"

	if req.Married {
		ms = "Evli"
	}

	return &proto.GreetingResponse{Message: fmt.Sprintf("Merhaba %s %s", req.FirstName, req.LastName), MaritalStatus: ms}, nil
}

func (s *GreetingServer) HelloEN(context context.Context, req *proto.GreetingRequest) (*proto.GreetingResponse, error) {
	ms := "Single"

	if req.Married {
		ms = "Married"
	}

	return &proto.GreetingResponse{Message: fmt.Sprintf("Hello %s %s", req.FirstName, req.LastName), MaritalStatus: ms}, nil
}

func main() {
	listener, e := net.Listen("tcp", ":50500")

	if e != nil {
		log.Fatalf("Problem occurred:%s\n", e.Error())
	}

	s := grpc.NewServer()
	proto.RegisterGreetingServiceServer(s, &GreetingServer{})
	log.Printf("Serving on localhost:%s", listener.Addr())
	if e = s.Serve(listener); e != nil {
		log.Fatalf("Failed to serve: %s\n", e.Error())
	}
}
