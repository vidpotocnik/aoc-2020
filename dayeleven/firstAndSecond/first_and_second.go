package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	puzzleFile, err := os.Open("../puzzle.txt")
	if err != nil {
		panic(err)
	}
	defer puzzleFile.Close()

	scanner := bufio.NewScanner(puzzleFile)

	var lines [][]string
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		log.Print(err)
	}

	var firstResult int
	var secondResult int

	firstSeatsLayout, changed := firstTask(lines)
	for {
		firstSeatsLayout, changed = firstTask(firstSeatsLayout)
		if !changed {
			break
		}
	}

	firstResult = 0
	for row := range firstSeatsLayout {
		for col := range firstSeatsLayout[row] {
			if firstSeatsLayout[row][col] == "#" {
				firstResult++
			}
		}
	}

	secondSeatsLayout, changed := secondTask(lines)
	for {
		secondSeatsLayout, changed = secondTask(secondSeatsLayout)
		if !changed {
			break
		}
	}

	secondResult = 0
	for row := range secondSeatsLayout {
		for col := range secondSeatsLayout[row] {
			if secondSeatsLayout[row][col] == "#" {
				secondResult++
			}
		}
	}

	log.Print("First task: ", firstResult)
	log.Print("Second task: ", secondResult)
	elapsed := time.Since(start)
	log.Printf("Time elapsed: %s", elapsed)

}

func firstTask(seatsLayout [][]string) ([][]string, bool) {
	changedOrder := false
	newOrder := make([][]string, len(seatsLayout))

	for i := range seatsLayout {
		newOrder[i] = make([]string, len(seatsLayout[i]))
		copy(newOrder[i], seatsLayout[i])
	}

	for row := 0; row < len(seatsLayout); row++ {
		upperRow := row-1 >= 0
		bottomRow := row+1 < len(seatsLayout)

		for col := 0; col < len(seatsLayout[row]); col++ {
			if seatsLayout[row][col] == "." {
				continue
			}

			rightSeatsCol := col+1 < len(seatsLayout[row])
			leftSeatsCol := col-1 >= 0

			adjacentSeats := 0

			if upperRow {
				if seatsLayout[row-1][col] == "#" {
					adjacentSeats++
				}

				if rightSeatsCol {
					if seatsLayout[row-1][col+1] == "#" {
						adjacentSeats++
					}
				}

				if leftSeatsCol {
					if seatsLayout[row-1][col-1] == "#" {
						adjacentSeats++
					}
				}
			}

			if bottomRow {
				if seatsLayout[row+1][col] == "#" {
					adjacentSeats++
				}

				if rightSeatsCol {
					if seatsLayout[row+1][col+1] == "#" {
						adjacentSeats++
					}
				}

				if leftSeatsCol {
					if seatsLayout[row+1][col-1] == "#" {
						adjacentSeats++
					}
				}
			}

			if rightSeatsCol {
				if seatsLayout[row][col+1] == "#" {
					adjacentSeats++
				}
			}

			if leftSeatsCol {
				if seatsLayout[row][col-1] == "#" {
					adjacentSeats++
				}
			}

			if adjacentSeats >= 4 && seatsLayout[row][col] == "#" {
				changedOrder = true
				newOrder[row][col] = "L"
			} else if adjacentSeats == 0 && seatsLayout[row][col] == "L" {
				changedOrder = true
				newOrder[row][col] = "#"
			}
		}
	}

	return newOrder, changedOrder
}

func secondTask(seatsLayout [][]string) ([][]string, bool) {
	changedOrder := false
	newOrder := make([][]string, len(seatsLayout))

	for i := range seatsLayout {
		newOrder[i] = make([]string, len(seatsLayout[i]))
		copy(newOrder[i], seatsLayout[i])
	}

	for row := 0; row < len(seatsLayout); row++ {
		upperRow := row-1 >= 0
		bottomRow := row+1 < len(seatsLayout)

		for col := 0; col < len(seatsLayout[row]); col++ {
			if seatsLayout[row][col] == "." {
				continue
			}

			rightSeatsCol := col+1 < len(seatsLayout[row])
			leftSeatsCol := col-1 >= 0

			visibleSeats := 0

			if upperRow {
				if followSeat(-1, 0, row, col, seatsLayout) {
					visibleSeats++
				}

				if rightSeatsCol {
					if followSeat(-1, +1, row, col, seatsLayout) {
						visibleSeats++
					}
				}

				if leftSeatsCol {
					if followSeat(-1, -1, row, col, seatsLayout) {
						visibleSeats++
					}
				}
			}

			if bottomRow {
				if followSeat(+1, 0, row, col, seatsLayout) {
					visibleSeats++
				}

				if rightSeatsCol {
					if followSeat(+1, +1, row, col, seatsLayout) {
						visibleSeats++
					}
				}

				if leftSeatsCol {
					if followSeat(+1, -1, row, col, seatsLayout) {
						visibleSeats++
					}
				}
			}

			if rightSeatsCol {
				if followSeat(0, +1, row, col, seatsLayout) {
					visibleSeats++
				}
			}

			if leftSeatsCol {
				if followSeat(0, -1, row, col, seatsLayout) {
					visibleSeats++
				}
			}

			if visibleSeats >= 5 && seatsLayout[row][col] == "#" {
				changedOrder = true
				newOrder[row][col] = "L"
			} else if visibleSeats == 0 && seatsLayout[row][col] == "L" {
				changedOrder = true
				newOrder[row][col] = "#"
			}
		}
	}

	return newOrder, changedOrder
}

func followSeat(dirRow, dirCol, startingSeatsRow, startingSeatsCol int, seatsLayout [][]string) bool {
	row := startingSeatsRow
	col := startingSeatsCol

	for {
		row += dirRow
		col += dirCol

		if row < 0 || col < 0 {
			break
		}

		if row > len(seatsLayout)-1 || col > len(seatsLayout[row])-1 {
			break
		}

		if seatsLayout[row][col] == "L" {
			return false
		}

		if seatsLayout[row][col] == "#" {
			return true
		}
	}

	return false
}