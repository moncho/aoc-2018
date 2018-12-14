package main

import (
	"fmt"
)

func main() {
	attempts := 909441
	firstElfRecipe := 3
	secondElfRecipe := 7
	s := scores(firstElfRecipe, secondElfRecipe, attempts)
	fmt.Printf("scores of the ten recipes after %d recipes: %v\n", attempts, s)
}

func scores(score1, score2, attempts int) []int {
	scores := []int{score1, score2}
	count := len(scores)
	firstElfRecipe := 0
	secondElfRecipe := 1

	var currentRecipe int
	for count-10 < attempts {

		currentRecipe = scores[firstElfRecipe] + scores[secondElfRecipe]
		if currentRecipe > 9 {
			scores = append(scores, (currentRecipe/10)%10)
			scores = append(scores, currentRecipe-10)
		} else {
			scores = append(scores, currentRecipe)
		}
		count = len(scores)
		firstElfRecipe = (firstElfRecipe + (scores[firstElfRecipe] + 1)) % count
		secondElfRecipe = (secondElfRecipe + (scores[secondElfRecipe] + 1)) % count
	}

	return scores[count-10 : count]
}
