package commands

import (
	"errors"
	"fmt"
	"time"

	"github.com/boltdb/bolt"

	"github.com/urfave/cli"
)

func listTasks(c *cli.Context) error {
	fmt.Println("Your listed tasks: ")

	db, err := bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	defer db.Close()

	// TODO: Initialize bucket in main or somewhere
	// Perhaps use Store, aka struct that houses db for command package

	var tasksToList []string

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("todo"))
		if bucket == nil {
			return errors.New("bucket does not exist")
		}
		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			task := fmt.Sprintf("%s) %s\n", string(k), v)
			tasksToList = append(tasksToList, task)
		}
		return nil
	})

	if err != nil {
		return err
	}

	for _, task := range tasksToList {
		fmt.Println(task)
	}

	return nil
}
