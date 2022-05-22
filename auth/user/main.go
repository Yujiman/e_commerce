package main

import (
	"github.com/Yujiman/e_commerce/auth/user/bootstrap"
	"github.com/Yujiman/e_commerce/auth/user/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
