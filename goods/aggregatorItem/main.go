package main

import (
	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/bootstrap"
	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/server"
)

func init() {
	bootstrap.InitConfig()
}

func main() {
	server.InitServer()
}
