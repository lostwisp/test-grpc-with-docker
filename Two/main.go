package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/lostwisp/test-grpc-with-docker/pr"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedScoreServer // Встраиваем для совместимости
}

func (s *server) UpdScore(ctx context.Context, req *pb.ScoreRequest) (*pb.ScoreResponse, error) {
	// Бизнес-логика: например, обновляем счёт (здесь просто +1)

	newScore := req.Score + 1
	fmt.Println(newScore)
	log.Printf("Received score: %d, returning: %d", req.Score, newScore)
	return &pb.ScoreResponse{Score: newScore}, nil
}

func main() {
	// Создаём TCP-listener
	fmt.Println("gRPC server starting")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Создаём gRPC-сервер
	grpcServer := grpc.NewServer()

	// Регистрируем сервис
	pb.RegisterScoreServer(grpcServer, &server{})

	// Запускаем сервер
	log.Println("gRPC server listening on :50051")
	fmt.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
