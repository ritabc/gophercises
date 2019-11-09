package main

import (
	"gophercises/07-task-manager-cli/commands"
	"gophercises/07-task-manager-cli/db"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"

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
	// setup app info and commands
	info()
	app.Commands = commands.GetAll()

	// set dbPath
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "task.db")
	err := db.Init(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
