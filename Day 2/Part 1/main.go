package main

import (
	"bufio"
	"fmt"
	"os"
)

const ASCII_OFFSET_ABC = 64
const ASCII_OFFSET_XYZ = 87

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
		// save each move to variable
		opponent := string(game[0])
		player := string(game[2])
		// evaluate the game and add the score
		score += evalGame(opponent, player)
	}

	fmt.Printf("Score: %d\n", score)
}

// loss = 0, tie = 3, win = 6
// rock = 1, paper = 2, scissors = 3
func evalGame(opponent string, player string) int {
	// convert the strings to ints based on the scoring of each move
	// (take their ascii value and subtract based on who's move it is)
	playerScore := int(player[0]) - ASCII_OFFSET_XYZ
	opponentScore := int(opponent[0]) - ASCII_OFFSET_ABC
	// it's a draw
	if opponentScore == playerScore {
		return playerScore + 3
	}
	switch opponentScore {
	// opponent played rock
	case 1:
		// player played paper
		if playerScore == 2 {
			return playerScore + 6
		}
	// opponent played paper
	case 2:
		// player played scissors
		if playerScore == 3 {
			return playerScore + 6
		}
	// opponent played scissors
	case 3:
		// player played rock
		if playerScore == 1 {
			return playerScore + 6
		}
	}
	// if we didn't already return, we know it's a loss
	return playerScore
}
