package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"proto/product"
)

type server struct {
}

func (*server) List(_ context.Context, request *product.Empty) (*product.ProductPayload, error) {
	sp := product.ProductPayload{
		Name:  "Bánh xèo",
		Price: "20.000đ",
	}
	return &sp, nil
}

func main()  {
	address := "0.0.0.0:50051"

	listen, err := net.Listen("tcp",address)
	if err != nil {
		log.Fatal("Error %v", err)
	}

	fmt.Printf("Server is listening on %v ...", address)

	s:= grpc.NewServer()

	product.RegisterProductServiceServer(s, &server{})

	s.Serve(listen)
}