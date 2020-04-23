// Reading a file 
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)
func main() {
	readFile, err := os.Open("questions")
	var score int
	var answer string
	var givenAnswer = []string{}
	var correctAnswers = []string{"F", "T", "T", "F", "F", "T", "F", "F", "T", "T", "F", "T"}
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
	readFile.Close()
	fmt.Println("Please answer True(t)/False(f) for below Questios")
	for _, line := range fileTextLines {
		fmt.Println(line)
		_, err := fmt.Scan(&answer)
		if err != nil {
			fmt.Printf("Something went wrong: %s\n", err)
		} else {
			givenAnswer = append(givenAnswer, answer)
		}
	}
	for i, v := range givenAnswer {
		if strings.ToUpper(v) == correctAnswers[i] {
			score++
		}
	}
	fmt.Println("Your score is: ", score)
}
