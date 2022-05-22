package bootstrap

import (
	"log"
	"strconv"

	"github.com/Yujiman/e_commerce/auth/passwordHasher/config"

	metricService "github.com/autokz/go-monitor"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln("Loading dotenv file failed " + err.Error())
	}
	log.Println("Loading dotenv finished.")
	metricConf := config.GetMetricConf()
	metricTimeout, err := strconv.Atoi(metricConf.MetricTimeout)
	if err != nil {
		log.Panicln("Metric environment not valid" + err.Error())
	}
	metricService.Handle(
		metricConf.MetricServerAddress,
		metricConf.MetricServerPort,
		"10s",
		metricConf.MetricAppName,
		metricTimeout,
	)
	log.Println("Metric initialization finished.")
}
