package main

import (
	"bufio"
	"fmt"
	"os"
)

const ASCII_OFFSET_ABC = 64
const ASCII_OFFSET_XYZ_OUTCOME = 88

func main() {
	// open the input file
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// initialize scanner and counter variable
	sc := bufio.NewScanner(input)
	score := 0

	// main read loop
	for sc.Scan() {
		// save each game per line
		game := sc.Text()
		// save opponent move and game outcome to variables
		opponent := string(game[0])
		outcome := string(game[2])
		// evaluate the game and add the score
		gameScore := evalGame(opponent, outcome)
		fmt.Println(gameScore)
		score += gameScore
	}

	fmt.Printf("Score: %d\n", score)
}

// opponent = rock = 1, paper = 2, scissors = 3
// outcomeScore = loss = 0, tie = 3, win = 6
func evalGame(opponent string, outcome string) int {
	// convert the opponent's move into its score
	opponentScore := int(opponent[0]) - ASCII_OFFSET_ABC
	// convert the outcome into its score
	outcomeScore := (int(outcome[0]) - ASCII_OFFSET_XYZ_OUTCOME) * 3

	// set the playerMove assuming its a draw
	var playerMove int = opponentScore

	switch outcomeScore {
	// win, set playerMove to the move that's one more than the opponents
	// (rock -> paper, paper -> scissors, scissors -> rock)
	case 6:
		playerMove = (opponentScore + 1)
	// loss, set playerMove to the move that's one less than the opponents
	// (rock -> scissors, paper -> rock, scissors -> paper)
	case 0:
		playerMove = (opponentScore + 2)
	}
	// if we're over 3, subtract 3 to get the correct move
	if playerMove > 3 {
		playerMove -= 3
	}
	return playerMove + outcomeScore
}
