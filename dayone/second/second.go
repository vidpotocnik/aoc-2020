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
	threeSum(puzzleItems, 2020)
}

func threeSum(list []int, target int) {
	var result []int
	length := len(list)
	sort.Ints(list)
	for i := 0; i < length-2; i++ {

		low, high := i+1, length-1
		sum := target - list[i]
		for {
			if low >= high {
				break
			}

			switch {
			case list[low]+list[high] == sum:
				var sumNum = []int{list[i], list[low], list[high]}
				result = sumNum
				low++
				high = high - 1
			case list[low]+list[high] < sum:
				low++
			case list[low]+list[high] > sum:
				high = high - 1
			}

		}
	}
	
	if result != nil {
		log.Print("2020 sum numbers: ", result)
		log.Print("Multiply result: ", result[0]*result[1]*result[2])
	}
}
