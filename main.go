package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"trailer-demo-app/trailer"

	"google.golang.org/grpc"
)

type testRpcServer struct {
	trailer.UnimplementedTrailerServer
}

func (testRpcServer) Init(context.Context, *trailer.Config) (*trailer.Response, error) {
	log.Println("Init")
	return &trailer.Response{}, nil
}
func (testRpcServer) Start(context.Context, *trailer.Request) (*trailer.Response, error) {
	log.Println("Start")
	return &trailer.Response{}, nil
}
func (testRpcServer) Status(context.Context, *trailer.Request) (*trailer.Response, error) {
	log.Println("Status")
	return &trailer.Response{Code: 200}, nil
}
func (testRpcServer) Service(ctx context.Context, request *trailer.ServiceRequest) (*trailer.ServiceResponse, error) {
	if string(request.Args) == "query" {
		return &trailer.ServiceResponse{Data: []byte{}}, nil
	}
	return &trailer.ServiceResponse{}, nil
}
func (testRpcServer) Schema(context.Context, *trailer.SchemaRequest) (*trailer.SchemaResponse, error) {
	log.Println("Schema")
	Columns := []*trailer.Column{
		{Name: []byte("temp"), Type: []byte("NUMBER")},
		{Name: []byte("humi"), Type: []byte("NUMBER")},
		{Name: []byte("co2"), Type: []byte("NUMBER")},
		{Name: []byte("weather"), Type: []byte("STRING")},
	}
	return &trailer.SchemaResponse{Columns: Columns}, nil
}
func (testRpcServer) Stop(context.Context, *trailer.Request) (*trailer.Response, error) {
	log.Println("Stop")
	return &trailer.Response{}, nil
}
func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 7798))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("net.Listen at localhost:7798")
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	trailer.RegisterTrailerServer(grpcServer, testRpcServer{})
	grpcServer.Serve(lis)
}
