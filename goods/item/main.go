package main

import (
	"github.com/Yujiman/e_commerce/goods/item/internal/bootstrap"
	"github.com/Yujiman/e_commerce/goods/item/server"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitMetrics()
	bootstrap.PingDbConnect()
	bootstrap.Migrate("./storage/db/migration/")
}

func main() {
	server.InitServer()
}
