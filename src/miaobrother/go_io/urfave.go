package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "green"

	app.Usage = "fight the loneliness"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}
	app.Run(os.Args)
}
