package bootstrap

import (
	"log"
	"strconv"

	"github.com/Yujiman/e_commerce/auth/domain/config"
	"github.com/Yujiman/e_commerce/auth/domain/storage/db"

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

	log.Println("Check DB started")
	db.GetDbConnection()
	log.Println("Check DB finished")
}
