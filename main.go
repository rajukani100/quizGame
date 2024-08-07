package main

import (
	"bufio"
	"encoding/csv"
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
	file, err := os.Open("problem.csv")
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

	case <-time.After(time.Second * 50): //wait for 50 second
		fmt.Println("Time up's !")
	}
	fmt.Printf("\nYou got %v of %v", correct, len(parsedRecords))

}
