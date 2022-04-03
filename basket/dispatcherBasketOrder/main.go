package main

import (
	"github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/bootstrap"
	"github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/server"
)

func init() {
	bootstrap.InitConfig()
}

func main() {
	server.InitServer()
}
