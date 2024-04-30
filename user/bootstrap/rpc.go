package bootstrap
//
//import (
//	"github.com/LucienVen/go-web-demo/user/internal/rpc"
//	"github.com/LucienVen/go-web-demo/utils/rpc/pb"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/reflection"
//)
//
//func NewRpcServer(env *Env) *grpc.Server {
//	s := grpc.NewServer()
//	// TODO 使用 serverName 进行注册
//	pb.RegisterHelloServer(s, rpc.NewServer1Rpc())
//	reflection.Register(s)
//	return s
//}
