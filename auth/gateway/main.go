package main

import (
	"github.com/Yujiman/e_commerce/auth/gateway/bootstrap"
	"github.com/Yujiman/e_commerce/auth/gateway/server/http"
)

func init() {
	bootstrap.InitEnv("./.env")
}

func main() {
	http.InitServer()
}
