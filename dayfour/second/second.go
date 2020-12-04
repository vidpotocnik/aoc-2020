package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	data, err := ioutil.ReadFile("../puzzle.txt")
	if err != nil {
		log.Print(err)
	}
	lines := strings.Split(string(data), "\n\n")
	passports := make([]string, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		passports = append(passports, strings.ReplaceAll(line, "\n", " "))
	}

	fields := getAllowedFields()

	parsedPassports := retrievePassports(passports)
	validPassports := make([][]string, 0)

	validatedPassports := 0

	for _, passport := range parsedPassports {
		valid := allPresent(passport, fields)

		if valid {
			validPassports = append(validPassports, passport)
		}
	}

	for _, passport := range validPassports {
		valid := validatePassport(passport)

		if valid {
			validatedPassports++
		}
	}

	log.Printf("There are %d valid passports", validatedPassports)
}

func getAllowedFields() []string {
	return []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
}


func retrievePassports(passports []string) [][]string {
	newPassports := make([][]string, 0, len(passports))

	for _, passport := range passports {
		currFields := make([]string, 0)
		fields := strings.Split(passport, " ")

		for _, field := range fields {
			currFields = append(currFields, field)
		}

		newPassports = append(newPassports, currFields)
	}

	return newPassports
}

func allPresent(passport []string, fields []string) bool {
	for _, prop := range fields {

		_, match := findByValue(passport, prop)

		if !match {
			return false
		}
	}

	return true
}

func validatePassport(passport []string) bool {

	for _, val := range passport {
		currVal := strings.Split(val, ":")

		switch currVal[0] {
		case "byr":
			if !isBetween(currVal[1], 1920, 2002) {
				return false
			}
		case "iyr":
			if !isBetween(currVal[1], 2010, 2020) {
				return false
			}

		case "eyr":
			if !isBetween(currVal[1], 2020, 2030) {
				return false
			}
		case "hgt":
			if !validateHeight(currVal[1]) {
				return false
			}
		case "hcl":
			if !validateHairColor(currVal[1]) {
				return false
			}
		case "ecl":
			if !validateEyeColor(currVal[1]) {
				return false
			}
		case "pid":
			if !validateId(currVal[1]) {
				return false
			}
		}
	}

	return true
}

func validateId(val string) bool {
	isValid, _ := regexp.MatchString("^[0-9]{9}$", val)

	return isValid
}

func validateEyeColor(val string) bool {
	colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, color := range colors {
		if color == val {
			return true
		}
	}

	return false
}

func validateHairColor(val string) bool {
	isValid, _ := regexp.MatchString("^#[a-z0-9]{6}$", val)

	return isValid
}

func isBetween(val string, min int, max int) bool {
	intVal, err := strconv.Atoi(val)

	if err != nil {
		return false
	}

	return intVal >= min && intVal <= max
}

func validateHeight(val string) bool {
	if strings.HasSuffix(val, "in") {
		intVal, _ := strconv.Atoi(strings.ReplaceAll(val, "in", ""))

		return intVal >= 59 && intVal <= 76
	} else if strings.HasSuffix(val, "cm") {
		intVal, _ := strconv.Atoi(strings.ReplaceAll(val, "cm", ""))

		return intVal >= 150 && intVal <= 193
	}
	return false
}

func findByValue(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if strings.Contains(item, val) {
			return i, true
		}
	}
	return -1, false
}
