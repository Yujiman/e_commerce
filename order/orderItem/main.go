package main

import (
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/bootstrap"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/server"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.PingDbConnect()
	bootstrap.Migrate("./storage/db/migration/")
}

func main() {
	server.InitServer()
}
