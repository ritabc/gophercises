package commands

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/urfave/cli"
)

// GetAll aggregates and returns all commands in this package
func GetAll() []cli.Command {
	return []cli.Command{
		{
			Name:        "list",
			Usage:       "task list",
			Description: "List all incomplete tasks",
			Action:      listTasks,
		},
		{
			Name:        "add",
			Usage:       "task add walk the dog",
			Description: "Add a single task: whatever comes after 'add'",
			Action:      addTask,
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
