package main

import "github.com/LucienVen/go-web-demo/base/bootstrap"

func main()  {
	instance := bootstrap.App()
	instance.PrintConfig()

}