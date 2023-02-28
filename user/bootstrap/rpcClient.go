package bootstrap

import (
	"fmt"
	"github.com/LucienVen/go-web-demo/utils/rpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// TODO 当前grpc client 无法动态创建
//type RpcClient struct {
//	Name      string
//	Conn      *grpc.ClientConn
//	CloseChan <-chan struct{} // TODO 监听chan，关闭连接
//}
//
//// rpc client
//func NewRpcClient(env *Env, closeChan chan struct{}) map[string]*RpcClient {
//	rpcClientMap := make(map[string]*RpcClient)
//	clientSettings := env.RpcClient
//	clients := strings.Split(clientSettings, ",")
//	for _, item := range clients {
//		rpcSetting := strings.Split(item, ",")
//		rpcClientName := rpcSetting[0]
//		//rpcClientPort := rpcSetting[1]
//		if rpcClientName == env.AppName {
//			continue
//		}
//
//		conn, err := grpc.Dial(item, grpc.WithTransportCredentials(insecure.NewCredentials()))
//		if err != nil {
//			log.Fatal("grpc did not connect: %v, rpc client name:%s", err, rpcClientName)
//		}
//
//		rpcClientMap[rpcClientName] = &RpcClient{
//			Name:      rpcClientName,
//			Conn:      conn,
//			CloseChan: closeChan,
//		}
//	}
//
//	return rpcClientMap
//}

// *********** 手动创建 grpc client ***********

var (
	Server2RpcClient pb.HelloClient // Order
)

type RpcClient struct {
	Order pb.HelloClient
}

func InitRpcClient(env *Env) RpcClient {
	addr := fmt.Sprintf("%s:%s", "order", env.OrderServicePort)

	opts := make([]grpc.DialOption, 0)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatal("grpc did not connect: %v, rpc client name:%s", err)
	}

	Server2RpcClient = pb.NewHelloClient(conn)

	return RpcClient{
		Order: Server2RpcClient,
	}

}
