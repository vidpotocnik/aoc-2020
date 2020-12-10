package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Counter struct {
	sum map[int]int
}

func main() {
	var allPuzzles []int

	puzzleFile := "../puzzle.txt"
	puzzles, err := os.Open(puzzleFile)
	if err != nil {
		log.Fatal(err)
	}
	defer puzzles.Close()
	scanner := bufio.NewScanner(puzzles)
	for scanner.Scan() {
		toInt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Print(err)
		}
		allPuzzles = append(allPuzzles, toInt)
	}

	firstTask, _ := firstTask(allPuzzles)
	secondTask, _ := secondTask(allPuzzles)

	log.Print("First task: ", firstTask)
	log.Print("Second task: ", secondTask)
}

func firstTask(puzzles []int) (int, error) {
	outletAndAdapter := prepareOutletAndAdapter(puzzles)
	diff1, diff3 := 0, 0
	for i := range outletAndAdapter[:len(outletAndAdapter)-1] {
		diff := outletAndAdapter[i+1] - outletAndAdapter[i]
		switch diff {
		case 3:
			diff3++
		case 0, 2:
			break
		case 1:
			diff1++
		default:
			return 0, fmt.Errorf("wrong adapter")
		}
	}

	return diff1 * diff3, nil
}

func secondTask(puzzles []int) (int, error) {

	sort.Ints(puzzles)
	outletAndAdapter := prepareOutletAndAdapter(puzzles)
	counter := &Counter{sum: make(map[int]int)}

	return counter.count(outletAndAdapter) + 1, nil
}

func prepareOutletAndAdapter(puzzles []int) []int {
	sort.Ints(puzzles)
	puzzles = append([]int{0}, puzzles...)               // OUTLET
	puzzles = append(puzzles, puzzles[len(puzzles)-1]+3) // ADAPTER

	return puzzles
}

func (c *Counter) count(nums []int) int {
	if counted, exists := c.sum[nums[0]]; exists {
		return counted
	}

	counted := 0
	for i := 0; i < len(nums)-1; i++ {
		for n := 1; i+n+1 < len(nums); n++ {
			diff := nums[i+n+1] - nums[i]
			if diff <= 3 {
				counted = counted + 1 + c.count(nums[i+n+1:])
			}
		}
	}

	c.sum[nums[0]] = counted
	return counted
}