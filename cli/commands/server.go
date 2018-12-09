package commands

import (
	"github.com/minhajuddinhkhan/ttd/httpsvr"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var addr string

//Serve serves the http server
var Serve = &cli.Command{
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "addr, a",
			Usage:       "start server on address",
			Destination: &addr,
		},
	},
	Before: func(c *cli.Context) error {
		if addr == "" {
			cli.NewExitError("invalid address", 1)
		}
		logrus.Info("starting server on addr")
		return nil
	},
	Action: func(c *cli.Context) error {
		return httpsvr.ServeHTTP(addr)
	},
}
