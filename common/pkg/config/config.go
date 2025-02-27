package config

import (
	"log"

	"github.com/spf13/viper"
)

// Env interface 定义
type Env interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	IsDevelopment() bool
}

// NewEnv 函数
func NewEnv(env Env) Env {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.IsDevelopment() {
		log.Println("The App is running in development env")
	}

	return env
}
