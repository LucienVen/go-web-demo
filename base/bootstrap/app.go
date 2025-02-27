package bootstrap

import (
	"fmt"

	"github.com/LucienVen/go-web-demo/base/config"
)

type Application struct{
	Env *config.Env
}

func App() Application {
	config.ConfigInit()
	return Application{
		Env: config.Config,
	}

}


func (app *Application) PrintConfig() {
	fmt.Printf("App Environment: %s\n", app.Env.AppEnv)
	fmt.Printf("App Name: %s\n", app.Env.AppName)
	fmt.Printf("Server Address: %s\n", app.Env.ServerAddress)
	fmt.Printf("Context Timeout: %d\n", app.Env.ContextTimeout)
	fmt.Printf("Database Host: %s\n", app.Env.DBHost)
	fmt.Printf("Database Port: %s\n", app.Env.DBPort)
	fmt.Printf("Database User: %s\n", app.Env.DBUser)
	fmt.Printf("Database Password: %s\n", app.Env.DBPass)
	fmt.Printf("Database Name: %s\n", app.Env.DBName)
	fmt.Printf("Access Token Expiry Hour: %d\n", app.Env.AccessTokenExpiryHour)
	fmt.Printf("Refresh Token Expiry Hour: %d\n", app.Env.RefreshTokenExpiryHour)
	fmt.Printf("Access Token Secret: %s\n", app.Env.AccessTokenSecret)
	fmt.Printf("Refresh Token Secret: %s\n", app.Env.RefreshTokenSecret)
	fmt.Printf("User Service Port: %s\n", app.Env.UserServicePort)
	fmt.Printf("Order Service Port: %s\n", app.Env.OrderServicePort)
	fmt.Printf("RPC Client: %s\n", app.Env.RpcClient)
}