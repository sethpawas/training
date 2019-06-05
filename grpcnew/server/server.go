package server

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"training/grpcnew/fibonaccipb/proto"
)

func Start(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &Fibonacci_Handler{})

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
