package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type question struct {
	text   string
	answer string
}

func main() {
	filenamePtr := flag.String("filename", "questions.csv", "File name for quiz questions")
	flag.Parse()
	fmt.Println("Welcome to the CLI")
	fmt.Println("Filename Entered: ", *filenamePtr)
	csvFile, err := os.Open(*filenamePtr)
	if err != nil {
		fmt.Errorf("Error opening CSV file: ", err.Error)
	}
	reader := csv.NewReader(csvFile)
	var questions []question
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Errorf("Error reading question: ", err.Error)
		}
		questions = append(questions, question{
			text:   line[0],
			answer: line[1],
		})
	}
	var score int
	for _, question := range questions {
		fmt.Println("What is: ", question.text)
		ansReader := bufio.NewReader(os.Stdin)
		userAnswer, err := ansReader.ReadString('\n')
		if err != nil {
			fmt.Errorf("Error reading answer: ", err.Error)
		}
		userAnswer = strings.TrimSuffix(userAnswer, "\n")
		if question.answer == userAnswer {
			score++
		}
	}
	fmt.Printf("You correctly answered %d out of %d\n", score, len(questions))
}
