package main

import (
	"time"
	"training/NEW/client"
	"training/NEW/server"
)

func serv() {

}
func main() {

	go func() {
		server.Start("50051")
	}()
	time.Sleep(2 * time.Second)
	client.Start("50051")
}
