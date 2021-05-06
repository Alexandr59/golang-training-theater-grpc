package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/Alexandr59/golang-training-theater-grpc/api"
)

var (
	serverAddr = os.Getenv("product-server")
	listen     = os.Getenv("LISTEN")
)

func init() {
	if serverAddr == "" {
		serverAddr = "localhost:8080"
	}
	if listen == "" {
		listen = "localhost:8181"
	}
}

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcMux := runtime.NewServeMux()
	api.RegisterAllServiceHandler(context.Background(), grpcMux, conn)
	log.Fatal(http.ListenAndServe(listen, grpcMux))
}
