package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func commands() {
	app.Commands = []cli.Command{
		{
			Name:        "list",
			Usage:       "task list",
			Description: "List all incomplete tasks",
			Action: func(c *cli.Context) {
				fmt.Println("Your listed tasks: ")
			},
		},
		{
			Name:        "add",
			Usage:       "task add walk the dog",
			Description: "Add a single task: whatever comes after 'add'",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return errors.New("missing task to add")
				}
				var task strings.Builder
				for _, el := range c.Args() {
					task.WriteString(el)
					task.WriteString(" ")
				}
				fmt.Printf("We should add %v\n", task.String())
				return nil
			},
		},
		{
			Name:        "do",
			Usage:       "task do 3",
			Description: "Complete the task with the given id",
			Action: func(c *cli.Context) error {
				if c.NArg() != 1 {
					return errors.New("cmd 'do' completes exacly 1 task, by id number")
				}
				id, err := strconv.Atoi(c.Args()[0])
				if err != nil {
					return fmt.Errorf("error converting '%s' into id or integer", c.Args()[0])
				}
				fmt.Printf("We should complete the task with id %d\n", id)
				return nil
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
