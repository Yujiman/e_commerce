package main

import (
	"github.com/Yujiman/e_commerce/auth/authentication/bootstrap"
	"github.com/Yujiman/e_commerce/auth/authentication/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
