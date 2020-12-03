package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var puzzleItems []string

	file, err := os.Open("../puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		puzzleItems = append(puzzleItems, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	validPasswords := 0

	for _, password := range puzzleItems {

		mainSplitted := strings.Split(password, " ")
		firstSplitted := strings.Split(mainSplitted[0], ":")
		ruleMinMaxSplitted := strings.Split(firstSplitted[0], "-")
		ruleCharSplitted := strings.Split(mainSplitted[1], ":")

		ruleChar := ruleCharSplitted[0]

		rulePassword := mainSplitted[2]

		ruleMin, err := strconv.Atoi(ruleMinMaxSplitted[0])
		if err != nil {
			log.Print(err)
		}
		ruleMax, err := strconv.Atoi(ruleMinMaxSplitted[1])
		if err != nil {
			log.Print(err)
		}

		if checkValidPassword(ruleMin, ruleMax, ruleChar, rulePassword) {
			validPasswords = validPasswords + 1
		}
	}

	log.Printf("There are %d valid passwords.", validPasswords)
}

func checkValidPassword(min int, max int, char string, password string) bool {
	charCount := strings.Count(password, char)
	if charCount >= min && charCount <= max {
		return true
	}

	return false
}
