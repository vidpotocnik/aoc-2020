package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	var posRight [5]int
	var foundTrees [5]int
	var line int

	file, err := os.Open("../puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		for slopes := 0; slopes < 5; slopes++ {
			if slopes == 4 && line%2 != 0 {
				continue
			}
			if scanner.Text()[posRight[slopes]%len(scanner.Text())] == '#' {
				foundTrees[slopes]++
			}
			posRight[slopes] = posRight[slopes] + (2*slopes + 1) % 8
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Print(foundTrees[0] * foundTrees[1] * foundTrees[2] * foundTrees[3] * foundTrees[4])
}
