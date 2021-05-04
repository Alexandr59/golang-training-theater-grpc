package main

import (
	"context"
	"golang-training-theater-grpc/api"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8181", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcMux := runtime.NewServeMux()
	//err = pb.RegisterAccountServiceHandler(context.Background(), grpcMux, conn)
	//if err != nil {
	//	log.Fatal(err)
	//}
	api.RegisterAllServiceHandler(context.Background(), grpcMux, conn)
	log.Fatal(http.ListenAndServe("localhost:8080", grpcMux))
}
