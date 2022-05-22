package main

import (
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/bootstrap"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
