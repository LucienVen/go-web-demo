package bootstrap

import (
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type Application struct {
	Env   *Env
	Mysql *sqlx.DB
	Rpc   *grpc.Server
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mysql = NewMysqlDatabase(app.Env)
	app.Rpc = NewRpcServer(app.Env)
	return *app
}

func (app *Application) CloseApplication() {
	CloseMysqlDbConnection(app.Mysql)
}
