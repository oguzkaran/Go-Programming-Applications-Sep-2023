package main

import (
	"Client/proto"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		_, _ = fmt.Fprintf(os.Stderr, "usage: client <address>")
		os.Exit(1)
	}

	n, e := strconv.Atoi(os.Args[2])
	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Invalid value for n")
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

	streamEN, e := c.GenerateAndJoinTextsEN(context.Background())

	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error occurred GenerateAndJoinTextsEN:%s\n", e.Error())
		os.Exit(1)
	}

	streamTR, e := c.GenerateAndJoinTextsTR(context.Background())

	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error occurred GenerateAndJoinTextsTR:%s\n", e.Error())
		os.Exit(1)
	}

	for i := 0; i < n; i++ {
		count := int32(rand.Intn(7) + 3)

		log.Printf("Count:%d", count)
		if e := streamEN.Send(&proto.TextGenerateInfo{Count: count}); e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error occurred Send EN:%s\n", e.Error())
			os.Exit(1)
		}

		if e := streamTR.Send(&proto.TextGenerateInfo{Count: count}); e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error occurred SendTR:%s\n", e.Error())
			os.Exit(1)
		}
	}

	resEN, e := streamEN.CloseAndRecv()

	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error occurred CloseAndRecv EN:%s\n", e.Error())
		os.Exit(1)
	}
	fmt.Printf("Text:%s\n", resEN.Text)

	resTR, e := streamTR.CloseAndRecv()

	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error occurred CloseAndRecv TR:%s\n", e.Error())
		os.Exit(1)
	}
	fmt.Printf("Yazı:%s\n", resTR.Text)
}
