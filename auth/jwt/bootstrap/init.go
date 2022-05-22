package bootstrap

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Yujiman/e_commerce/auth/jwt/config"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db"
	"github.com/Yujiman/e_commerce/auth/jwt/utils"

	metricService "github.com/autokz/go-monitor"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("Loading dotenv file failed " + err.Error())
	}
	log.Println("Init godotenv")

	metricConf := config.GetMetricConf()
	metricTimeout, err := strconv.Atoi(metricConf.MetricTimeout)
	if err != nil {
		log.Panicf("METRIC_TIMEOUT environment not valid")
	}
	metricService.Handle(
		metricConf.MetricServerAddress,
		metricConf.MetricServerPort,
		"10s",
		metricConf.MetricAppName,
		metricTimeout,
	)
	log.Println("Init metricService")

	err = migrate()
	if err != nil {
		log.Panicf("Migration failed: " + err.Error())
	}
	log.Println("Init migrate finished")

	tokenClaimsInit()
	log.Println("Init tokens finished")
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

func tokenClaimsInit() {
	// call init functions for access token
	config.GetAccessTokenClaims()
	config.GetAccessTokenLifeTimeMinutes()

	// call init functions for refresh token
	config.GetRefreshTokenClaims()
	config.GetRefreshTokenLifeTimeMinutes()

	// call init function for init token keys
	config.GetKeysStorage()
}
