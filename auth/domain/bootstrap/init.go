package bootstrap

import (
	"log"

	"github.com/Yujiman/e_commerce/auth/domain/storage/db"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("Loading dotenv file failed " + err.Error())
	}
	log.Println("Init godotenv")

	log.Println("Check DB started")
	db.GetDbConnection()
	log.Println("Check DB finished")
}
