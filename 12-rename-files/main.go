package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// Run after deleting samples dir
	err := createFiles()
	if err != nil {
		log.Fatalln(err)
	}

	// Print each filename in samples
	root := "./samples"
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			return nil
		}
		filePieces := strings.Split(path, "_")
		prefix := filePieces[0]
		suffix := filePieces[1]
		numberExtension := strings.Split(suffix, ".")
		stringNumber := numberExtension[0]
		number, err := strconv.Atoi(stringNumber)
		if err != nil {
			return err
		}
		var newName strings.Builder
		newName.WriteString(prefix)
		newName.WriteString(" - ")
		newName.WriteString(strconv.Itoa(number))
		newName.WriteString(".")
		newName.WriteString(numberExtension[1])
		err = os.Rename(path, newName.String())
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func createFiles() error {
	// Make new directory
	mkdir1 := exec.Command("mkdir", "samples")
	err := mkdir1.Run()
	if err != nil {
		return err
	}
	mkdir2 := exec.Command("mkdir", "samples/nested")
	err = mkdir2.Run()
	if err != nil {
		return err
	}

	// Add files at first level
	filenames := []string{"samples/birthday_001.txt", "samples/birthday_002.txt", "samples/birthday_003.txt", "samples/birthday_004.txt", "samples/birthday_005.txt", "samples/nested/birthday_006.txt", "samples/nested/birthday_007.txt", "samples/nested/birthday_008.txt", "samples/nested/birthday_009.txt", "samples/nested/birthday_010.txt"}
	for _, filename := range filenames {
		_, err = os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
