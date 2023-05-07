package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"

	"github.com/iratewami/cachemeoutside/grpc_stuff"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	exitFail = 1
)

func run() error {
	/*
		cert, err := tls.LoadX509KeyPair("server-cert.pem", "server-key.pem")
		if err != nil {
			return err
		}

		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	*/
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{},
	}

	grpcServer := grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsConfig)))

	grpc_stuff.RegisterServeCacheServer(grpcServer, &grpc_stuff.Server{})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}
