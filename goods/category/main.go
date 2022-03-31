package main

import (
	"github.com/Yujiman/e_commerce/goods/category/internal/bootstrap"
	"github.com/Yujiman/e_commerce/goods/category/internal/server"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.PingDbConnect()
}

func main() {
	server.InitServer()
}
