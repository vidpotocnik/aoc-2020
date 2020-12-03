package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	var posRight int
	var foundTrees int

	file, err := os.Open("../puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text()[posRight%len(scanner.Text())] == '#' {
			foundTrees++
		}
		posRight += 3
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Print(foundTrees)
}
