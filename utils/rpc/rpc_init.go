package rpc

import (
	"github.com/LucienVen/go-web-demo/utils/rpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

// 初始化rpc server
func RunGrpcServer(serverName, port string) error {
	s := grpc.NewServer()
	// TODO 使用 serverName 进行注册
	pb.RegisterHelloServer(s, NewServer1Rpc())
	reflection.Register(s)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	return s.Serve(lis)
}
