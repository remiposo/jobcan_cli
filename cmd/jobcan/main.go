package main

import (
	"fmt"
	"log"
	"os"

	"github.com/remiposo/jobcan_cli/internal/jobcan"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "email",
				Usage:    "email for jobcan",
				EnvVars:  []string{"JOBCAN_EMAIL"},
				Required: true,
			},
			&cli.StringFlag{
				Name:        "password",
				Usage:       "password for jobcan",
				EnvVars:     []string{"JOBCAN_PASSWORD"},
				Required:    true,
				DefaultText: "****",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "attend",
				Usage: "Attend work",
				Action: func(c *cli.Context) error {
					email := c.String("email")
					password := c.String("password")
					err := jobcan.AttendWork(email, password)
					if err != nil {
						return err
					}
					fmt.Println("attended!!")
					return nil
				},
			},
			{
				Name:  "leave",
				Usage: "Leave work",
				Action: func(c *cli.Context) error {
					email := c.String("email")
					password := c.String("password")
					err := jobcan.LeaveWork(email, password)
					if err != nil {
						return err
					}
					fmt.Println("leaved!!")
					return nil
				},
			},
			{
				Name:  "status",
				Usage: "Show working status",
				Action: func(c *cli.Context) error {
					email := c.String("email")
					password := c.String("password")
					status, err := jobcan.GetStatus(email, password)
					if err != nil {
						return err
					}
					fmt.Println(status)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
