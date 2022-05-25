package main

import (
	"github.com/Yujiman/e_commerce/goods/gatway/internal/bootstrap"
	"github.com/Yujiman/e_commerce/goods/gatway/internal/server/http"
)

func init() {
	bootstrap.InitConfig("./.env")
}

func main() {
	http.InitServer()
}
