package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type question struct {
	text   string
	answer string
}

func main() {
	filename := flag.String("csv", "questions.csv", "File name for quiz questions")
	timeLimitPtr := flag.Int("limit", 3, "Time limit for entire quiz, in seconds")
	flag.Parse()
	fmt.Println("Welcome to the Quiz")
	csvFile, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Error opening CSV file: %s\n", *filename)
		os.Exit(1)
	}
	csvReader := csv.NewReader(csvFile)
	var questions []question
	lines, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error parsing the provided CSV file.")
		os.Exit(1)
	}
	for _, line := range lines {
		questions = append(questions, question{
			text:   line[0],
			answer: line[1],
		})
	}
	var score int
	timeUpChan := make(chan struct{})
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press enter to start quiz clock")
	_, err = reader.ReadString('\n')
	if err == nil {
		go func() {
			limit := time.Duration(*timeLimitPtr) * time.Second
			fmt.Println(limit)
			time.Sleep(limit)
			timeUpChan <- struct{}{}
		}()
	}
	for _, question := range questions {
		fmt.Println("What is: ", question.text)
		answerCh := make(chan string)
		go func() {
			userAnswer, _ := reader.ReadString('\n')
			userAnswer = strings.TrimSuffix(userAnswer, "\n")
			answerCh <- userAnswer
		}()
		select {
		case <-timeUpChan:
			fmt.Println("\nTime's Up!")
			fmt.Printf("You correctly answered %d out of %d\n", score, len(questions))
			return
		case userAnswer := <-answerCh:
			if question.answer == userAnswer {
				score++
			}
		}
	}
	fmt.Printf("You correctly answered %d out of %d\n", score, len(questions))
}
