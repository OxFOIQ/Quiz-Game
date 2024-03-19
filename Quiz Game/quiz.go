package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func exit (msg string){
	fmt.Println(msg)
	os.Exit(1)
}

type problems struct {
     answer string 
     question string
}

func parseLine(lines [][]string) []problems {
    resultt := make([]problems, len(lines))
    for index , line := range lines {
        resultt[index] = problems{
            question: line[0],
            answer: strings.TrimSpace(line[1]),
        }
    }
    return resultt
}

func main() {
    csvFilename := flag.String("csv", "problems.csv", "Questions&Answers File")
    timeLimit := flag.Int("time",15 , "time limit for the quiz in seconds")
    flag.Parse()
    file, err := os.Open(*csvFilename)
     if err != nil {
     exit(fmt.Sprintf("Failed to Open CSV File \n: %v", err))
    }

    reader := csv.NewReader(file)
    lines, err := reader.ReadAll()
    if err != nil {
        exit("Failed to Parse the CSV File")
    }

    problems := parseLine(lines)
    timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
    correct :=0
    problemloop :
    for index , problem := range problems {
        fmt.Printf("Problem #%d: %s\n", index+1, problem.question)
        answerCh := make(chan string)
        go func () {
            var answer string 
            fmt.Scanf("%s", &answer)
            answerCh <- answer
        }()
        select {
            case <-timer.C:
            fmt.Println()
            break problemloop
        case answer := <-answerCh:
            if answer == problem.answer{
                correct++
            }
        }
    }
    fmt.Printf("\nYou scored %d points out of %d tests \n" , correct , len(problems))
}
