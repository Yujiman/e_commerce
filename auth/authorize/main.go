package main

import (
	"github.com/Yujiman/e_commerce/auth/authorize/internal/bootstrap"
	"github.com/Yujiman/e_commerce/auth/authorize/internal/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
