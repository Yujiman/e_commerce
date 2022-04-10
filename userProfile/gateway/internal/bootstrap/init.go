package bootstrap

import (
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/config"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/utils"

	"github.com/joho/godotenv"
)

func InitConfig(envDir string) {
	utils.LogPrintln(utils.Yellow("Loading .env file..."))
	err := godotenv.Load(envDir)
	if err != nil {
		utils.LogPanicf(utils.Fata("Loading .env file failed "))
	}

	config.GetConfig()

	utils.LogPrintln(utils.Green("Initialized .env file"))
}
