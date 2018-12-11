package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var players int
	var lastMarble int
	for scanner.Scan() {
		_, err = fmt.Sscanf(scanner.Text(), "%d players; last marble is worth %d points", &players, &lastMarble)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("%d players; last marble is worth %d points: high score is %v\n", players, lastMarble, highestScore(scores(players, lastMarble)))

	fmt.Printf("%d players; last marble is worth %d points: high score is %v\n", players, 100*lastMarble, highestScore(scores(players, 100*lastMarble)))
}

func highestScore(scores []int) int {
	topScore := 0
	for _, score := range scores {
		if score > topScore {
			topScore = score
		}
	}

	return topScore
}

func scores(players, plays int) []int {
	scores := make([]int, players)
	board := make(map[int]int)
	board[0] = 0
	board[1] = 2
	board[2] = 1
	currentPos := 1
	for marble := 3; marble <= plays; marble++ {
		boardLen := len(board)
		if marble%23 == 0 {
			if currentPos < 7 {
				currentPos = boardLen - (7 - currentPos)
			} else {
				currentPos = currentPos - 7
			}
			scores[marble%players] += marble
			scores[marble%players] += board[currentPos]
			for i := currentPos; i < boardLen; i++ {
				board[i] = board[i+1]
			}
			delete(board, boardLen-1)
			continue
		}
		currentPos = nextPosition(currentPos, boardLen)
		if currentPos == marble {
			board[currentPos] = marble
		} else {
			m := marble
			for i := currentPos; i <= boardLen; i++ {
				board[i], m = m, board[i]
			}
		}

	}

	return scores
}

func nextPosition(cursor, count int) int {
	switch count {
	case 0:
		return 0
	case 1:
		return 1
	case cursor + 1:
		return 1
	}

	return cursor + 2
}
