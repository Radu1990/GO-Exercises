package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

/* Simple Quiz Game written in GO

Create a program that will read in a quiz provided via a CSV file (more details below)
and will then give the quiz to a user keeping track of how many questions they get right
and how many they get incorrect. Regardless of whether the answer is correct or wrong
the next question should be asked immediately afterwards.

The CSV file should default to problems.csv (example shown below), but the user should
be able to customize the filename via a flag.

The CSV file will be in a format where the first column is a question and the second column
in the same row is the answer to that question.

You can assume that quizzes will be relatively short (< 100 questions)
and will have single word/number answers.

At the end of the quiz the program should output the total number of questions correct
and how many questions there were in total. Questions given invalid answers are considered incorrect.

Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds,
but should also be customizable via a flag.

Your quiz should stop as soon as the time limit has exceeded. That is, you shouldn’t wait
for the user to answer one final questions but should ideally stop the quiz entirely even
if you are currently waiting on an answer from the end user.

Users should be asked to press enter (or some other key) before the timer starts, and then
the questions should be printed out to the screen one at a time until the user provides an answer.
Regardless of whether the answer is correct or wrong the next question should be asked.

At the end of the quiz the program should still output the total number of questions correct and
how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.

*/

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of " +
		"'question / answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the QUIZ in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s \n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s= \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				fmt.Println("Correct!")
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem {
			q: line[0],
			a: line[1],
		}
	}
	return ret
}
