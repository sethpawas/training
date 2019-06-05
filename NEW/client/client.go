package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	fibonaccipb "training/NEW/fibonaccipb/proto"
)

func Start(port string) {
	cc, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	ctx := context.Background()
	client := fibonaccipb.NewFibonacciServiceClient(cc)

	// SERVER STREAMING
	fmt.Println("Fibonacci Series")
	n := 120
	stream, err := client.FibonacciSeries(ctx, &fibonaccipb.Request{
		Number: int32(n),
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
