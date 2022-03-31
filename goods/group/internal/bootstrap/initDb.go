package bootstrap

import (
	"github.com/Yujiman/e_commerce/goods/group/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/group/internal/utils"
)

func PingDbConnect() {
	utils.LogPrintln(utils.Yellow("Starting checking DB..."))
	db.GetDbConnection()
	utils.LogPrintln(utils.Green("Checking DB successfully finished!"))
}
