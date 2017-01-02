package command

import (
	"fmt"
	"log"
	"os"

	"github.com/comail/colog"
	"github.com/k0kubun/pp"
	"github.com/nlopes/slack"
	"github.com/urfave/cli"
)

type (
	SlackDndCtl struct {
		Token string
		Api   *slack.Client
	}
)

func setLoggerColog() {
	colog.SetDefaultLevel(colog.LInfo)
	colog.SetMinLevel(colog.LInfo)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime,
	})
	colog.Register()
	/*
		log.Printf("trace: this is a trace log.")
		log.Printf("debug: this is a debug log.")
		log.Printf("info: this is an info log.")
		log.Printf("warn: this is a warning log.")
		log.Printf("error: this is an error log.")
		log.Printf("alert: this is an alert log.")
		log.Printf("this is a default level log.")
	*/
}

func newSlackDndCtl(c *cli.Context) (sdc *SlackDndCtl) {
	if emptyString(c.GlobalString("token")) {
		fmt.Fprintf(os.Stderr, "Arg: 'token' is must be required. See '--help'.")
		os.Exit(2)
	}
	sdc = &SlackDndCtl{
		Token: c.GlobalString("token"),
	}
	sdc.auth()
	return
}

func emptyString(s string) (b bool) {
	switch s {
	case "":
		b = true
		return
	default:
		b = false
		return
	}
}

func (sdc *SlackDndCtl) auth() {
	sdc.Api = slack.New(sdc.Token)
}

func (sdc *SlackDndCtl) Test() {
	r, err := sdc.Api.AuthTest()
	pp.Print(err)
	pp.Print(r)
}

func (sdc *SlackDndCtl) EndDND() (err error) {
	err = sdc.Api.EndDND()
	return
}
