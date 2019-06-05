package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"training/grpcnew/fibonaccipb/proto"
)

func Start(port string) {
	cc, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	ctx := context.Background()
	client := proto.NewCalculatorServiceClient(cc)

	// UNARY
	fmt.Println("Multiplication")
	fb, err := client.FibonacciSeries(ctx, &proto.Request{})

	fmt.Println("Multiplication of two numbers is", fb.GetResult())

}
