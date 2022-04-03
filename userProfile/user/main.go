package main

import (
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/bootstrap"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/server"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.PingDbConnect()
	bootstrap.Migrate("./internal/storage/db/migration/")
}

func main() {
	server.InitServer()
}
