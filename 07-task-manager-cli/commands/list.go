package commands

import (
	"fmt"
	"gophercises/07-task-manager-cli/db"
	"os"

	"github.com/urfave/cli"
)

func listTasks(c *cli.Context) {
	tasks, err := db.AllTasks()
	if err != nil {
		fmt.Println("Something went wrong:", err.Error())
		os.Exit(1)
	}

	fmt.Println("You have the following tasks:")
	for _, task := range tasks {
		fmt.Printf("%d) %s\n", task.Key, task.Value)
	}
}
