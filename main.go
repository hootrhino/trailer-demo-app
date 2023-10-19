package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"trailer-demo-app/trailer"

	"google.golang.org/grpc"
)

type testRpcServer struct {
	trailer.UnimplementedTrailerServer
}

func (testRpcServer) Init(ctx context.Context, cfg *trailer.Config) (*trailer.Response, error) {
	log.Println("来自协议包的日志 Init, Config=", string(cfg.GetKv()))
	return &trailer.Response{}, nil
}
func (testRpcServer) Start(context.Context, *trailer.Request) (*trailer.Response, error) {
	log.Println("来自协议包的日志 Start")
	return &trailer.Response{}, nil
}
func (testRpcServer) Status(context.Context, *trailer.Request) (*trailer.StatusResponse, error) {
	log.Println("来自协议包的日志 Status")
	return &trailer.StatusResponse{Status: trailer.StatusResponse_RUNNING, Message: "Success"}, nil
}
func (testRpcServer) Service(ctx context.Context, request *trailer.ServiceRequest) (*trailer.ServiceResponse, error) {
	if string(request.Args) == "query" {
		return &trailer.ServiceResponse{Data: []byte{}}, nil
	}
	return &trailer.ServiceResponse{}, nil
}
func (testRpcServer) Schema(ctx context.Context, req *trailer.SchemaRequest) (*trailer.SchemaResponse, error) {
	log.Println("来自协议包的日志 Schema")
	Columns := []*trailer.Column{
		{Name: []byte("temp"), Type: trailer.ValueType_NUMBER, Description: []byte("温度")},
		{Name: []byte("humi"), Type: trailer.ValueType_NUMBER, Description: []byte("湿度")},
		{Name: []byte("co2"), Type: trailer.ValueType_NUMBER, Description: []byte("二氧化碳")},
		{Name: []byte("weather"), Type: trailer.ValueType_STRING, Description: []byte("天气")},
		{Name: []byte("isOk"), Type: trailer.ValueType_BOOL, Description: []byte("你还好吗")},
	}
	return &trailer.SchemaResponse{Columns: Columns}, nil
}
func (testRpcServer) Query(ctx context.Context, req *trailer.DataRowsRequest) (*trailer.DataRowsResponse, error) {
	log.Println("来自协议包的日志 Query", string(req.GetQuery()))
	// [
	//     {
	//         "co2": 13.5,
	//         "humi": 65,
	//         "isOk": false,
	//         "temp": 15.34,
	//         "weather": "SUNNY"
	//     }
	// ]
	Values1 := []*trailer.ColumnValue{
		{Name: []byte("temp"), Type: trailer.ValueType_NUMBER, Value: []byte("15.34")},
		{Name: []byte("humi"), Type: trailer.ValueType_NUMBER, Value: []byte("65")},
		{Name: []byte("co2"), Type: trailer.ValueType_NUMBER, Value: []byte("13.5")},
		{Name: []byte("weather"), Type: trailer.ValueType_STRING, Value: []byte("SUNNY")},
		{Name: []byte("isOk"), Type: trailer.ValueType_BOOL, Value: []byte("false")},
	}
	Values2 := []*trailer.ColumnValue{
		{Name: []byte("temp"), Type: trailer.ValueType_NUMBER, Value: []byte("15.34")},
		{Name: []byte("humi"), Type: trailer.ValueType_NUMBER, Value: []byte("65")},
		{Name: []byte("co2"), Type: trailer.ValueType_NUMBER, Value: []byte("13.5")},
		{Name: []byte("weather"), Type: trailer.ValueType_STRING, Value: []byte("SUNNY")},
		{Name: []byte("isOk"), Type: trailer.ValueType_BOOL, Value: []byte("false")},
	}
	Rows := []*trailer.DataRow{
		{Column: Values1},
		{Column: Values2},
	}
	return &trailer.DataRowsResponse{
		Row: Rows,
	}, nil
}
func (testRpcServer) Stop(context.Context, *trailer.Request) (*trailer.Response, error) {
	log.Println("来自协议包的日志 Stop")
	return &trailer.Response{}, nil
}
func main() {
	port := flag.Int("port", 7700, "port")
	flag.Parse()
	fmt.Printf("来自协议包的日志, main 参数 %v\n", os.Args)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("来自协议包的日志 failed to listen: %v", err)
	}
	log.Println("来自协议包的日志 Listen at localhost:", *port)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	trailer.RegisterTrailerServer(grpcServer, testRpcServer{})
	grpcServer.Serve(lis)
	log.Println("来自协议包的日志 Stop")

}
