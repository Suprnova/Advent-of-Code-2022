/*
 * this was done on the hour of day 11, so the code isn't as clean as
 * it probably should be. please forgive me i don't do coding competitions
 * i just like to practice languages, i might create a post version to refactor
 * the general filth
 */

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// struct for the monkeys:
// current items, whether the operation is multiplication and/or squaring, the
// the second value of the operation, the divisor for the test, the monkey ID for a
// successful test, the monkey ID for a failed test, and the running total of inspections
type Monkey struct {
	items        []int
	oMultiply    bool
	oSquare      bool
	oModifier    int
	tDivisor     int
	tTrueMonkey  int
	tFalseMonkey int
	inspectCount int
}

func main() {
	// open the input file
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// initialize scanner and monkeys array
	sc := bufio.NewScanner(input)
	monkeys := make([]Monkey, 0)

	// loop through the input
	for sc.Scan() {
		// skip monkey ID line
		sc.Scan()
		// initialize items, remove commas and split into array
		items := make([]int, 0)
		itemString := strings.Trim(sc.Text(), " ")
		itemString = strings.ReplaceAll(itemString, ",", "")
		for _, item := range strings.Split(itemString, " ") {
			// this will catch the "Starting items:" sections, but it'll throw an error
			// so we can just ignore it and move on to the next item
			itemInt, err := strconv.Atoi(item)
			if err != nil {
				continue
			}
			items = append(items, itemInt)
		}
		// move on to the operation line, store the operation line in an array
		sc.Scan()
		operation := strings.Split(strings.Trim(sc.Text(), " "), " ")
		// check if the operation is multiplication
		multiply := operation[4] == "*"
		// check if the operation is squaring
		square := operation[5] == "old"
		// if operation is squaring, operation[5] will be "old", not a number
		// so Atoi will return 0, which is fine since we never use the modifier
		// for a square
		modifier, _ := strconv.Atoi(operation[5])
		// move on to the test line, store the test line in an array
		sc.Scan()
		test := strings.Split(strings.Trim(sc.Text(), " "), " ")
		// store the number to divide the worry level by
		divisor, _ := strconv.Atoi(test[3])
		// store the monkey ID for a successful test
		sc.Scan()
		trueMonkey, _ := strconv.Atoi(strings.Split(strings.Trim(sc.Text(), " "), " ")[5])
		// store the monkey ID for a failed test
		sc.Scan()
		falseMonkey, _ := strconv.Atoi(strings.Split(strings.Trim(sc.Text(), " "), " ")[5])
		// create the monkey and append it to the monkeys array
		monkeys = append(monkeys, Monkey{items, multiply, square, modifier, divisor, trueMonkey, falseMonkey, 0})
		sc.Scan()
	}

	// do 20 rounds
	for i := 0; i < 20; i++ {
		// go through each monkey starting from 0
		for monkeyId, monkey := range monkeys {
			// go through every item and run the evaluations
			for itemId, item := range monkey.items {
				// instead of removing items we just set them to -1, so we can skip
				// every item of value -1
				// i know this is bad, i fixed it in part 2
				if item == -1 {
					continue
				}
				// increment the inspection count
				monkeys[monkeyId].inspectCount++

				// do the operation, check if it's multiplying, squaring, or addition
				if monkey.oMultiply {
					if monkey.oSquare {
						monkeys[monkeyId].items[itemId] *= monkey.items[itemId]
					} else {
						monkeys[monkeyId].items[itemId] *= monkey.oModifier
					}
				} else {
					monkeys[monkeyId].items[itemId] += monkey.oModifier
				}

				// divide the worry level by 3
				monkeys[monkeyId].items[itemId] /= 3
				// check if the worry level is divisible by the divisor, then give the item to
				// the appropriate monkey
				if monkeys[monkeyId].items[itemId]%monkey.tDivisor == 0 {
					monkeys[monkey.tTrueMonkey].items = append(monkeys[monkey.tTrueMonkey].items, monkeys[monkeyId].items[itemId])
				} else {
					monkeys[monkey.tFalseMonkey].items = append(monkeys[monkey.tFalseMonkey].items, monkeys[monkeyId].items[itemId])
				}
				// we removed the item, just set it to -1
				monkeys[monkeyId].items[itemId] = -1
			}
		}
	}

	// find the two highest inspection counts
	firstHighest := 0
	secondHighest := 0
	for _, monkey := range monkeys {
		if monkey.inspectCount > firstHighest {
			// higher than first = downgrades first to second, sets first to new value
			secondHighest = firstHighest
			firstHighest = monkey.inspectCount
		} else if monkey.inspectCount > secondHighest {
			// higher than second = replace only the second
			secondHighest = monkey.inspectCount
		}
	}

	// print the product of the two highest inspection counts (solution)
	println(firstHighest * secondHighest)
}
