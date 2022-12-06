package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// open the input file
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// initialize scanner, strings, and counter
	sc := bufio.NewScanner(input)
	var buffer, lastFourteen string
	i := 0
	sc.Scan()
	buffer = sc.Text()

	// loop through the string constantly checking if the marker is present
	// with additional logic to ensure we don't index out of bounds
	for !containsMarker(lastFourteen) && i+14 < len(buffer) {
		// store last four runes and increment counter
		lastFourteen = buffer[i : i+14]
		i++
	}

	// this must mean we broke out of the loop due to reaching the end of the string
	// print an error
	if i+14 >= len(buffer) {
		fmt.Println("No marker found")
	} else {
		// we found the marker, it's the last character of the lastFourteen string, so it's
		// in the buffer string at index i+3
		fmt.Println("Marker found at position", i+13)
	}
}

// function to check if the string contains a marker (all runes are unique)
func containsMarker(lastFourteen string) bool {
	// edge case for checking while we're still building the string
	if len(lastFourteen) < 14 {
		return false
	}
	// for each rune in the string...
	for i := 0; i < 13; i++ {
		// for every rune after it...
		for j := i + 1; j < 14; j++ {
			// ensure they don't match, and if they do, return false
			if lastFourteen[i] == lastFourteen[j] {
				return false
			}
		}
	}
	// if we made it here, all runes are unique, so return true
	return true
}
