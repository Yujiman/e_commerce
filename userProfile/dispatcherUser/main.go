package main

import (
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/bootstrap"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/server"
)

func init() {
	bootstrap.InitConfig()
}

func main() {
	server.InitServer()
}
