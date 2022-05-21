package main

import (
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/bootstrap"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
