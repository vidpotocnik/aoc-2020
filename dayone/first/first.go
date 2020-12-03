package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	var puzzleItems []int

	file, err := os.Open("../puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		puzzleItem, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		puzzleItems = append(puzzleItems, puzzleItem)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(puzzleItems)
	var result = twoSum(puzzleItems, 2020)
	if result != nil {
		log.Print("2020 sum numbers: ", result)
		log.Print("Multiply result: ", result[0]*result[1])
	}
}

func twoSum(nums []int, target int) []int {
	var result []int
	if len(nums) <= 1 {
		return nil
	}
	for _, item := range nums {
		var currentSumItem = item
		for _, try := range nums {
			if (currentSumItem + try) == target {
				result = append(result, currentSumItem, try)
				return result
			}
		}
	}

	return nil
}
