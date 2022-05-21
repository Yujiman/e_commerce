package main

import (
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/internal/bootstrap"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/internal/server"
)

func init() {
	bootstrap.Init()
}

func main() {
	server.InitServer()
}
