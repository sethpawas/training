package main

import (
	"training/grpcnew/client"
	"training/grpcnew/server"
)

func main() {

	go func() {
		server.Start("50051")
	}()

	//time.Sleep(2 * time.Second)

	client.Start("50051")
}
