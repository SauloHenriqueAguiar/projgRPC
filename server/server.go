package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"

	pb "projgRPC/proto"

	"google.golang.org/grpc"
)

type productServer struct {
	pb.UnimplementedProductServiceServer
	products map[int32]*pb.Product
	mu       sync.Mutex
}

func (s *productServer) AddProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	product := req.GetProduct()
	s.products[product.GetId()] = product
	return &pb.ProductResponse{Product: product}, nil
}

func (s *productServer) GetProduct(ctx context.Context, req *pb.ProductID) (*pb.ProductResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	product, exists := s.products[req.GetId()]
	if !exists {
		return nil, errors.New("product not found")
	}
	return &pb.ProductResponse{Product: product}, nil
}

func (s *productServer) UpdateProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	product := req.GetProduct()
	if _, exists := s.products[product.GetId()]; !exists {
		return nil, errors.New("product not found")
	}
	s.products[product.GetId()] = product
	return &pb.ProductResponse{Product: product}, nil
}

func (s *productServer) DeleteProduct(ctx context.Context, req *pb.ProductID) (*pb.DeleteResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.products[req.GetId()]; !exists {
		return nil, errors.New("product not found")
	}
	delete(s.products, req.GetId())
	return &pb.DeleteResponse{Message: "product deleted successfully"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &productServer{products: make(map[int32]*pb.Product)})

	fmt.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
