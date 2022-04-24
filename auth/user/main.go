package main

import (
	"github.com/Yujiman/e_commerce/auth/user/internal/bootstrap"
	"github.com/Yujiman/e_commerce/auth/user/internal/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
