package main

import (
	"context"
	"github.com/DABronskikh/ago-7/cmd/fine/server/app"
	fineV1Pb "github.com/DABronskikh/ago-7/pkg/fine/v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const (
	defaultPort = "9998"
	defaultHost = "0.0.0.0"

	defaultDSN = "postgres://app:pass@localhost:5432/db"
	serv2DSN   = "postgres://app:pass@localhost:5433/db"
	serv3DSN   = "postgres://app:pass@localhost:5434/db"
)

func main() {
	port, ok := os.LookupEnv("APP_PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("APP_HOST")
	if !ok {
		host = defaultHost
	}

	log.Println(host)
	log.Println(port)

	dsn, ok := os.LookupEnv("APP_DSN")
	if !ok {
		dsn = defaultDSN
	}

	dsnServ2, ok := os.LookupEnv("APP_DSN_SERV_2")
	if !ok {
		dsnServ2 = serv2DSN
	}

	dsnServ3, ok := os.LookupEnv("APP_DSN_SERV_3")
	if !ok {
		dsnServ3 = serv3DSN
	}

	if err := execute(net.JoinHostPort(host, port), dsn, dsnServ2, dsnServ3); err != nil {
		os.Exit(1)
	}
}

func execute(addr string, dsn string, dsnServ2 string, dsnServ3 string) (err error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Print(err)
		return err
	}

	poolServ2, err := pgxpool.Connect(ctx, dsnServ2)
	if err != nil {
		log.Print(err)
		return err
	}

	poolServ3, err := pgxpool.Connect(ctx, dsnServ3)
	if err != nil {
		log.Print(err)
		return err
	}

	businessSvc := app.NewService(pool, poolServ2, poolServ3)

	grpcServer := grpc.NewServer()
	server := app.NewServer(businessSvc)

	fineV1Pb.RegisterAggregatorServiceServer(grpcServer, server)

	return grpcServer.Serve(listener)
}
