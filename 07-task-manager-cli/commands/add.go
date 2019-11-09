package commands

import (
	"fmt"
	"gophercises/07-task-manager-cli/db"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func addTask(c *cli.Context) {
	if c.NArg() == 0 {
		fmt.Println("missing task to add")
		os.Exit(1)
	}
	s := strings.Join(c.Args(), " ")
	task, err := db.AddTask(s)
	if err != nil {
		fmt.Println("Something went wrong:", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Added task: '%s' with ID of %d\n", task.Value, task.Key)
}
