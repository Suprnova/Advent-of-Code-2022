package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var highestY int
var sandDropped int

func main() {
	// open the input file
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// initialize the scanner and cave
	sc := bufio.NewScanner(input)
	cave := make([][]rune, 500)
	for i := range cave {
		cave[i] = make([]rune, 1000)
		for j := range cave[i] {
			cave[i][j] = '.'
		}
	}

	for sc.Scan() {
		// read the input, format it into a list of coordinates [x1, y1, x2, y2, ...]
		input := sc.Text()
		input = strings.ReplaceAll(input, " -> ", ",")
		inputVals := strings.Split(input, ",")
		// start at the second coordinate since we use the previous coordinate as the base
		for i := 2; i < len(inputVals); i += 2 {
			var x1, y1, x2, y2 int
			// convert the coordinates into integers
			fmt.Sscanf(strings.Join(inputVals[i-2:i+2], " "), "%d %d %d %d", &x1, &y1, &x2, &y2)
			// draw the line represented by the coords
			drawLine(&cave, x1, y1, x2, y2)
			// update the highest y value if applicable
			if y2 > highestY {
				highestY = y2
			} else if y1 > highestY {
				highestY = y1
			}
		}
	}

	// start the initial dropping of sand at the source of 500, 0
	dropSand(&cave, 500, 0)
	println(sandDropped)
}

// function to drop sand and handle its movement
func dropSand(cave *[][]rune, x, y int) {
	// if we're past the bottom, we've hit the void, stop dropping sand
	if y > highestY {
		// we also decrement the counter since the sand never settled
		sandDropped--
		return
	}
	if (*cave)[y][x] == '.' {
		// if the space is empty, drop the sand
		(*cave)[y][x] = 'O'
		sandDropped++
		dropSand(cave, x, y)
	} else if (*cave)[y][x] == 'O' {
		// we're on a space with sand, process its movement
		if (*cave)[y+1][x] == '.' {
			// if the space below is empty, move the sand down one
			(*cave)[y][x] = '.'
			(*cave)[y+1][x] = 'O'
			dropSand(cave, x, y+1)
		} else if (*cave)[y+1][x-1] == '.' {
			// if the space below and to the left is empty, move the sand down and left one
			(*cave)[y][x] = '.'
			(*cave)[y+1][x-1] = 'O'
			dropSand(cave, x-1, y+1)
		} else if (*cave)[y+1][x+1] == '.' {
			// if the space below and to the right is empty, move the sand down and right one
			(*cave)[y][x] = '.'
			(*cave)[y+1][x+1] = 'O'
			dropSand(cave, x+1, y+1)
		} else {
			// if we can't move the sand, we know it's settled. drop a new piece of sand at the source
			dropSand(cave, 500, 0)
		}
	}
}

// function to draw a line on the cave
func drawLine(cave *[][]rune, x1, y1, x2, y2 int) {
	if x1 == x2 {
		// must be a vertical line
		// if y1 is greater than y2, swap them
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		// for every value between, create a rock tile
		for i := y1; i <= y2; i++ {
			(*cave)[i][x1] = '#'
		}
	} else {
		// must be a horizontal line
		// if x1 is greater than x2, swap them
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		// for every value between, create a rock tile
		for i := x1; i <= x2; i++ {
			(*cave)[y1][i] = '#'
		}
	}
}
