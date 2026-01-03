package main

import (
	"fmt"
	pb "github.com/HamsterNiki/test-grpc-with-docker/one/generated"
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
