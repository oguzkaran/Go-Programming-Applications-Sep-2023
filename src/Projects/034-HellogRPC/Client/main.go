package main

import (
	"Client/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "usage: client <address>")
		os.Exit(1)
	}
	con, e := grpc.NewClient(os.Args[1], grpc.WithTransportCredentials(insecure.NewCredentials()))

	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Can not connect to server")
		os.Exit(1)
	}

	defer func(con *grpc.ClientConn) {
		_ = con.Close()
	}(con)

	c := proto.NewGreetingServiceClient(con)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	response, e := c.HelloTR(ctx, &proto.GreetingRequest{Id: 1, FirstName: "Oğuz", LastName: "Karan", Married: true})
	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Can not connect to server")
		os.Exit(1)
	}

	fmt.Printf("%d, %s, %s\n", response.Id, response.Message, response.MaritalStatus)
	response, e = c.HelloEN(ctx, &proto.GreetingRequest{Id: 1, FirstName: "Oğuz", LastName: "Karan", Married: true})
	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Can not connect to server")
		os.Exit(1)
	}

	fmt.Printf("%d, %s, %s\n", response.Id, response.Message, response.MaritalStatus)
}
