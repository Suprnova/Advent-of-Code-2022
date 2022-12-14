package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

// list object that can hold more lists or a value
type List struct {
	value    int
	elements []*List
	parent   *List
}

func main() {
	// open the input file
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// initialize the scanner, index counter, and sum
	sc := bufio.NewScanner(input)

	var masterList []List

	// read through the input file
	for sc.Scan() {
		if sc.Text() == "" {
			sc.Scan()
		}
		// parse the input, save it as a list
		masterList = append(masterList, parseInput(sc.Text()))
	}
	// append the two decoder packets
	dPacket1 := parseInput("[[2]]")
	dPacket2 := parseInput("[[6]]")
	masterList = append(masterList, dPacket1, dPacket2)

	// sort the list using the orderedStatus function
	sort.Slice(masterList, func(i, j int) bool {
		return orderedStatus(&masterList[i], &masterList[j]) == 2
	})

	decoderKey := 1
	// loop through the list, multiply the decoder key by the index if the elements are equal to the decoder packets
	for i, list := range masterList {
		if orderedStatus(&dPacket1, &list) == 1 || orderedStatus(&dPacket2, &list) == 1 {
			decoderKey *= (i + 1)
		}
	}
	println(decoderKey)
}

// function to parse the input into a list object
func parseInput(input string) List {
	// the "root" list, contains all other lists
	result := List{-1, []*List{}, nil}
	// set a pointer to the current list
	curList := &result
	// a string to hold the current number as the string is being read
	var n string
	for _, c := range input {
		// if the current character is a comma or a closing bracket, save the current number
		if c == ',' || c == ']' {
			if n != "" {
				// convert the number to an int, save it to the current list, and reset n
				fullNumber, _ := strconv.Atoi(n)
				curList.value = fullNumber
				n = ""
			}
			// exit out of the current list, set the current list to the parent instead
			curList = curList.parent
		}

		if c == '[' || c == ',' {
			// create a new list, add the new list to the parent, then navigate to the new list
			list := List{-1, []*List{}, curList}
			curList.elements = append(curList.elements, &list)
			curList = &list
		} else if c != ']' {
			// character isnt a bracket or comma, so it's a number. add it to n
			n += string(c)
		}
	}
	return result
}

// function to check the ordered status of two lists
// returns ints to indicate the status
// 0 = not ordered
// 1 = continue
// 2 = ordered
func orderedStatus(a *List, b *List) int {
	// if both lists have a value, (i.e., not -1) compare the values
	if a.value != -1 && b.value != -1 {
		if a.value > b.value {
			return 0
		} else if a.value == b.value {
			return 1
		} else {
			return 2
		}
	} else if a.value != -1 {
		// if a is a value and b is a list, make 'a' a list and recurse
		return orderedStatus(&List{-1, []*List{a}, nil}, b)
	} else if b.value != -1 {
		// if b is a value and a is a list, make 'b' a list and recurse
		return orderedStatus(a, &List{-1, []*List{b}, nil})
	} else {
		// both a and b are lists, compare them
		var i int
		// loop through each element as long as they're both in bounds
		for i = 0; i < len(a.elements) && i < len(b.elements); i++ {
			// check the ordered status of the current element
			status := orderedStatus(a.elements[i], b.elements[i])
			// a status of 1 means to continue, so only return if it's not 1
			if status != 1 {
				return status
			}
		}
		if i < len(a.elements) {
			// if we ran out of elements on b, then it's not ordered
			return 0
		} else if i < len(b.elements) {
			// if we ran out of elements on a, then it's ordered
			return 2
		}
		return 1
	}
}
