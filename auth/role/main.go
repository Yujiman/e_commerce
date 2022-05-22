package main

import (
	"github.com/Yujiman/e_commerce/auth/role/bootstrap"
	"github.com/Yujiman/e_commerce/auth/role/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
