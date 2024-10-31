package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "projgRPC/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	// Adicionar um produto
	product := &pb.Product{Id: 1, Name: "Laptop", Price: 1200.00}
	res, err := client.AddProduct(context.Background(), &pb.ProductRequest{Product: product})
	if err != nil {
		log.Fatalf("AddProduct error: %v", err)
	}
	fmt.Printf("Product added: %v\n", res.GetProduct())

	// Buscar um produto
	time.Sleep(1 * time.Second)
	res, err = client.GetProduct(context.Background(), &pb.ProductID{Id: 1})
	if err != nil {
		log.Fatalf("GetProduct error: %v", err)
	}
	fmt.Printf("Product fetched: %v\n", res.GetProduct())

	// Atualizar o produto
	product.Price = 1100.00
	res, err = client.UpdateProduct(context.Background(), &pb.ProductRequest{Product: product})
	if err != nil {
		log.Fatalf("UpdateProduct error: %v", err)
	}
	fmt.Printf("Product updated: %v\n", res.GetProduct())

	// Deletar o produto
	delRes, err := client.DeleteProduct(context.Background(), &pb.ProductID{Id: 1})
	if err != nil {
		log.Fatalf("DeleteProduct error: %v", err)
	}
	fmt.Printf("Product deleted: %s\n", delRes.GetMessage())
}
