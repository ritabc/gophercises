package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func commands() {
	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "List all incomplete tasks",
			Action: func(c *cli.Context) {
				fmt.Println("Your listed tasks: ")
			},
		},
		{
			Name:  "add",
			Usage: "Add the given task",
			Action: func(c *cli.Context) {
				if c.Args()[0] != "" {
					fmt.Printf("We should add %v\n", c.Args()[0])
				}
			},
		},
		{
			Name:  "do",
			Usage: "Complete the task with the given id",
			Action: func(c *cli.Context) {
				if c.Args()[0] != "" {
					fmt.Printf("We should complete the task with id %v\n", c.Args()[0])
				}
			},
		},
	}
}

func info() {
	app.Name = "task"
	app.Usage = "A task manager built using urfave/cli and boltdb"
	app.Author = "Rita"
	app.Version = "1.0.0"
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

/* Note that after building, if we run:
$ ./task add task1
Then these two expressions both produce 'task1'
c.Args()[0]
os.Args[2]
*/
