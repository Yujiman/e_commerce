package main

import (
	"github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/bootstrap"
	"github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/server"
)

func init() {
	bootstrap.InitConfig()
}

func main() {
	server.InitServer()
}
