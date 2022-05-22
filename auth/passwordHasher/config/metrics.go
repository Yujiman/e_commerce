package config

import "os"

type MetricsConf struct {
	MetricServerAddress string
	MetricServerPort    string
	MetricAppName       string
	MetricTimeout       string
}

func GetMetricConf() *MetricsConf {
	return &MetricsConf{
		MetricServerAddress: os.Getenv("METRIC_SERVER_ADDRESS"),
		MetricServerPort:    os.Getenv("METRIC_SERVER_PORT"),
		MetricAppName:       os.Getenv("METRIC_APP_NAME"),
		MetricTimeout:       os.Getenv("METRIC_TIMEOUT"),
	}
}
