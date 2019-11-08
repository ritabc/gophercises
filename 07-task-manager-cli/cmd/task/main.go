package main

import (
	"gophercises/07-task-manager-cli/commands"
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "task"
	app.Usage = "A task manager built using urfave/cli and boltdb"
	app.Author = "Rita"
	app.Version = "1.0.0"
}

func main() {
	info()
	app.Commands = commands.GetAll()

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
