package main

import (
	"github.com/Yujiman/e_commerce/auth/role/internal/bootstrap"

	"github.com/Yujiman/e_commerce/auth/role/internal/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
