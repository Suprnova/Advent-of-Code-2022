package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// a unit of the rope
type Body struct {
	x int
	y int
}

// map used to store every position visited by the tail
var positionsVisited = make(map[string]bool)

func main() {
	// open the input file
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// initialize scanner and rope
	sc := bufio.NewScanner(input)
	rope := make([]Body, 10)
	// populate the rope with bodies at 0,0. rope[0] is the head, rope[0] is the tail
	for i := 0; i < 10; i++ {
		rope[i] = Body{0, 0}
	}

	// loop through the input
	for sc.Scan() {
		inputLine := sc.Text()
		// format: "D 123", <direction> <steps>
		commands := strings.Split(inputLine, " ")
		direction := commands[0]
		steps, _ := strconv.Atoi(commands[1])
		// move the rope by the specified steps to the specified direction
		rope = moveRope(rope, steps, direction)
	}

	// print the number of positions visited by the tail
	println("Positions visited:", len(positionsVisited))
}

func moveRope(rope []Body, steps int, direction string) []Body {
	// if the direction is L or D, we subtract from the x or y coordinate
	negative := direction == "L" || direction == "D"
	// if the direction is vertical...
	if direction == "U" || direction == "D" {
		// move the rope every step
		for i := 0; i < steps; i++ {
			if negative {
				rope[0].y--
			} else {
				rope[0].y++
			}
			// we only moved the head before, so we need to move the rest of the rope now
			rope = moveBody(rope)
			// save the position of the tail, so we can count how many positions it visited
			// duplicates won't show up, so counting the same position more than one is a non-issue
			positionsVisited[strconv.Itoa(rope[9].x)+","+strconv.Itoa(rope[9].y)] = true
		}
	} else {
		// the direction must be horizontal
		// move the rope every step
		for i := 0; i < steps; i++ {
			if negative {
				rope[0].x--
			} else {
				rope[0].x++
			}
			// as above, move the body after moving the head
			rope = moveBody(rope)
			// as above, save the position of the tail
			positionsVisited[strconv.Itoa(rope[9].x)+","+strconv.Itoa(rope[9].y)] = true
		}
	}
	// sanity check to ensure all parts of the rope are still connected
	// this probably has non-negligible performance impact, but it's good to have
	// and i'm not focused on performance
	for i := 0; i < 9; i++ {
		if abs(rope[i].x-rope[i+1].x) > 1 || abs(rope[i].y-rope[i+1].y) > 1 {
			panic("Body desynced")
		}
	}
	return rope
}

// function to move the body of the rope
func moveBody(rope []Body) []Body {
	// loop through every body in the rope besides the head
	for i := 1; i <= 9; i++ {
		// if the body is already in the correct position, we can stop moving the rope
		if abs(rope[i].x-rope[i-1].x) <= 1 && abs(rope[i].y-rope[i-1].y) <= 1 {
			break
		}

		// horizontal movement checking
		if rope[i-1].x > rope[i].x {
			rope[i].x++
		} else if rope[i-1].x < rope[i].x {
			rope[i].x--
		}

		// vertical movement checking
		if rope[i-1].y > rope[i].y {
			rope[i].y++
		} else if rope[i-1].y < rope[i].y {
			rope[i].y--
		}

		// these movement steps are not mutually exclusive, since the body needs to move diagonally
		// in some cases, so we check for both
	}
	return rope
}

// function to return the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
