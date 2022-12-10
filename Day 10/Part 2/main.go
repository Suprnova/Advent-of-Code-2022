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

	// initialize scanner, register, and next increment
	sc := bufio.NewScanner(input)
	register, nextIncrement := 1, 0

	// initialize the CRT, with 6 rows and 40 columns
	crt := make([][]rune, 6)
	for i := 0; i < 6; i++ {
		crt[i] = make([]rune, 40)
	}

	// go through each row of the CRT
	for i := 0; i < 6; i++ {
		// go through 40 cycles per row, populating columns 0-39
		// this abstracts away the true value of the cycle, since we're only interested in
		// the value of cycle % 40 when populating the CRT
		for cycle := 0; cycle < 40; cycle++ {
			// since the register is responsible for the center of the sprite, and the sprite is
			// 3 units wide, we can determine if the pixel should be lit based on if it's within
			// 1 unit of the cycle
			if abs(cycle-register) <= 1 {
				crt[i][cycle] = '#'
			} else {
				crt[i][cycle] = '.'
			}

			// if we have a next increment, add it to the register and move to the next cycle
			if nextIncrement != 0 {
				register += nextIncrement
				nextIncrement = 0
			} else {
				// sanity check to ensure we haven't exhausted our input before cycle 240
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
	}

	// print the CRT
	renderCRT(crt)
}

// function to render the CRT to the standard output
func renderCRT(crt [][]rune) {
	for _, row := range crt {
		for _, pixel := range row {
			print(string(pixel))
		}
		println()
	}
}

// function to return the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
