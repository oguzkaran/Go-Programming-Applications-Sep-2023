package main

import (
	"Server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type GreetingServer struct {
	proto.UnimplementedGreetingServiceServer
}

func (s *GreetingServer) HelloTR(context context.Context, req *proto.GreetingRequest) (*proto.GreetingResponse, error) {
	ms := "Bekar"

	if req.Married {
		ms = "Evli"
	}

	return &proto.GreetingResponse{Id: req.Id + 1, Message: fmt.Sprintf("Merhaba %s %s", req.FirstName, req.LastName), MaritalStatus: ms}, nil
}

func (s *GreetingServer) HelloEN(context context.Context, req *proto.GreetingRequest) (*proto.GreetingResponse, error) {
	ms := "Single"

	if req.Married {
		ms = "Married"
	}

	return &proto.GreetingResponse{Id: req.Id + 2, Message: fmt.Sprintf("Hello %s %s", req.FirstName, req.LastName), MaritalStatus: ms}, nil
}

func main() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "usage: server <address>")
		os.Exit(1)
	}

	listener, e := net.Listen("tcp", os.Args[1])

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
