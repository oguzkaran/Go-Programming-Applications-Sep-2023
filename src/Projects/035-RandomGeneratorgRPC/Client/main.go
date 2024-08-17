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

	c := proto.NewRandomTextGeneratorServiceClient(con)
	for i := 0; i < 50; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		defer cancel()
		response, e := c.GenerateTextEN(ctx, &proto.TextGenerateInfo{Count: 4})
		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Can not connect to server")
			os.Exit(1)
		}

		fmt.Printf("Text: %s\n", response.GetText())
		response, e = c.GenerateTextTR(ctx, &proto.TextGenerateInfo{Count: 2})
		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Can not connect to server")
			os.Exit(1)
		}

		fmt.Printf("Yazı: %s\n", response.GetText())
	}
}
