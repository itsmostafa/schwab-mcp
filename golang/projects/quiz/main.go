// CSV Reader
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

// parseLines ...
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// exit: print message and exit with status code 1
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	// add flags "-csv <filename> and "-limit <time limit (seconds)>""
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	// open csv file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s ", *csvFilename))

	}
	r := csv.NewReader(file)

	// read entire csv file
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file.")
	}

	problems := parseLines(lines)

	// start timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// loop through problems
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d correctly.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.answer {
				correct++
			}
		}
	}
}
