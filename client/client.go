package main

import (
	"calculator/calculatorpb"
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50069", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("err while dial %v", err)
	}

	// close connection <after return>
	defer cc.Close()

	client := calculatorpb.NewCalculatorServiceClient(cc)

	// callSum(client)
	callPND(client)
}

func callSum(c calculatorpb.CalculatorServiceClient) {
	log.Printf("Calling sum api")
	resp, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		Num1: 7,
		Num2: 6,
	})

	if err != nil {
		log.Fatalf("Call sum api error: %v", err)
	}

	log.Printf("sum api response: %v\n", resp.GetResult())
}

func callPND(c calculatorpb.CalculatorServiceClient) {
	stream, err := c.PrimeNumberDecomposition(context.Background(), &calculatorpb.PNDRequest{
		Number: 120,
	})

	if err != nil {
		log.Fatalf("Call PND failed: %v", err)
	}

	for {
		resp, recvErr := stream.Recv()
		if recvErr == io.EOF {
			log.Printf("server finished streaming")
			return
		}
		log.Printf("Prime number: %v", resp.GetResult())
	}

}
