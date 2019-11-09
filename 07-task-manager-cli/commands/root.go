package commands

import (
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
			Action:      doTask,
		},
		{
			Name:        "done",
			Usage:       "task done",
			Description: "List done tasks",
			Action:      listDone,
		},
	}
}
