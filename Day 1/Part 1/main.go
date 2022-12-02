package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// open the input file
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// initialize scanner and counter variables
	sc := bufio.NewScanner(input)
	highestCalories := 0
	currentCalories := 0

	// for each token in scan
	for sc.Scan() {
		line := sc.Text()
		// if true, we are at the end of a group
		if line == "" {
			// if bigger than high score...
			if currentCalories > highestCalories {
				// ...set new high score
				highestCalories = currentCalories
			}
			// reset calorie count
			currentCalories = 0
			continue
		}
		// convert the token to an int
		count, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		// add to counter
		currentCalories += count
	}

	fmt.Printf("Highest calories: %d\n", highestCalories)
}
