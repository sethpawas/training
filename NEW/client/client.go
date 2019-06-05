package client

import (
	"fmt"
	"github.com/sethpawas/training/NEW/calculatorpb/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

func Start(port string) {
	cc, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	// SERVER STREAMING
	fmt.Println("Fibonacci Series")
	n = 120
	stream, err := client.FibonacciSeries(ctx, &calculatorpb.CalculatorRequest{
		Number: n,
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		resPF, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(resPF.GetResult())
		}
	}
}
