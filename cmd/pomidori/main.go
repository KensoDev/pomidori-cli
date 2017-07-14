package main

import (
	"fmt"
	"github.com/kensodev/pomidori"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.2"

	app.Commands = []cli.Command{
		{
			Name:  "register",
			Usage: "Register an account on the pomidori server",
			Action: func(c *cli.Context) error {
				client := pomidori.NewClient()
				token, err := client.Register()

				tokenBytes := []byte(token)
				err = ioutil.WriteFile(".token", tokenBytes, 0644)

				if err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name:  "work",
			Usage: "Start working on a task",
			Action: func(c *cli.Context) error {
				tokenBytes, err := ioutil.ReadFile(".token")

				if err != nil {
					return fmt.Errorf("Could not find a token. You need to register/login")
				}

				token := string(tokenBytes)

				client := pomidori.NewRegisteredClient(token)
				err = client.CreateTask("sample", "10")

				if err != nil {
					fmt.Println(err)
				}

				return nil

			},
		},
	}

	app.Run(os.Args)
}
