package server

import (
	"github.com/perennial-go-training/training/20_grpc/calculatorpb/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Start(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &CalculatorHandler{})

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
