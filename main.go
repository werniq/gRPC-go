package main

import (
	"context"
	"github.com/werniq/grpc-todo/invoicer/invoicer"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponseMessage, error) {
	return &invoicer.CreateResponseMessage{
		Pdf:  []byte("teset"),
		Docx: []byte("teset"),
	}, nil
}

// firstly, we need to create a server, so we can recieve request, and send responses
func main() {
	listener, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}

	// service registrar, server
	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	if err := serverRegistrar.Serve(listener); err != nil {
		log.Fatalf("Impossible to serve: %v", err)
	}
}
