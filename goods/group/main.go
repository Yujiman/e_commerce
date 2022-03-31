package main

import (
	"github.com/Yujiman/e_commerce/goods/group/internal/bootstrap"
	"github.com/Yujiman/e_commerce/goods/group/internal/server"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.PingDbConnect()
}

func main() {
	server.InitServer()
}
