package main

import (
	"github.com/Yujiman/e_commerce/auth/jwt/authentication/internal/bootstrap"
	"github.com/Yujiman/e_commerce/auth/jwt/authentication/internal/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
