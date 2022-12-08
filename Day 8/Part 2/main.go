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

	// initialize scanner, grid, and counter
	sc := bufio.NewScanner(input)
	treeGrid := make([][]int, 0)
	i := 0

	// loop through the input
	for sc.Scan() {
		// get the input line, and create a row in the tree grid for it
		inputLine := sc.Text()
		treeGrid = append(treeGrid, make([]int, len(inputLine)))

		// loop through every rune in the input line and save it to the grid as an int
		for j := 0; j < len(inputLine); j++ {
			treeGrid[i][j], _ = strconv.Atoi(string(inputLine[j]))
		}
		// increment the row counter
		i++
	}

	// initialize the sum of the visible trees
	highestValue := 0

	for rowKey, row := range treeGrid {
		for k := range row {
			// calculate the score for the tree
			score := checkScore(&treeGrid, rowKey, k)
			// if bigger than the current highest, replace it
			if score > highestValue {
				highestValue = score
			}
		}
	}

	fmt.Println("Highest score:", highestValue)
}

// function to check if a tree is visible
func checkScore(treeGrid *[][]int, i int, j int) int {
	// if the tree is on the edge, one of its scores is 0, don't waste compute time
	if (i == 0 || i == len(*treeGrid)-1) || (j == 0 || j == len((*treeGrid)[i])-1) {
		return 0
	}
	// initialize the height of the tree and the scenic score
	baseTree := (*treeGrid)[i][j]
	scenicScore := 1

	// check the scenic score south
	for k := i + 1; k < len(*treeGrid); k++ {
		if (*treeGrid)[k][j] >= baseTree || k == len(*treeGrid)-1 {
			scenicScore *= k - i
			break
		}
	}

	// check the scenic score north
	for k := i - 1; k >= 0; k-- {
		if (*treeGrid)[k][j] >= baseTree || k == 0 {
			scenicScore *= i - k
			break
		}
	}

	// check the scenic score east
	for k := j + 1; k < len((*treeGrid)[i]); k++ {
		if (*treeGrid)[i][k] >= baseTree || k == len((*treeGrid)[i])-1 {
			scenicScore *= k - j
			break
		}
	}

	// check the scenic score west
	for k := j - 1; k >= 0; k-- {
		if (*treeGrid)[i][k] >= baseTree || k == 0 {
			scenicScore *= j - k
			break
		}
	}
	return scenicScore
}
