package main

import (
	"Server/csd/util/str"
	"Server/proto"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"os"
)

const MinCount = 3
const MaxCount = 10

type RandomGeneratorService struct {
	proto.UnimplementedRandomTextGeneratorServiceServer
}

func (s *RandomGeneratorService) GenerateTextEN(_ context.Context, req *proto.TextGenerateInfo) (*proto.TextInfo, error) {
	count := req.Count
	if count <= 0 {
		return nil, errors.New("count must be greater than zero")
	}
	if count < MinCount || count > MaxCount {
		return nil, errors.New(fmt.Sprintf("count must be in interval [%d, %d]", MinCount, MaxCount))
	}
	text := str.GenerateRandomTextEN(int(count))
	log.Printf("Text:%s", text)

	return &proto.TextInfo{Text: text}, nil
}

func (s *RandomGeneratorService) GenerateTextTR(_ context.Context, req *proto.TextGenerateInfo) (*proto.TextInfo, error) {
	count := req.Count
	if count <= 0 {
		return nil, errors.New("count sıfırdan büyük olmalıdır")
	}
	if count < MinCount || count > MaxCount {
		return nil, errors.New(fmt.Sprintf("count [%d, %d] aralığında olmalıdır", MinCount, MaxCount))
	}
	text := str.GenerateRandomTextTR(int(count))
	log.Printf("Yazı:%s", text)

	return &proto.TextInfo{Text: text}, nil
}

func (s *RandomGeneratorService) GetTextBound(_ context.Context, _ *proto.NoParam) (*proto.TextBound, error) {
	return &proto.TextBound{MinCount: MinCount, MaxCount: MaxCount}, nil
}

func (s *RandomGeneratorService) GenerateTextsEN(req *proto.TextsGenerateInfo,
	server proto.RandomTextGeneratorService_GenerateTextsENServer) error {

	count := int(req.Count)
	if count <= 0 {
		return errors.New("count must be greater than zero")
	}
	if count < MinCount || count > MaxCount {
		return errors.New(fmt.Sprintf("count [%d, %d] aralığında olmalıdır", MinCount, MaxCount))
	}

	n := int(req.N)
	if n <= 0 {
		return errors.New("N must be greater than zero")
	}

	log.Printf("Count:%d, N:%d", count, n)

	for i := 0; i < count; i++ {
		text := str.GenerateRandomTextEN(count)
		log.Printf("Text:%s", text)
		response := &proto.TextInfo{Text: text}

		if e := server.Send(response); e != nil {
			log.Printf("Error occurred:%s", e.Error())
			return e
		}
	}

	return nil
}

func (s *RandomGeneratorService) GenerateInt32(_ context.Context, req *proto.Int32GenerateInfo) (*proto.Int32Result, error) {
	if req.Min >= req.Bound {
		return nil, errors.New("invalid values")
	}

	return &proto.Int32Result{Value: int32(rand.Intn(int(req.Bound-req.Min)) + int(req.Min))}, nil
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
	proto.RegisterRandomTextGeneratorServiceServer(s, &RandomGeneratorService{})
	log.Printf("Serving on localhost:%s", listener.Addr())
	if e = s.Serve(listener); e != nil {
		log.Fatalf("Failed to serve: %s\n", e.Error())
	}
}
