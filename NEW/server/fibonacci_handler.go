package server

import (
	"context"
	"fmt"
	"training/grpcnew/fibonaccipb/proto"
)

type Fibonacci_Handler struct{}

func (ch *Fibonacci_Handler) FibonacciSeries(ctx context.Context, request *proto.Request) (*proto.FibonacciResponse, error) {
	response := &proto.FibonacciResponse{}

	var n int32
	fmt.Println("Enter Number")
	fmt.Scan(&n)

	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; int32(i) <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	response.Result = int32(f[n])

	return response, nil
}
