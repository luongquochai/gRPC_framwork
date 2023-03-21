package main

import (
	service "calculator/calculatorpb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *service.SumRequest) (*service.SumResponse, error) {
	log.Printf("Sum called")
	resp := &service.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}

	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50069")
	if err != nil {
		log.Fatalf("Could not listen %v", lis)
	}
	s := grpc.NewServer()

	service.RegisterCalculatorServiceServer(s, &server{})
	fmt.Println("Calculator is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("Could not serve with err %v", err)
	}

}
