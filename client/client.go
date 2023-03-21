package main

import (
	"calculator/calculatorpb"
	"context"
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

	callSum(client)
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
