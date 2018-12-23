package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var importantNumbers map[int]int = map[int]int{
	10000: 0,
	1000:  0,
	100:   0,
	10:    0,
	5:     0,
	4:     0,
	3:     0,
	2:     0,
}

var importantCounts map[int]int = map[int]int{}

var i int = 0

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalf("Invalid args: %v", args)
	}
	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
		line := strings.Split(scanner.Text(), ":")
		count, _ := strconv.Atoi(line[1])

		if _, ok := importantNumbers[count]; ok {
			importantNumbers[count] = i
		}
		if i%10000000 == 0 {
			fmt.Println(i)
			importantCounts[i] = count
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	importantNumbersOut, err := json.MarshalIndent(importantNumbers, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	importantCountsOut, err := json.MarshalIndent(importantCounts, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", string(importantNumbersOut))
	fmt.Printf("%v\n", string(importantCountsOut))
}
