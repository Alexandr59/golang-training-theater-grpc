package main

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/Alexandr59/golang-training-theater-grpc/api"
	"github.com/Alexandr59/golang-training-theater-grpc/pkg/db"
)

var (
	listen   = os.Getenv("LISTEN")
	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")
)

func init() {
	if listen == "" {
		listen = ":8080"
	}
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

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ctx, cancel = context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	conn, err := connectToDbWithTimeout(ctx)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}

	listener, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	api.RegisterAllServiceServer(server, conn)
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}

func connectToDbWithTimeout(ctx context.Context) (*gorm.DB, error) {
	for {
		time.Sleep(2 * time.Second)
		conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
		if err == nil {
			return conn, nil
		}
		select {
		case <-ctx.Done():
			return nil, err
		default:
			continue
		}
	}
}
