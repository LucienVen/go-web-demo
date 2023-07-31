package rpc

import "github.com/LucienVen/go-web-demo/utils/rpc/pb"

//type Server1Rpc struct {
//	pb.UnimplementedHelloServer
//}

//
//func NewServer1Rpc() *Server1Rpc {
//	return &Server1Rpc{}
//}
//
////func (r *Server1Rpc) mustEmbedUnimplementedHelloServer() {
////	//TODO implement me
////	panic("implement me")
////}
//
//func (r *Server1Rpc) SayHello(ctx context.Context, req *pb.SayHelloReq) (*pb.SayHelloRes, error) {
//	greet := fmt.Sprintf("Hello, %s!", req.ServerName)
//	return &pb.SayHelloRes{Greet: greet}, nil
//}

type UserRpc struct {
	pb.UnimplementedHelloServer
}

func NewUserRpc() *UserRpc {
	return &UserRpc{}
}
