package main

import (
	"bufio"
	"os"
)

func main() {
	// open the input file
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// initialize scanner, grid, end position, and lowest step count
	sc := bufio.NewScanner(input)
	grid := make([][]byte, 0)
	endPos := make([]int, 2)
	lowestSteps := 0

	// read the input file and populate it into the grid
	for i := 0; sc.Scan(); i++ {
		grid = append(grid, make([]byte, 0))
		for j, c := range sc.Text() {
			if c == 'S' {
				// start position; set to 'a'
				c = 'a'
			} else if c == 'E' {
				// end position; store in endPos and set to 'z'
				endPos[1], endPos[0] = i, j
				c = 'z'
			}
			// regardless of what the character is, append it to the grid
			grid[i] = append(grid[i], byte(c))
		}
	}

	// loop through entire grid, find all 'a' nodes
	for y, row := range grid {
		for x, i := range row {
			if i == 'a' {
				// create a queue with a starting position of this 'a'
				queue := [][]int{{x, y, 0, 'a'}}
				// calculate the shortest path from this 'a' to the end position
				result := calculate(grid, []int{x, y}, endPos, queue)
				// if smaller than lowestSteps, and if result didn't return -1,
				// which means no path found, then update lowestSteps
				if (result < lowestSteps || lowestSteps == 0) && result != -1 {
					lowestSteps = result
				}
			}
		}
	}

	// print the result
	println(lowestSteps)
}

func calculate(base [][]byte, startPos, endPos []int, queue [][]int) int {
	// create a copy of the grid
	grid := make([][]byte, len(base))
	for i := range base {
		grid[i] = make([]byte, len(base[i]))
		copy(grid[i], base[i])
	}

	// QOL variables for width and height
	w := len(grid[0])
	h := len(grid)

	// while the queue is not empty
	for len(queue) > 0 {
		// pop the first element, store it in pos and update queue
		// pos contains x, y, steps, and the current letter
		pos := queue[0]
		queue = queue[1:]

		// breadth first search implementation:

		// for each of the four cardinal directions from the current position
		for _, c := range [][]int{{pos[0] - 1, pos[1]}, {pos[0], pos[1] - 1}, {pos[0] + 1, pos[1]}, {pos[0], pos[1] + 1}} {
			// initialize x and y with the values from the loop
			x, y := c[0], c[1]
			// check if out of bounds or on visited tile
			if y < 0 || y >= h || x < 0 || x >= w || grid[y][x] == '.' {
				continue
			}
			// check if the next letter is not more than 1 size greater than the current position
			if int(grid[y][x])-pos[3] > 1 {
				continue
			}
			// check if on the end tile
			if x == endPos[0] && y == endPos[1] {
				return pos[2] + 1
			}
			// the position is valid, add it to the queue
			queue = append(queue, []int{x, y, pos[2] + 1, int(grid[y][x])})
			// mark the position as visited
			grid[y][x] = '.'
		}
	}
	// no path found
	return -1
}
