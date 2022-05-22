package main

import (
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/bootstrap"
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
