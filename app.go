package main

import (
	"net"
	"google.golang.org/grpc"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"github.com/onezerobinary/db-box/mygrpc"
	"github.com/goinggo/tracelog"
	"os"
)

const (
	GRPC_PORT = ":1982"
)

func main() {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	listen, err := net.Listen("tcp", GRPC_PORT)
	if err != nil {
		tracelog.Errorf(err, "app", "main", "Failed to start the service")
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	// Add to the grpcServer the Service
	pb_account.RegisterAccountServiceServer(grpcServer, &mygrpc.AccountServiceServer{})

	tracelog.Trace("app", "main", "Grpc Server Listening on port 1982")

	grpcServer.Serve(listen)
}