package main

import (
	"log"
	"sync"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/bootstrap"
	"github.com/Yujiman/e_commerce/auth/jwt/handler"
	"github.com/Yujiman/e_commerce/auth/jwt/server"
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
		log.Println("Check for expired tokens.")
		err := handler.RevokeOldTokens()
		if err != nil {
			log.Println(err)
			break
		}
		time.Sleep(3 * time.Hour)
	}

	wg.Wait()
}
