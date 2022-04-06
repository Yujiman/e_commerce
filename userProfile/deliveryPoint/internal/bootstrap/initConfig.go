package bootstrap

import (
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/config"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/utils"

	"github.com/joho/godotenv"
)

func InitConfig() {
	utils.LogPrintln(utils.Yellow("Loading .env file..."))
	err := godotenv.Load()
	if err != nil {
		utils.LogPanicf(utils.Fata("Loading .env file failed "))
	}

	config.GetConfig()

	utils.LogPrintln(utils.Green("Initialized .env file"))
}
