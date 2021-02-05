package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	questionAndAnswer := readCSV()
	totalQuestion := 0
	score := 0
	fmt.Println("================= Quiz Game ====================")
	for _, qa := range questionAndAnswer {
		totalQuestion++
		if userAnswer(qa[0]) == qa[1] {
			score++
		}
		fmt.Printf("You score %v/%v \n", score, totalQuestion)
	}
	fmt.Println("================ Game Over ==============================")
	fmt.Printf("Your total score is %v out of %v", score, totalQuestion)

}

type questions struct {
	Question string
	Answer   string
}

func userAnswer(question string) string {
	fmt.Print(question + ": ")
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
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return csvLines
}
