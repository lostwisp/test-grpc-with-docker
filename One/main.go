package main

import (
	"context"
	"log"
	"time"

	pb "github.com/lostwisp/test-grpc-with-docker/pr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Подключаемся к серверу (без TLS для примера)
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Создаём клиента
	client := pb.NewScoreClient(conn)

	// Создаём запрос
	req := &pb.ScoreRequest{Score: 42}

	// Вызываем метод
	for {
		resp, err := client.UpdScore(context.Background(), req)
		if err != nil {
			log.Fatalf("Failed to call UpdScore: %v", err)
		}

		// Выводим ответ
		log.Printf("Received response: score = %d", resp.Score) // Ожидаемо: 43
		req = &pb.ScoreRequest{Score: resp.Score}
		time.Sleep(1 * time.Second)
	}
}
