package db

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

// Init db & buckets: todo & done
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("todo"))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("done"))
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

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("todo"))
		if bucket == nil {
			return errors.New("bucket does not exist")
		}
		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			task := Task{
				Key:   btoi(k),
				Value: string(v),
			}
			tasks = append(tasks, task)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func AddTask(taskString string) (Task, error) {
	var task Task
	task.Value = taskString
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("todo"))
		if err != nil {
			return err
		}
		// Ignore the error b/c only way that could happen is if we're inside View instead of Update, or the transaction was closed already
		id64, _ := bucket.NextSequence()
		id := int(id64)
		task.Key = id
		key := itob(id)
		err = bucket.Put(key, []byte(taskString))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return task, err
	}
	return task, nil
}

func MarkDone(key int) (Task, error) {

	var doneTask Task

	// Delete from todo, populate doneTask fields
	err := db.Update(func(tx *bolt.Tx) error {
		todoBucket, err := tx.CreateBucketIfNotExists([]byte("todo"))
		if err != nil {
			return err
		}
		// verify the key exists in the bucket
		bKey := itob(key)
		value := todoBucket.Get(itob(key))
		if value == nil {
			return fmt.Errorf("could not find task with id %d", key)
		}

		// iff we find the key in 'todo' bucket, populate doneTask
		doneTask.Key = key
		doneTask.Value = string(value)

		// if found, we're also good to delete!
		c := todoBucket.Cursor()
		_, _ = c.Seek(bKey)
		err = c.Delete()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return doneTask, err
	}

	// Add to done bucket
	err = db.Update(func(tx *bolt.Tx) error {
		doneBucket, err := tx.CreateBucketIfNotExists([]byte("done"))
		if err != nil {
			return err
		}
		// The only way a task will arrive in done bucket is via todo bucket. All tasks will have had unique ID's in todo bucket, so use their same ID in done bucket
		err = doneBucket.Put(itob(doneTask.Key), []byte(doneTask.Value))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return doneTask, err
	}
	return doneTask, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func Close() {
	db.Close()
}
