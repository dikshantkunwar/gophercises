package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func readCSVFile(filepath string) [][]string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Unable to read input file: "+filepath, err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for: "+filepath, err)
	}
	return records
}

func main() {
	records := readCSVFile("test.csv")

	fmt.Println("Enter a time limit (Default is 20): ")
	var timeLimit int
	var timeLimitSeconds int
	var totalCorrectAnswers int

	fmt.Scanln(&timeLimit)

	timeLimitSeconds = 60 * timeLimit
	timer := time.NewTimer(time.Duration(timeLimitSeconds) * time.Second)

	for _, element := range records {
		select {
		case <-timer.C:
			fmt.Println("You got " + string(totalCorrectAnswers) + " correct answers!")
			return
		default:
			fmt.Println("Question " + element[0] + " = ? ")

			var answer int
			var correctAnswer int

			correctAnswer, err := strconv.Atoi(element[1])
			_, scanErr := fmt.Scanln(&answer)
			if scanErr != nil {
				fmt.Println("Error scanning in the answer! ", err)
			}
			if answer == correctAnswer {
				fmt.Println("Right answer! Good job")
				totalCorrectAnswers++
			} else {
				fmt.Println("Sorry wrong answer")
			}
		}
	}
	fmt.Println("Total correct answers: ", totalCorrectAnswers)

}
