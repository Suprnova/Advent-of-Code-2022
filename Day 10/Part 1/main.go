package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open the input file
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// initialize scanner, register, next increment and sum
	sc := bufio.NewScanner(input)
	register, nextIncrement := 1, 0
	sum := 0

	// start from cycle 1, loop up to cycle 220
	for cycle := 1; cycle <= 220; cycle++ {
		// if we're on cycle 20, 60, 100, 140, etc., report the register times the cycle
		if cycle%40 == 20 {
			sum += register * cycle
		}
		// if we have a next increment, add it to the register and move to the next cycle
		if nextIncrement != 0 {
			register += nextIncrement
			nextIncrement = 0
		} else {
			// sanity check to ensure we haven't exhausted our input before cycle 220
			if !sc.Scan() {
				panic("IO unexpectedly exhausted")
			}

			// read the next instruction, parse it into a command and an argument
			inputLine := sc.Text()
			commands := strings.Split(inputLine, " ")
			instruction := commands[0]
			// if the instruction is "addx", set the next increment to the argument
			if instruction == "addx" {
				nextIncrement, _ = strconv.Atoi(commands[1])
			}
			// otherwise, the instruction must be noop, so we do nothing and let the cycle increment
		}
	}

	// print the final sum
	println("Sum:", sum)
}
