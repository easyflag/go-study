package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	opPriorityInit()

	app := cli.NewApp()
	app.Name = "calculater"
	app.Usage = "calculate expression"
	app.Version = "1.0"

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			cmd := c.Args().First()
			fmt.Println("cmd is", cmd)
		}
		return process()
	}

	app.Run(os.Args)
}
