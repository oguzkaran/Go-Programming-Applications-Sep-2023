package main

import (
	"Client/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		_, _ = fmt.Fprintf(os.Stderr, "usage: client <address>")
		os.Exit(1)
	}
	count, e := strconv.Atoi(os.Args[2])

	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "invalid count value!...")
		os.Exit(1)
	}

	n, e := strconv.Atoi(os.Args[3])

	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "invalid n value!...")
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

	c := proto.NewRandomTextGeneratorServiceClient(con)

	s, e := c.GenerateTextsEN(context.Background(), &proto.TextsGenerateInfo{Count: int32(count), N: int32(n)})

	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error occurred:%s\n", e.Error())
		os.Exit(1)
	}

	for {
		res, e := s.Recv()
		if e == io.EOF {
			break
		}
		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error occurred:%s\n", e.Error())
			os.Exit(1)
		}

		fmt.Printf("Text:%s\n", res.Text)
	}
}
