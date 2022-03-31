package main

import (
	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/bootstrap"
	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/server"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitMetrics()
}

func main() {
	server.InitServer()
}
