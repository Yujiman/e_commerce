package main

import (
	"github.com/Yujiman/e_commerce/auth/jwt/internal/bootstrap"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/handler"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/server"

	"log"
	"sync"
	"time"
)

func init() {
	bootstrap.Init()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.InitServer()
	}()

	for {
		err := handler.RevokeOldTokens()
		if err != nil {
			log.Println(err)
			break
		}
		time.Sleep(3 * time.Hour)
	}

	wg.Wait()
}
