package main

import (
	service "calculator/calculatorpb"
	"context"
	"fmt"
	"log"
	"net"
	"time"

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

/*
*	Alogrithm
*
 */

func (*server) PrimeNumberDecomposition(req *service.PNDRequest, stream service.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Printf("PND called")
	k := int32(2)
	N := req.GetNumber()
	log.Printf("Number: %d", N)
	for N > 1 {
		if N%k == 0 {
			N = N / k
			log.Printf("New Number: %d", N)
			// send response to client
			stream.Send(&service.PNDResponse{
				Result: k,
			})
			time.Sleep(2000 * time.Millisecond)
		} else {
			k++
			log.Printf("k increase to %d", k)
		}
	}

	return nil
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
