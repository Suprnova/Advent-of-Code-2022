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
	var highestCalories [3]int
	currentCalories := 0

	// for each token in scan
	for sc.Scan() {
		line := sc.Text()
		// if true, we are at the end of a group
		if line == "" {
			// check each of the 3 high scores
			for i := 0; i < 3; i++ {
				// if higher...
				if currentCalories > highestCalories[i] {
					// ...shift down all the scores below it...
					for j := i + 1; j < 3; j++ {
						highestCalories[j] = highestCalories[i]
					}
					// ...and set the new high score
					highestCalories[i] = currentCalories
					break
				}
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

	fmt.Printf("Highest 3 calories: %d\n", highestCalories)
	fmt.Printf("Highest calories total: %d\n", highestCalories[0]+highestCalories[1]+highestCalories[2])
}
