package main

import (
	"github.com/Yujiman/e_commerce/basket/gatway/internal/bootstrap"
	"github.com/Yujiman/e_commerce/basket/gatway/internal/server/http"
)

func init() {
	bootstrap.InitConfig("./.env")
}

func main() {
	http.InitServer()
}
