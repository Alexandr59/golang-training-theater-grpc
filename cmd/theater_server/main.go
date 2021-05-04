package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	//"github.com/Alexandr59/golang-training-theater-grpc/pkg/db"
	"golang-training-theater-grpc/pkg/db"
	//"github.com/Alexandr59/golang-training-theater-grpc/api"
	"golang-training-theater-grpc/api"
)

var (
	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")
)

func init() {
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "Theater_db"
	}
	if password == "" {
		password = "5959"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
}

func main() {
	conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}
	listener, err := net.Listen("tcp", ":8181")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	//pb.RegisterAccountServiceServer(server, api.NewAccountServer(*data.NewAccountData(conn)))
	api.RegisterAllServiceServer(server, conn)
	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
