package main

import (
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/bootstrap"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
