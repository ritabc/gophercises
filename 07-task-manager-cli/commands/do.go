package commands

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/boltdb/bolt"

	"github.com/urfave/cli"
)

func doTask(c *cli.Context) error {
	if c.NArg() != 1 {
		return errors.New("cmd 'do' completes exacly 1 task, by id number")
	}
	stringID := c.Args()[0]

	db, err := bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	defer db.Close()

	var completedTask string
	// Delete from todo, get task text
	err = db.Update(func(tx *bolt.Tx) error {
		todoBucket, err := tx.CreateBucketIfNotExists([]byte("todo"))
		if err != nil {
			return err
		}
		// verify the key exists in the bucket
		value := todoBucket.Get([]byte(stringID))
		if value == nil {
			return fmt.Errorf("could not find task with id %v", stringID)
		}
		completedTask = string(value)
		// if so, we're good to delete!
		c := todoBucket.Cursor()
		_, _ = c.Seek([]byte(stringID))
		err = c.Delete()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Add to done bucket
	err = db.Update(func(tx *bolt.Tx) error {
		doneBucket, err := tx.CreateBucketIfNotExists([]byte("done"))
		if err != nil {
			return err
		}
		nextID, err := doneBucket.NextSequence()
		if err != nil {
			return err
		}
		doneID := []byte(strconv.Itoa(int(nextID)))
		err = doneBucket.Put(doneID, []byte(completedTask))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
