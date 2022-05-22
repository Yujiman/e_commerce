package config

import "os"

type LoggerConf struct {
	ServerAddress string
	ServerPort    string
	AppName       string
	Timeout       string
}

func GetLoggerConf() *LoggerConf {
	return &LoggerConf{
		ServerAddress: os.Getenv("LOGGER_SERVER_ADDRESS"),
		ServerPort:    os.Getenv("LOGGER_SERVER_PORT"),
		AppName:       os.Getenv("LOGGER_APP_NAME"),
		Timeout:       os.Getenv("LOGGER_TIMEOUT"),
	}
}
