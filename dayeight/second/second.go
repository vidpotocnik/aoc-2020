package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	arg       int
}

func getAfterTerminateValue(instructions *[]instruction) (bool, int) {
	instructionIndex := 0
	accumulator := 0
	executed := make(map[int]struct{}, len(*instructions))
	for {
		ins := (*instructions)[instructionIndex]

		if _, ok := executed[instructionIndex]; ok {
			return false, accumulator
		}
		executed[instructionIndex] = struct{}{}

		switch ins.operation {
		case "nop":
			instructionIndex++
		case "acc":
			accumulator += ins.arg
			instructionIndex++
		case "jmp":
			instructionIndex += ins.arg
		}
		if instructionIndex > len(*instructions)-1 {
			return true, accumulator
		}
	}
}

func main() {
	instructions := make([]instruction, 0)
	puzzleFile, _ := ioutil.ReadFile("../puzzle.txt")
	for _, instructionStr := range strings.Split(string(puzzleFile), "\n") {
		if instructionStr == "" {
			continue
		}
		ins := strings.Split(instructionStr, " ")
		operator, _ := strconv.Atoi(ins[1])
		instructions = append(instructions, instruction{ins[0], operator})
	}

	for index := range instructions {
		instruction := &(instructions)[index]

		switch instruction.operation {
		case "jmp":
			instruction.operation = "nop"
		case "nop":
			instruction.operation = "jmp"
		default:
			continue
		}

		if yes, accumulator := getAfterTerminateValue(&instructions); yes {
			log.Print("After terminate value of accumulator is: ", accumulator)
			return
		}

		switch instruction.operation {
		case "jmp":
			instruction.operation = "nop"
		case "nop":
			instruction.operation = "jmp"
		}
	}
}
