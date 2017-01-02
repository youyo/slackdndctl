package command

import (
	"log"

	"github.com/urfave/cli"
)

func CmdEndDND(c *cli.Context) (err error) {
	setLoggerColog()
	sdc := newSlackDndCtl(c)
	err = sdc.EndDND()
	log.Printf("info: Ends the current user's Do Not Disturb session.")
	return
}
