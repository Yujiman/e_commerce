package main

import (
	"log"
	"os"

	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/bootstrap"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/console"

	"github.com/urfave/cli/v2"
)

func init() {
	bootstrap.Init()
}

func main() {
	commands := cli.Commands{
		console.AddCommand,
		console.UpdateCommand,
		console.RemoveCommand,
	}
	app := &cli.App{Name: "domain-manager", Commands: commands}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
