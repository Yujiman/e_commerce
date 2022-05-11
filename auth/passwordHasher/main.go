package main

import (
	"github.com/Yujiman/e_commerce/auth/passwordHasher/internal/bootstrap"
	"github.com/Yujiman/e_commerce/auth/passwordHasher/internal/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
