package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "ttd"
	app.Usage = "server using ttd"
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
