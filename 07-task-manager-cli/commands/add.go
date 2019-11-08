package commands

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/boltdb/bolt"

	"github.com/urfave/cli"
)

func addTask(c *cli.Context) error {
	if c.NArg() == 0 {
		return errors.New("missing task to add")
	}
	var taskToAdd bytes.Buffer
	for _, el := range c.Args() {
		taskToAdd.WriteString(el)
		taskToAdd.WriteString(" ")
	}

	db, err := bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return err
		}
		nextID, err := bucket.NextSequence()
		if err != nil {
			return err
		}
		taskID := []byte(strconv.Itoa(int(nextID)))

		err = bucket.Put(taskID, taskToAdd.Bytes())

		return nil
	})

	fmt.Printf("Added task: %v\n", taskToAdd.String())
	return nil
}
