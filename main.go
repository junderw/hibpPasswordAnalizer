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

var targetsList []int = []int{
	100000,
	50000,
	10000,
	5000,
	1000,
	500,
	100,
	50,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

var keys []int = []int{}

var targetIndex int = 0

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

		if count < targetsList[targetIndex] {
			keys = append(keys, targetsList[targetIndex])
			targetIndex++
		}
		importantNumbers[targetsList[targetIndex]] = i
		if i%10000000 == 0 {
			fmt.Println(i)
		}
	}
	// for the final 1
	keys = append(keys, targetsList[targetIndex])
	reverse(&keys)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, k := range keys {
		fmt.Printf("%9v passwords were leaked %6v times or more\n", importantNumbers[k], k)
	}
}

func reverse(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
