package bootstrap

import (
	"github.com/Yujiman/e_commerce/userProfile/city/internal/storage/db"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/utils"
)

func Migrate(dir string) {
	utils.LogPrintln(utils.Yellow("Start migration..."))

	dbConn := db.GetDbConnection()

	files, err := utils.GetFiles(dir)
	if err != nil {
		utils.LogPrintln(utils.Fata("Migration failed!"))
		utils.LogFatalf("Get file error: %v", err)
	}

	for key, fileName := range files {
		queryString := utils.ReadFile(dir + fileName)

		utils.LogPrintf("(%d/%d) %v%v", key+1, len(files), dir, fileName)

		_, err = dbConn.Exec(queryString)
		if err != nil {
			utils.LogPrintln(utils.Fata("Migration failed!"))
			utils.LogFatalf("Migration=%s, Query error=%v\n", fileName, err)
		}
	}

	utils.LogPrintln(utils.Green("Migrations done!"))
}

func PingDbConnect() {
	utils.LogPrintln(utils.Yellow("Starting checking DB..."))
	db.GetDbConnection()
	utils.LogPrintln(utils.Green("Checking DB successfully finished!"))
}
