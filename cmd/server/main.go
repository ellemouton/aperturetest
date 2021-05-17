package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	pb "aperturetest/pricesrpc"

	"github.com/lightningnetwork/lnd/cert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	secure = flag.Bool("tls", false, "use tls")
)

var bookPrices = map[int]int64{
	1: 10,
	2: 20,
	3: 0,
}

func main() {
	flag.Parse()

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
	var grpcServer *grpc.Server
	if *secure {
		log.Printf("init secure gRPC server")
		certData, _, err := cert.LoadCert("./tls.cert", "./tls.key")
		if err != nil {
			log.Fatal(err)
		}

		tlsCfg := cert.TLSConfFromCert(certData)
		tlsCredentials := credentials.NewTLS(tlsCfg)
		grpcServer = grpc.NewServer(grpc.Creds(tlsCredentials))

	} else {
		log.Printf("init insecure gRPC server")
		grpcServer = grpc.NewServer()
	}

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
