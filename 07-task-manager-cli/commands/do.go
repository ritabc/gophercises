package commands

import (
	"fmt"
	"gophercises/07-task-manager-cli/db"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

func doTask(c *cli.Context) {
	if c.NArg() < 1 {
		fmt.Println("cmd 'do' expects at least 1 task id to complete")
		os.Exit(1)
	}
	for _, idString := range c.Args() {
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			fmt.Println("error parsing %s as ID", idString)
			continue
		}
		completed, err := db.MarkDone(idInt)
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Good job! Task: '%s' with ID of %d is now marked as completed!\n", completed.Value, completed.Key)
	}

}
