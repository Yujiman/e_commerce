package main

import (
	"github.com/Yujiman/e_commerce/goods/item/internal/bootstrap"
	"github.com/Yujiman/e_commerce/goods/item/internal/server"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.PingDbConnect()
	bootstrap.Migrate("./internal/storage/db/migration/")
}

func main() {
	server.InitServer()
}
