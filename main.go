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

type problem struct {
	q string
	a string
}

func parseRecords(records [][]string) []problem {
	problemsList := make([]problem, len(records))
	for i, v := range records {
		problemsList[i] = problem{
			q: v[0],
			a: v[1],
		}
	}
	return problemsList
}

func main() {
	fileName := flag.String("f", "problem.csv", "provide csv file name")
	var quizTime int
	flag.IntVar(&quizTime, "t", 50, "provide time in seconds")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("error", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err1 := reader.ReadAll()
	if err1 != nil {
		fmt.Println("error", err1)
	}

	parsedRecords := parseRecords(records)

	correct := 0
	inputReader := bufio.NewReader(os.Stdin)

	ch := make(chan int)

	go func(ch chan int) {
		for i, v := range parsedRecords {
			fmt.Printf("Problem %v: %s ->", i+1, v.q)
			ans, _ := inputReader.ReadString('\n')
			ans = strings.TrimSpace(ans)
			if ans == v.a {
				correct++
			}
		}
		ch <- 0
	}(ch)

	select {
	case correct = <-ch:

	case <-time.After(time.Duration(quizTime * int(time.Second))): //wait for 50 second
		fmt.Println("Time up's !")
	}
	fmt.Printf("\nYou got %v of %v", correct, len(parsedRecords))

}
