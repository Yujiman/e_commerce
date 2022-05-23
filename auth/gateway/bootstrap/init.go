package bootstrap

import (
	"github.com/Yujiman/e_commerce/auth/gateway/utils"

	"github.com/joho/godotenv"
)

func InitEnv(envDir string) {
	err := godotenv.Load(envDir)
	if err != nil {
		utils.LogPanicf(utils.Fata("Loading .env file failed "))
	}
	utils.LogPrintln(utils.Yellow("Initialized .env file"))
}
