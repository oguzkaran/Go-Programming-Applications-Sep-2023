package app

import (
	"Server/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func Run() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "usage: server <address>")
		os.Exit(1)
	}

	listener, e := net.Listen("tcp", os.Args[1])

	if e != nil {
		log.Fatalf("Problem occurred:%s\n", e.Error())
	}

	s := grpc.NewServer()
	proto.RegisterRandomTextGeneratorServiceServer(s, &RandomGeneratorService{})
	log.Printf("Serving on localhost:%s", listener.Addr())
	if e = s.Serve(listener); e != nil {
		log.Fatalf("Failed to serve: %s\n", e.Error())
	}
}
