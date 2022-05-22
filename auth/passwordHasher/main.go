package main

import (
	"github.com/Yujiman/e_commerce/auth/passwordHasher/bootstrap"
	"github.com/Yujiman/e_commerce/auth/passwordHasher/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
