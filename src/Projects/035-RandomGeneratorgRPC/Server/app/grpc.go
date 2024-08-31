package app

import (
	"Server/csd/util/str"
	"Server/proto"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strings"
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
		return errors.New(fmt.Sprintf("count must be in [%d, %d] interval", MinCount, MaxCount))
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

func (s *RandomGeneratorService) GenerateTextsTR(req *proto.TextsGenerateInfo,
	server proto.RandomTextGeneratorService_GenerateTextsTRServer) error {

	count := int(req.Count)
	if count <= 0 {
		return errors.New("count değeri pozitif olmalıdır")
	}
	if count < MinCount || count > MaxCount {
		return errors.New(fmt.Sprintf("count [%d, %d] aralığında olmalıdır", MinCount, MaxCount))
	}

	n := int(req.N)
	if n <= 0 {
		return errors.New("N sıfırdan büyük olmalıdır")
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

func (s *RandomGeneratorService) GenerateAndJoinTextsEN(server proto.RandomTextGeneratorService_GenerateAndJoinTextsENServer) error {
	sb := strings.Builder{}

	for {
		req, e := server.Recv()

		if e == io.EOF {
			return server.SendAndClose(&proto.TextInfo{Text: sb.String()})
		}

		if e != nil {
			return e
		}

		count := int(req.Count)

		if count <= 0 {
			return errors.New("count must be positive")
		}
		if count < MinCount || count > MaxCount {
			return errors.New(fmt.Sprintf("count must be in [%d, %d] interval", MinCount, MaxCount))
		}

		text := str.GenerateRandomTextEN(count)
		log.Printf("Text:%s", text)
		sb.WriteString(text)
	}
}

func (s *RandomGeneratorService) GenerateAndJoinTextsTR(server proto.RandomTextGeneratorService_GenerateAndJoinTextsENServer) error {
	sb := strings.Builder{}

	for {
		req, e := server.Recv()

		if e == io.EOF {
			return server.SendAndClose(&proto.TextInfo{Text: sb.String()})
		}

		if e != nil {
			return e
		}

		count := int(req.Count)

		if count <= 0 {
			return errors.New("count pozitif olmalıdır")
		}
		if count < MinCount || count > MaxCount {
			return errors.New(fmt.Sprintf("count [%d, %d] aralığında olmalıdır", MinCount, MaxCount))
		}

		text := str.GenerateRandomTextTR(count)
		log.Printf("Yazı:%s", text)
		sb.WriteString(text)
	}
}

func (s *RandomGeneratorService) GenerateInt32(_ context.Context, req *proto.Int32GenerateInfo) (*proto.Int32Result, error) {
	if req.Min >= req.Bound {
		return nil, errors.New("invalid values")
	}

	return &proto.Int32Result{Value: int32(rand.Intn(int(req.Bound-req.Min)) + int(req.Min))}, nil
}
