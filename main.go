package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var importantNumbers map[int]int = map[int]int{}

var targetsList []int = []int{}

var keys []int = []int{}

var targetIndex int = 0

var i int = 0

func main() {
	// Grab the filename
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalf("Invalid args: %v", args)
	}

	// Open the file
	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Open a new file for the output
	file2, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	// Add the numbers we need to check for the targets
	for j := 100000; j > 0; j-- {
		targetsList = append(targetsList, j)
	}

	// Get the scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// How many rows have we checked?
		i++

		// Grab the count from the line
		line := strings.Split(scanner.Text(), ":")
		count, _ := strconv.Atoi(line[1])

		// If the pwned count is lower than the current target
		// which the slice is in descending order like the password list
		if count < targetsList[targetIndex] {
			// append the keys slice with the current target
			keys = append(keys, targetsList[targetIndex])
			// If we're more than one below the next target, we'll continue incrementing and appending
			for count < targetsList[targetIndex]-1 {
				targetIndex++
				keys = append(keys, targetsList[targetIndex])
				importantNumbers[targetsList[targetIndex]] = i
			}
			// move to the next target
			targetIndex++
		}

		// Store the number of passwords scanned in the key of the next "pwned count" target
		importantNumbers[targetsList[targetIndex]] = i

		// Every 10m just print the number to give a sense of when it will finish
		if i%10000000 == 0 {
			fmt.Println(i)
		}
	}
	// for the final target
	keys = append(keys, targetsList[targetIndex])

	// Reverse the order of the keys slice
	reverse(&keys)

	// Check if any of the scanner code caused an error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// Write CSV header
	file2.WriteString(fmt.Sprintf("%v,%v\n", "pwnedCount", "numberOfPasswords"))
	for _, k := range keys {
		// keys is sorted via ascending
		file2.WriteString(fmt.Sprintf("%v,%v\n", k, importantNumbers[k]))
	}
}

func reverse(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
