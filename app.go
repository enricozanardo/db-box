package main

import (
	"net"
	"google.golang.org/grpc"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"github.com/onezerobinary/db-box/mygrpc"
	"github.com/goinggo/tracelog"
	"os"
	"github.com/spf13/viper"
	"github.com/onezerobinary/db-box/job"
)

const (
	GRPC_PORT = ":1982"
)

func main() {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	//development environment
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		tracelog.Errorf(err, "main", "main", "Error reading config file")
	}

	tracelog.Warning("main", "main", "Using config file")

	// Start the cronjob that must check the status of the accounts
	go job.CheckAccountStatus()

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