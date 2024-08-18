package main

import (
	"Client/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"strconv"
	"time"
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
	for i := 0; i < count; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		defer cancel()
		response, e := c.GenerateTextEN(ctx, &proto.TextGenerateInfo{Count: int32(n)})
		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error occurred:%s\n", e.Error())
		} else {

			fmt.Printf("Text: %s\n", response.GetText())
		}
		response, e = c.GenerateTextTR(ctx, &proto.TextGenerateInfo{Count: int32(n)})
		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error occurred:%s\n", e.Error())
		} else {

			fmt.Printf("Yazı: %s\n", response.GetText())
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, e := c.GetTextBound(ctx, &proto.NoParam{})
	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error occurred:%s\n", e.Error())
	} else {
		fmt.Printf("Min: %d, Max:%d\n", response.MinCount, response.MaxCount)
	}
}
