package main

import (
	"github.com/Yujiman/e_commerce/auth/authorize/bootstrap"
	"github.com/Yujiman/e_commerce/auth/authorize/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
