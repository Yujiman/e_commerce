package console

import (
	"fmt"
	"log"

	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/handler/add"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/utils"

	"github.com/urfave/cli/v2"
)

var AddCommand = &cli.Command{
	Name:  "add",
	Usage: "add domain",
	Action: func(ctx *cli.Context) error {
		var (
			name string
			url  string
		)

		fmt.Print("Enter domain's name: ")
		utils.Scan(&name)

		fmt.Print("Enter domain's url: ")
		utils.Scan(&url)

		resp, err := add.Handle(&add.RequestDTO{
			Name: name,
			Url:  url,
		})
		if err != nil {
			return err
		}

		log.Println("Domain has been added with id=" + resp.DomainId)
		return nil
	},
}
