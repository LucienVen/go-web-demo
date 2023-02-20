package bootstrap

import (
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type Application struct {
	Env       *Env
	Mysql     *sqlx.DB
	Rpc       *grpc.Server
	RpcClient RpcClient
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mysql = NewMysqlDatabase(app.Env)
	app.Rpc = NewRpcServer(app.Env)
	app.RpcClient = InitRpcClient(app.Env)
	return *app
}

func (app *Application) CloseApplication() {
	CloseMysqlDbConnection(app.Mysql)
}

//func (app *Application) InitRpcClient(closeChan chan struct{}) {
//	app.RpcClient = NewRpcClient(app.Env, closeChan)
//}
