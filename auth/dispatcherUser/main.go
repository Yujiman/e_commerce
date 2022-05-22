package main

import (
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/bootstrap"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
