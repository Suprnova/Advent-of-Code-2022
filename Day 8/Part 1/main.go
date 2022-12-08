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
	sum := 0

	for rowKey, row := range treeGrid {
		if rowKey == 0 || rowKey == len(treeGrid)-1 {
			// if the row is the first or last row, all trees are visible
			sum += len(row)
		} else {
			for k := range row {
				if k == 0 || k == len(row)-1 {
					// if the tree is the first or last tree in the row, it is visible
					sum++
				} else {
					if checkVisibility(&treeGrid, rowKey, k) {
						// if checkVisibility returns true, the tree is visible
						sum++
					}
				}
			}
		}
	}

	fmt.Println("Sum:", sum)
}

// function to check if a tree is visible
func checkVisibility(treeGrid *[][]int, i int, j int) bool {
	// initialize the height of the tree and the visibility bool
	baseTree := (*treeGrid)[i][j]
	visible := true

	// check for visibility south
	for k := i + 1; k < len(*treeGrid); k++ {
		if (*treeGrid)[k][j] >= baseTree {
			visible = false
			break
		}
	}
	if visible {
		return true
	}
	visible = true

	// check for visibility north
	for k := i - 1; k >= 0; k-- {
		if (*treeGrid)[k][j] >= baseTree {
			visible = false
			break
		}
	}
	if visible {
		return true
	}
	visible = true

	// check for visibility east
	for k := j + 1; k < len((*treeGrid)[i]); k++ {
		if (*treeGrid)[i][k] >= baseTree {
			visible = false
			break
		}
	}
	if visible {
		return true
	}
	visible = true

	// check for visibility west
	for k := j - 1; k >= 0; k-- {
		if (*treeGrid)[i][k] >= baseTree {
			visible = false
			break
		}
	}
	return visible
}
