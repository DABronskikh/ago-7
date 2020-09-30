package main

import (
	"context"
	fineV1Pb "github.com/DABronskikh/ago-7/pkg/fine/v1"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const defaultPort = "9998"
const defaultHost = "0.0.0.0"

func main() {
	port, ok := os.LookupEnv("APP_PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("APP_HOST")
	if !ok {
		host = defaultHost
	}

	if err := execute(net.JoinHostPort(host, port)); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func execute(addr string) (err error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			if err == nil {
				err = cerr
				return
			}
			log.Print(err)
		}
	}()

	client := fineV1Pb.NewAggregatorServiceClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 15)

	stream, err := client.SearchFlights(ctx, &fineV1Pb.SearchData{
		FromIATA:      1111,
		ToIATA:        2222,
	})
	if err != nil {
		return err
	}

	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		log.Print("response: ", response)
	}

	log.Print("finished")
	return nil
}
