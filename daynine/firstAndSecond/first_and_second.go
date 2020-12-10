package main

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	puzzleFile, err := os.Open("../puzzle.txt")
	if err != nil {
		log.Print(err)
	}
	puzzleReader, err := ioutil.ReadAll(puzzleFile)
	if err != nil {
		log.Print(err)
	}
	lines := strings.Split(string(puzzleReader), "\n")
	puzzles := make([]int, len(lines))

	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Print(err)
		}
		puzzles[i] = n
	}

	element := firstTask(puzzles, 25)
	log.Print("First task: ", element)
	log.Print("Second task: ", secondTask(puzzles, element))
}

func firstTask(data []int, preambleLength int) int {
loop:
	for startPoint := 0; startPoint < len(data)-preambleLength; startPoint++ {
		preamble := data[startPoint : startPoint+preambleLength]
		for i := 0; i < preambleLength-1; i++ {
			for j := i + 1; j < preambleLength; j++ {
				sum := preamble[i] + preamble[j]
				if data[startPoint+preambleLength] == sum {
					continue loop
				}
			}
		}
		return data[startPoint+preambleLength]
	}
	return -1
}

func sum(a []int) int {
	total := 0
	for _, i := range a {
		total += i
	}
	return total
}

func secondTask(puzzles []int, element int) int {
	for winSize := 2; winSize < len(puzzles); winSize++ {
		for start := 0; start < len(puzzles)-winSize; start++ {
			if sum(puzzles[start:start+winSize]) == element {
				result := make([]int, winSize)
				copy(result, puzzles[start:start+winSize])
				sort.Ints(result)
				return result[0] + result[winSize-1]
			}
		}
	}
	return -1
}
