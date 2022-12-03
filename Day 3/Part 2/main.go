package main

import (
	"bufio"
	"fmt"
	"os"
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
		elf1 := sc.Text()
		sc.Scan()
		elf2 := sc.Text()
		sc.Scan()
		elf3 := sc.Text()

		// for each rune in the first elf's compartment...
		for i := 0; i < len(elf1); i++ {
			// for each rune in the second elf's compartment...
			for j := 0; j < len(elf2); j++ {
				// if they're equal...
				if elf1[i] == elf2[j] {
					// then for each rune in the third elf's compartment...
					for k := 0; k < len(elf3); k++ {
						// if they're equal...
						if elf1[i] == elf3[k] {
							// we have a rune present in all three elves' compartments
							if elf1[i] > ASCII_OFFSET_LOWERCASE {
								// the rune is lowercase, subtract the offset
								sum += int(elf1[i]) - ASCII_OFFSET_LOWERCASE
							} else {
								// the rune is uppercase, subtract the offset
								sum += int(elf1[i]) - ASCII_OFFSET_UPPERCASE
							}
							// break out of these nested loops to move on to the next group of 3
							goto next
						}
					}
				}
			}
		}
	next:
	}

	fmt.Printf("Sum: %d\n", sum)
}
