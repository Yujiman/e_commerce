package main

import (
	"github.com/Yujiman/e_commerce/auth/aggregatorDomain/bootstrap"
	"github.com/Yujiman/e_commerce/auth/aggregatorDomain/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
