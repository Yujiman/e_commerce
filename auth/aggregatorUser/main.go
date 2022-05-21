package main

import (
	"github.com/Yujiman/e_commerce/auth/jwt/aggregatorUser/internal/bootstrap"
	"github.com/Yujiman/e_commerce/auth/jwt/aggregatorUser/internal/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
