package console

import (
	"fmt"
	"log"

	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/handler/remove"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/utils"

	"github.com/urfave/cli/v2"
)

var RemoveCommand = &cli.Command{
	Name:  "remove",
	Usage: "remove domain",
	Action: func(ctx *cli.Context) error {
		var domainId string

		fmt.Print("Enter domain_id: ")
		utils.Scan(&domainId)

		err := remove.Handle(&remove.RemoveDTO{DomainId: domainId})
		if err != nil {
			return err
		}

		log.Println("Domain has been removed")
		return nil
	},
}
