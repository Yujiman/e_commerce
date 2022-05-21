package bootstrap

import (
	"log"

	"github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/config"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("Loading dotenv file failed " + err.Error())
	}
	log.Println("Init godotenv")

	config.GetServicesParams()
	log.Println("Init services and their params")
}
