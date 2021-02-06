package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type questions struct {
	Question string
	Answer   string
}


func main() {
	questionAndAnswer := readCSV()
	totalQuestion := 0
	score := 0

	fmt.Println("================= Quiz Game ====================")
	for _, qa := range questionAndAnswer {
		totalQuestion++
		q := questions{
			Question: qa[0],
			Answer: qa[1],
		}
		ans := q.askQuestion()
		if q.isCorrect(ans) {
			score++
		}
		fmt.Printf("You score %v/%v \n", score, totalQuestion)
	}
	fmt.Println("================ Game Over ==============================")
	fmt.Printf("Your total score is %v out of %v", score, totalQuestion)
}

func (q questions) isCorrect(userInput string) bool {
	if userInput == q.Answer{
		return true
	}
	return false
} 

func (q questions) askQuestion() string {
	fmt.Print(q.Question + ": ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func readCSV() [][]string {
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return csvLines
}
