package main

import (
	"fmt"
	pb "github.com/lostwisp/test-grpc-with-docker/pr"
	"google.golang.org/grpc"
	"time"
)

type client struct {
	pd
}

func main() {
	for {
		fmt.Println("One working")
		time.Sleep(2 * time.Second)
	}
}
