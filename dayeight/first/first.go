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
	count     int
}

var accumulator = 0
var instructions = make([]instruction, 0)

func getAccumulatorValue() {
	instructionIndex := 0
	for {
		instruction := &instructions[instructionIndex]
		if instruction.count > 0 {
			log.Print("Accultor value is: ", accumulator)
			return
		}
		switch instruction.operation {
		case "nop":
			instructionIndex++
		case "acc":
			accumulator += instruction.arg
			instructionIndex++
		case "jmp":
			instructionIndex += instruction.arg
		}
		instruction.count++
	}
}

func main() {
	puzzleFile, _ := ioutil.ReadFile("../puzzle.txt")
	for _, instructionStr := range strings.Split(string(puzzleFile), "\n") {
		if instructionStr == "" {
			continue
		}
		ins := strings.Split(instructionStr, " ")
		operator, _ := strconv.Atoi(ins[1])
		instructions = append(instructions, instruction{ins[0], operator, 0})
	}
	getAccumulatorValue()
}
