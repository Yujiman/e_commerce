package bootstrap

import (
	"fmt"
	"log"

	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db"
	"github.com/Yujiman/e_commerce/auth/user/internal/utils"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("Loading dotenv file failed " + err.Error())
	}
	log.Println("Init godotenv")

	err = migrate()
	if err != nil {
		log.Panicf("Migration failed: " + err.Error())
	}
	log.Println("Init migrate finished")
}

func migrate() error {
	log.Println("Check DB started")
	dbConn := db.GetDbConnection()
	log.Println("Check DB finished")

	dir := "./storage/db/migration/"

	files, err := utils.GetFiles(dir)
	if err != nil {
		return err
	}

	for _, fileName := range files {
		queryString := utils.ReadFile(dir + fileName)
		fmt.Println(dir + fileName)
		_, err = dbConn.Exec(queryString)
		if err != nil {
			return fmt.Errorf("Migration=%s, Query error=%v\n", fileName, err)
		}
	}

	return nil
}
