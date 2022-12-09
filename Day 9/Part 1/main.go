// note: part 1 and part 2 have a different implementation. part 1 is
// not compatible with part 2, but part 2 is compatable with part 1 by
// simply changing the number of body parts from 9 to 1

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// the head of the rope
type Head struct {
	x int
	y int
}

// the tail of the rope
type Tail struct {
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

	// initialize scanner and rope head & tail
	sc := bufio.NewScanner(input)
	head := Head{100, 100}
	tail := Tail{100, 100}

	// loop through the input
	for sc.Scan() {
		inputLine := sc.Text()
		// format: "D 123", <direction> <steps>
		commands := strings.Split(inputLine, " ")
		direction := commands[0]
		steps, _ := strconv.Atoi(commands[1])
		// move the rope by the specified steps to the specified direction, save it to the head and tail
		head, tail = moveRope(head, tail, steps, direction)
	}

	// print the number of positions visited by the tail
	println("Positions visited:", len(positionsVisited))
}

func moveRope(head Head, tail Tail, steps int, direction string) (Head, Tail) {
	// if the direction is L or D, we subtract from the x or y coordinate
	negative := direction == "L" || direction == "D"
	// if the direction is vertical...
	if direction == "U" || direction == "D" {
		// move the rope every step
		for i := 0; i < steps; i++ {
			if negative {
				head.y--
			} else {
				head.y++
			}
			// we only moved the head, so we need to move the tail now
			tail = moveTail(head, tail)
			// save the position of the tail, so we can count how many positions it visited
			// duplicates won't show up, so counting the same position more than one is a non-issue
			positionsVisited[strconv.Itoa(tail.x)+","+strconv.Itoa(tail.y)] = true
		}
	} else {
		// the direction must be horizontal
		// move the rope every step
		for i := 0; i < steps; i++ {
			if negative {
				head.x--
			} else {
				head.x++
			}
			// as above, move the body after moving the head
			tail = moveTail(head, tail)
			// as above, save the position of the tail
			positionsVisited[strconv.Itoa(tail.x)+","+strconv.Itoa(tail.y)] = true
		}
	}
	// sanity check to ensure all parts of the rope are still connected
	if abs(head.x-tail.x) > 1 || abs(head.y-tail.y) > 1 {
		panic("Tail desynced")
	}
	return head, tail
}

// function to move the tail of the rope
func moveTail(head Head, tail Tail) Tail {
	// if the tail is already in the correct position, return it
	if abs(head.x-tail.x) <= 1 && abs(head.y-tail.y) <= 1 {
		return tail
	}

	// save a copy of the tail
	newTail := Tail{tail.x, tail.y}

	// horizontal movement checking
	if head.x > tail.x {
		newTail.x++
	} else if head.x < tail.x {
		newTail.x--
	}

	// vertical movement checking
	if head.y > tail.y {
		newTail.y++
	} else if head.y < tail.y {
		newTail.y--
	}
	// these movement steps are not mutually exclusive, since the tail needs to move diagonally
	// in some cases, so we check for both
	return newTail
}

// function to return the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
