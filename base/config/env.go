package config

import (
	"errors"

	"github.com/LucienVen/go-web-demo/common/pkg/config"

	"reflect"
)

var Config *Env

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	AppName                string `mapstructure:"APP_NAME"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	UserServicePort        string `mapstructure:"USER_SERVICE_PORT"`  // user rpc 监听端口
	OrderServicePort       string `mapstructure:"ORDER_SERVICE_PORT"` // order rpc 监听端口
	RpcClient              string `mapstructure:"RPC_CLIENT"`         // 服务名:端口，逗号分割
}

func (e *Env) Get(key string) (interface{}, error) {
	val := reflect.ValueOf(e).Elem()
	field := val.FieldByName(key)
	if !field.IsValid() {
		return nil, errors.New("key not found")
	}
	return field.Interface(), nil
}

func (e *Env) Set(key string, value interface{}) error {
	val := reflect.ValueOf(e).Elem()
	field := val.FieldByName(key)
	if !field.IsValid() {
		return errors.New("key not found")
	}
	if !field.CanSet() {
		return errors.New("cannot set value")
	}
	field.Set(reflect.ValueOf(value))
	return nil
}

func (e *Env) IsDevelopment() bool {
	return e.AppEnv == "development"
}

func ConfigInit() {
	Config = config.NewEnv()
}