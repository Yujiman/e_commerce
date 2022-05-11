package bootstrap

import (
	"log"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln("Loading dotenv file failed " + err.Error())
	}
	log.Println("Loading dotenv finished.")
}
