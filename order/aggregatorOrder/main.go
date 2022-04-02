package main

import (
	"github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/bootstrap"
	"github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/server"
)

func init() {
	bootstrap.InitConfig()
}

func main() {
	server.InitServer()
}
