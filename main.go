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
		{Name: ("temp"), Type: trailer.ValueType_NUMBER, Description: ("温度")},
		{Name: ("humi"), Type: trailer.ValueType_NUMBER, Description: ("湿度")},
		{Name: ("co2"), Type: trailer.ValueType_NUMBER, Description: ("二氧化碳")},
		{Name: ("weather"), Type: trailer.ValueType_STRING, Description: ("天气")},
		{Name: ("点位1"), Type: trailer.ValueType_NUMBER, Description: ("点位1参数")},
		{Name: ("点位2"), Type: trailer.ValueType_NUMBER, Description: ("点位2参数")},
		{Name: ("点位3"), Type: trailer.ValueType_NUMBER, Description: ("点位3参数")},
		{Name: ("点位4"), Type: trailer.ValueType_NUMBER, Description: ("点位4参数")},
		{Name: ("点位5"), Type: trailer.ValueType_NUMBER, Description: ("点位5参数")},
		{Name: ("点位6"), Type: trailer.ValueType_NUMBER, Description: ("点位6参数")},
		{Name: ("点位7"), Type: trailer.ValueType_NUMBER, Description: ("点位7参数")},
		{Name: ("点位8"), Type: trailer.ValueType_NUMBER, Description: ("点位8参数")},
		{Name: ("点位9"), Type: trailer.ValueType_NUMBER, Description: ("点位9参数")},
		{Name: ("点位10"), Type: trailer.ValueType_NUMBER, Description: ("点位10参数")},
		{Name: ("点位11"), Type: trailer.ValueType_NUMBER, Description: ("点位11参数")},
		{Name: ("点位12"), Type: trailer.ValueType_NUMBER, Description: ("点位12参数")},
		{Name: ("点位13"), Type: trailer.ValueType_NUMBER, Description: ("点位13参数")},
		{Name: ("点位14"), Type: trailer.ValueType_NUMBER, Description: ("点位14参数")},
		{Name: ("点位15"), Type: trailer.ValueType_NUMBER, Description: ("点位15参数")},
		{Name: ("点位16"), Type: trailer.ValueType_NUMBER, Description: ("点位16参数")},
		{Name: ("点位17"), Type: trailer.ValueType_NUMBER, Description: ("点位17参数")},
		{Name: ("点位18"), Type: trailer.ValueType_NUMBER, Description: ("点位18参数")},
		{Name: ("点位19"), Type: trailer.ValueType_NUMBER, Description: ("点位19参数")},
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
		{Name: ("temp"), Type: trailer.ValueType_NUMBER, Value: []byte("15.34")},
		{Name: ("humi"), Type: trailer.ValueType_NUMBER, Value: []byte("65")},
		{Name: ("co2"), Type: trailer.ValueType_NUMBER, Value: []byte("13.5")},
		{Name: ("weather"), Type: trailer.ValueType_STRING, Value: []byte("SUNNY")},
		{Name: ("isOk"), Type: trailer.ValueType_BOOL, Value: []byte("false")},
		{Name: ("点位1"), Type: trailer.ValueType_NUMBER, Value: []byte("23.15")},
		{Name: ("点位2"), Type: trailer.ValueType_NUMBER, Value: []byte("23.25")},
		{Name: ("点位3"), Type: trailer.ValueType_NUMBER, Value: []byte("23.35")},
		{Name: ("点位4"), Type: trailer.ValueType_NUMBER, Value: []byte("23.45")},
		{Name: ("点位5"), Type: trailer.ValueType_NUMBER, Value: []byte("23.55")},
		{Name: ("点位6"), Type: trailer.ValueType_NUMBER, Value: []byte("23.65")},
		{Name: ("点位7"), Type: trailer.ValueType_NUMBER, Value: []byte("23.75")},
		{Name: ("点位8"), Type: trailer.ValueType_NUMBER, Value: []byte("23.85")},
		{Name: ("点位9"), Type: trailer.ValueType_NUMBER, Value: []byte("23.95")},
		{Name: ("点位10"), Type: trailer.ValueType_NUMBER, Value: []byte("65530")},
		{Name: ("点位11"), Type: trailer.ValueType_NUMBER, Value: []byte("65531")},
		{Name: ("点位12"), Type: trailer.ValueType_NUMBER, Value: []byte("65532")},
		{Name: ("点位13"), Type: trailer.ValueType_NUMBER, Value: []byte("65533")},
		{Name: ("点位14"), Type: trailer.ValueType_NUMBER, Value: []byte("65534")},
		{Name: ("点位15"), Type: trailer.ValueType_NUMBER, Value: []byte("65535")},
		{Name: ("点位16"), Type: trailer.ValueType_NUMBER, Value: []byte("65536")},
		{Name: ("点位17"), Type: trailer.ValueType_NUMBER, Value: []byte("65537")},
		{Name: ("点位18"), Type: trailer.ValueType_NUMBER, Value: []byte("65538")},
		{Name: ("点位19"), Type: trailer.ValueType_NUMBER, Value: []byte("65539")},
	}

	Rows := []*trailer.DataRow{}
	for i := 0; i < 123; i++ {
		Rows = append(Rows, &trailer.DataRow{
			Column: Values1,
		})
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
	log.Println("来自协议包的日志, main 参数:", os.Args)
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
