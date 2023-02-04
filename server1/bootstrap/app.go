package bootstrap

import "github.com/jmoiron/sqlx"

type Application struct {
	Env   *Env
	Mysql *sqlx.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mysql = NewMysqlDatabase(app.Env)
	return *app
}

func (app *Application) CloseApplication() {
	CloseMysqlDbConnection(app.Mysql)
}
