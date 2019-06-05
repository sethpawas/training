package server

import (
	"context"
	"fmt"
	"training/grpcnew/fibonaccipb/proto"
)

type Fibonacci_Handler struct{}

func (ch *Fibonacci_Handler) FibonacciSeries(ctx context.Context, request *proto.Request) (*proto.FibonacciResponse, error) {
	response := &proto.FibonacciResponse{}

	var a, b int32
	fmt.Println("Enter Two numbers")
	fmt.Scan(&a, &b)

	response.Result = a * b

	return response, nil
}
