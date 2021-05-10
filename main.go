package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	pb "aperturetest/pricesrpc"

	"google.golang.org/grpc"
)

var bookPrices = map[int]int64{
	1: 10,
	2: 20,
	3: 0,
}

func main() {
	http.HandleFunc("/book/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "here is your book")
	})

	http.HandleFunc("/book/2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "here is another book")
	})

	http.HandleFunc("/book/3", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "here is yet another book")
	})

	go serverGRPCForever("localhost:9001")

	if err := http.ListenAndServe("localhost:9000", nil); err != nil {
		log.Fatal(err)
	}
}

func serverGRPCForever(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	pricesServer := Server{}
	grpcServer := grpc.NewServer()

	pb.RegisterPricesServer(grpcServer, pricesServer)

	log.Printf("serving gRPC server at %s", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server prices server", err)
	}
}

var _ pb.PricesServer = (*Server)(nil)

type Server struct{}

func (s Server) GetPrice(ctx context.Context, request *pb.GetPriceRequest) (*pb.GetPriceResponse, error) {
	if !strings.Contains(request.Path, "book") {
		return &pb.GetPriceResponse{}, nil
	}

	id, err := strconv.Atoi(strings.TrimLeft(request.Path, "/book/"))
	if err != nil {
		return &pb.GetPriceResponse{}, nil
	}

	price, ok := bookPrices[id]
	if !ok {
		return &pb.GetPriceResponse{}, nil
	}

	return &pb.GetPriceResponse{
		Price: price,
	}, nil
}
