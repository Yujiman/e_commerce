package main

import (
	"github.com/Yujiman/e_commerce/auth/domain/bootstrap"
	"github.com/Yujiman/e_commerce/auth/domain/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
