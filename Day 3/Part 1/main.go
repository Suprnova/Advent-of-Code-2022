package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const ASCII_OFFSET_LOWERCASE = 96
const ASCII_OFFSET_UPPERCASE = 38

func main() {
	// open the input file
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// initialize scanner and counter variable
	sc := bufio.NewScanner(input)
	sum := 0

	// main read loop
	for sc.Scan() {
		compartment := sc.Text()

		// for every rune in first half of compartment...
		for i := 0; i < len(compartment)/2; i++ {
			// save the rune
			char1 := compartment[i]
			// if the rune is *, we replaced it to avoid dupes, skip it
			if string(char1) == "*" {
				continue
			}
			// for every run in latter half of compartment...
			for j := len(compartment) / 2; j < len(compartment); j++ {
				// save the rune
				char2 := compartment[j]
				// check if the runes are equal
				if char1 == char2 {
					if char1 > ASCII_OFFSET_LOWERCASE {
						// the rune is lowercase, subtract the offset
						sum += int(char1) - ASCII_OFFSET_LOWERCASE
					} else {
						// the rune is uppercase, subtract the offset
						sum += int(char1) - ASCII_OFFSET_UPPERCASE
					}
					// replace the compared rune with * to avoid dupes
					compartment = strings.Replace(compartment, string(char1), "*", -1)
					// break to avoid checking for the same rune repeated elsewhere
					break
				}
			}
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
