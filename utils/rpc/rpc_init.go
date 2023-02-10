package rpc

import (
	"github.com/LucienVen/go-web-demo/utils/pb"
	"github.com/LucienVen/go-web-demo/utils/rpc/server1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

// 初始化rpc server
func RunGrpcServer(port string) error {
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, server1.NewServer1Rpc())
	reflection.Register(s)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	
	return s.Serve(lis)
}
