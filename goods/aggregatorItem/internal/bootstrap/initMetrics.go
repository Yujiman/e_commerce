package bootstrap

import (
	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/config"
	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/utils"
	"strconv"

	metricService "github.com/autokz/go-monitor"
)

func InitMetrics() {
	metricConf := config.GetConfig().MetricsParams
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
