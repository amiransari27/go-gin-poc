package main

import (
	"fmt"
	"go-gin-api/src/ioc"
	"go-gin-api/src/server"
)

func main() {
	fmt.Println("go gin api")

	kernal := ioc.NewKernal()
	apiServer := server.NewGinServer(kernal, ioc.LoadControllers)
	apiServer.Start()

}
