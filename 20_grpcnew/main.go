package main

import (
	"server"
	"time"
	"training/20_grpc/client"
)

func serv() {
	server.Start("50051")
}
func main() {
	go serv()
	time.Sleep(2 * time.Second)
	client.Start("50051")
}
