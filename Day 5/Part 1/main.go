package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// custom Stack of bytes
type Stack []byte

// Check if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Peek the top of the stack
func (s *Stack) Peek() byte {
	if s.IsEmpty() {
		return ' '
	}
	return (*s)[len(*s)-1]
}

// Pop the top of the stack and return it
func (s *Stack) Pop() byte {
	if s.IsEmpty() {
		return ' '
	}
	e := s.Peek()
	*s = (*s)[:len(*s)-1]
	return e
}

// Push a new element onto the stack
func (s *Stack) Push(e byte) {
	*s = append(*s, e)
}

// Reverse the stack
func (s *Stack) Reverse() {
	if s.IsEmpty() {
		return
	}
	length := len(*s)
	for i := 0; i < length/2; i++ {
		(*s)[i], (*s)[length-i-1] = (*s)[length-i-1], (*s)[i]
	}
}

func main() {
	// open the input file
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// initialize scanner and stack collection
	sc := bufio.NewScanner(input)
	var stacks []Stack

	// read loop to populate the stacks
	for sc.Scan() {
		// read the input, split it into a number of stacks
		input := sc.Text()
		// reasoning: for every 4 characters in the input, we have a stack
		// plus one more at the end since that once doesn't contain the same padding
		// as long as at least one stack is provided, this works
		numberOfStacks := len(input)/4 + 1
		// initialize the stack collection if it hasn't been already
		if len(stacks) == 0 {
			stacks = make([]Stack, numberOfStacks)
		}

		// this would mean that we're past populating the stacks, so break out of the read loop
		if input[1] == '1' {
			break
		}

		// populate the stacks
		for i := 0; i < numberOfStacks; i++ {
			// this always returns the character that would be within []s in the input
			item := input[i*4+1]
			// a space means that this section of the input is empty, so don't add it
			if item == ' ' {
				continue
			} else {
				stacks[i].Push(item)
			}
		}
	}

	// since we were reading the input backwards, we need to reverse the stacks
	for _, stack := range stacks {
		stack.Reverse()
	}

	// skip the empty line after the stacks
	sc.Scan()

	// read loop to process instructions
	for sc.Scan() {
		instructions := sc.Text()
		// split the instructions by spaces
		instructionsSplit := strings.Split(instructions, " ")

		// populate variables of instructions with the values from the input
		count, _ := strconv.Atoi(instructionsSplit[1])
		sourceStack, _ := strconv.Atoi(instructionsSplit[3])
		destinationStack, _ := strconv.Atoi(instructionsSplit[5])

		// perform the instructions
		for i := 0; i < count; i++ {
			stacks[destinationStack-1].Push(stacks[sourceStack-1].Pop())
		}
	}

	// print the stacks
	for _, stack := range stacks {
		fmt.Print(string(stack.Peek()))
	}
}
