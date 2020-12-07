package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	name     string
	contains map[string]int
}

func main() {
	puzzleFile, err := os.Open("../puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer puzzleFile.Close()

	bagsCollection := make(map[string]Bag)
	targetBag := "shiny gold"

	scanner := bufio.NewScanner(puzzleFile)

	for scanner.Scan() {
		line := scanner.Text()
		formattedLine := strings.Split(line, " bags contain ") // let's remove `bags contain` from string
		color := formattedLine[0]                              // take out bag color from formatted string
		separators := strings.Split(formattedLine[1], ", ")
		currentBag := Bag{color, make(map[string]int)} // make map with curret bag color
		for _, value := range separators {             // counted and separated color bags
			if value == "no other bags." {
				break
			}
			words := strings.Split(value, " ")
			containedColor := words[1] + " " + words[2]
			number, err := strconv.Atoi(words[0])
			if err != nil {
				log.Print(err)
			}
			currentBag.contains[containedColor] = number
		}
		bagsCollection[color] = currentBag // add bags inside parent bag so we have collection of bags counted
	}

	possibleOptions := 0
	for _, currentBag := range bagsCollection {
		if currentBag.name == targetBag {
			continue
		}
		containingBags := getContainingBags(bagsCollection, currentBag, targetBag)
		if containingBags > 0 {
			possibleOptions++
		}
	}
	log.Print(possibleOptions, " bag colors can eventually contain at least one shiny gold bag.")

	individualBagsInside := getBags(bagsCollection, bagsCollection[targetBag])

	log.Print(individualBagsInside, " individual bags are required inside a single shiny gold bag.")

	if err := scanner.Err(); err != nil {
		log.Print(err)
	}
}

func getContainingBags(bags map[string]Bag, currentBag Bag, target string) int {
	count := 0
	for color := range currentBag.contains {
		subBag := bags[color]
		if color == target {
			count++
		} else {
			count += getContainingBags(bags, subBag, target)
		}
	}
	return count
}

func getBags(bags map[string]Bag, currentBag Bag) int {
	count := 0
	for color, i := range currentBag.contains {
		subBag := bags[color]
		count += i + i*getBags(bags, subBag)
	}

	return count
}
