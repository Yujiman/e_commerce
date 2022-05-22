package bootstrap

import (
	"strconv"

	"github.com/Yujiman/e_commerce/auth/gateway/config"
	"github.com/Yujiman/e_commerce/auth/gateway/utils"

	metricService "github.com/autokz/go-monitor"
	"github.com/joho/godotenv"
)

func InitEnv(envDir string) {
	err := godotenv.Load(envDir)
	if err != nil {
		utils.LogPanicf(utils.Fata("Loading .env file failed "))
	}
	utils.LogPrintln(utils.Yellow("Initialized .env file"))
}

func InitMetrics() {
	metricConf := config.GetMetricConf()
	metricTimeout, err := strconv.Atoi(metricConf.MetricTimeout)
	if err != nil {
		utils.LogPanicf(utils.Fata("METRIC_TIMEOUT environment not valid"))
	}
	metricService.Handle(
		metricConf.MetricServerAddress,
		metricConf.MetricServerPort,
		"10s",
		metricConf.MetricAppName,
		metricTimeout,
	)
	utils.LogPrintln(utils.Yellow("Initialized metricService"))
}
