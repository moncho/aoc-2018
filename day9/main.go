package main

import (
	"bufio"
	"fmt"
	"os"
)

type marble struct {
	prev, next *marble
	val        int
}

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
	current := &marble{
		val: 0,
	}
	current.prev = current
	current.next = current
	prev, next := current, current

	for play := 1; play <= plays; play++ {
		if play%23 == 0 {
			for i := 0; i < 7; i++ {
				current = current.prev
			}
			prev = current.prev
			next = current.next
			scores[play%players] += play
			scores[play%players] += current.val
			prev.next, next.prev = next, prev
			current = next
		} else {
			prev, next = current.next, current.next.next
			current = &marble{
				val:  play,
				next: next,
				prev: prev,
			}
			prev.next = current
			next.prev = current
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
