package main

import (
	"bufio"
	"fmt"
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

	// initialize scanner and counter variable
	sc := bufio.NewScanner(input)
	sum := 0

	// main read loop
	for sc.Scan() {
		input := sc.Text()
		// objective is to split the input into 4 sections
		// 0 and 1 are the min and max for range 1, 2 and 3 are the min and max for range 2
		input = strings.ReplaceAll(input, "-", ",")
		rangesStrings := strings.Split(input, ",")

		// convert the strings to an array of ints
		var ranges [4]int
		for i, s := range rangesStrings {
			ranges[i], _ = strconv.Atoi(s)
		}

		// if the min and max of one range are within the range of another, add to the sum
		if ranges[0] >= ranges[2] && ranges[1] <= ranges[3] {
			sum++
		} else if ranges[2] >= ranges[0] && ranges[3] <= ranges[1] {
			sum++
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
