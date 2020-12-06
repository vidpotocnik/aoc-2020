package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	puzzleFile, err := os.Open("../puzzle.txt")
	if err != nil {
		log.Print(err)
	}
	defer puzzleFile.Close()

	var s []string
	var data [][]string

	scanner := bufio.NewScanner(puzzleFile)

	for scanner.Scan() {
		element := scanner.Text()
		if element == "" {
			data = append(data, s)
			s = []string{}
		} else {
			s = append(s, element)
		}
	}
	data = append(data, s)

	log.Print("First task: ", countElements(data, firstTask))
	log.Print("Second task: ", countElements(data, secondTask))
}

func firstTask(puzzle []string) int {
	var groupedString string
	var result string

	for _, v := range puzzle {
		groupedString += v
	}
	for _, char := range groupedString {
		newChar := string(char)
		if !strings.Contains(result, newChar) {
			result += newChar
		}
	}

	return len(result)
}

func secondTask(puzzles []string) (result int) {
	dict := map[string]int{}

	for _, puzzle := range puzzles {
		for _, char := range puzzle {
			newChar := string(char)
			if _, ok := dict[newChar]; !ok {
				dict[newChar] = 1
			} else {
				dict[newChar] += 1
			}
		}
	}

	for _, v := range dict {
		if v == len(puzzles) {
			result++
		}
	}

	return result
}

func countElements(data [][]string, f func(i []string) int) (result int) {
	for _, e := range data {
		result += f(e)
	}
	return result
}
