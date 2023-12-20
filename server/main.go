package main

import (
	"context"
	"fmt"
	"net"

	"github.com/tylerjgabb/go-grpc-sandbox/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(
	ctx context.Context,
	in *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	fmt.Printf("ADD: %d+%d\n", in.A, in.B)
	return &pb.CalculationResponse{
		Result: in.A + in.B,
	}, nil
}

func (s *server) Divide(
	ctx context.Context,
	in *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	fmt.Printf("DIVIDE: %d/%d\n", in.A, in.B)
	if in.B == 0 {
		fmt.Printf("DIVIDE: cannot divide by zero\n")
		return nil, status.Error(codes.InvalidArgument, "cannot divide by zero")
	}
	return &pb.CalculationResponse{
		Result: in.A / in.B,
	}, nil
}

func (s *server) Sum(
	ctx context.Context,
	in *pb.NumbersRequest,
) (*pb.CalculationResponse, error) {
	fmt.Printf("SUM: %v\n", in.Numbers)
	var sum int64
	for _, v := range in.Numbers {
		sum += v
	}
	return &pb.CalculationResponse{
		Result: sum,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterCalculatorServer(s, &server{})
	fmt.Printf("Starting server on port 8080\n")
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
