package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	puzzleFile, err := ioutil.ReadFile("../puzzle.txt")
	if err != nil {
		log.Print(err)
	}
	puzzleContent := string(puzzleFile)
	puzzles := strings.Split(puzzleContent, "\n\n")

	allowedFields := []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}
	validPassports := 0
	for _, passport := range puzzles {
		countCheck := 0
		for _, element := range allowedFields {
			if strings.Contains(passport, element) {
				countCheck++
			}
		}
		if countCheck >= len(allowedFields) {
			validPassports++
		}
	}

	log.Printf("There are %d valid passports.", validPassports)
}
