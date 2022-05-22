package console

import (
	"fmt"
	"log"

	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/handler/update"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/utils"

	"github.com/urfave/cli/v2"
)

var UpdateCommand = &cli.Command{
	Name:  "update",
	Usage: "update domain",
	Action: func(ctx *cli.Context) error {
		var (
			domainId string
			name     string
			url      string
		)

		fmt.Print("Enter domain's id: ")
		utils.Scan(&domainId)

		fmt.Print("Enter domain's name: ")
		utils.Scan(&name)

		fmt.Print("Enter domain's url: ")
		utils.Scan(&url)

		err := update.Handle(&update.RequestDTO{
			DomainId: domainId,
			Name:     name,
			Url:      url,
		})
		if err != nil {
			return err
		}

		log.Println("Domain has been updated.")
		return nil
	},
}
