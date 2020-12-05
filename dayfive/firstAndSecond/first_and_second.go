package main

import (
	"bufio"
	"log"
	"os"
)

type Options struct {
	Min int
	Max int
}

func main() {

	max := 0

	puzzleFile := "../puzzle.txt"
	puzzles, err := os.Open(puzzleFile)
	if err != nil {
		log.Fatal(err)
	}
	defer puzzles.Close()
	scanner := bufio.NewScanner(puzzles)
	allIds := make([]bool, 1000)

	min := 999999
	for scanner.Scan() {
		puzzle := scanner.Text()
		id := solveSeat(puzzle)

		if id < min {
			min = id
		}

		if id > max {
			max = id
		}

		allIds[id] = true
	}
	log.Println(max, min)

	for i := min; i < max-1; i++ {
		if allIds[i-1] && !allIds[i] && allIds[i+1] {
			log.Println("Seat ID is: ", i)
			break
		}
	}

}

func solveSeat(puzzle string) int {
	rows := Options{0, 128}
	cols := Options{0, 8}

	for _, char := range puzzle {
		switch string(char) {
		case "F":
			rows.Max -= (rows.Max - rows.Min) / 2
		case "B":
			rows.Min += (rows.Max - rows.Min) / 2
		case "L":
			cols.Max -= (cols.Max - cols.Min) / 2
		case "R":
			cols.Min += (cols.Max - cols.Min) / 2
		}
	}

	result := rows.Min*8 + cols.Min

	return result
}
