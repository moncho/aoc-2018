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

	p2 := part2Scores(firstElfRecipe, secondElfRecipe, attempts)
	fmt.Printf("recipes on the scoreboard to the left of the score sequence %d: %v\n", attempts, len(p2))

}
func part2Scores(score1, score2, expected int) []int {
	scores := make([]int, 2, expected+10)
	scores[0], scores[1] = 3, 7
	firstElfIndex := 0
	secondElfIndex := 1

	expectedI := toSlice(expected)
	le := len(expectedI)
	scoresL := len(scores)

	var currentRecipe int
	foundIndex := -1
	for {

		currentRecipe = scores[firstElfIndex] + scores[secondElfIndex]
		if currentRecipe > 9 {
			scores = append(scores, 1)
			scores = append(scores, currentRecipe%10)
		} else {
			scores = append(scores, currentRecipe)
		}
		scoresL = len(scores)
		firstElfIndex = (firstElfIndex + (scores[firstElfIndex] + 1)) % scoresL
		secondElfIndex = (secondElfIndex + (scores[secondElfIndex] + 1)) % scoresL

		if scoresL > le {
			foundIndex = index(scores[scoresL-le-1:], expectedI)
		}
		if foundIndex >= 0 {
			break
		}
	}
	return scores[:foundIndex+scoresL-le-1]
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

func eq(aa []int, bb []int) bool {
	if len(aa) != len(bb) {
		return false
	}
	for i, a := range aa {
		if a != bb[i] {
			return false
		}
	}
	return true
}

//toSlice returns an slice with the digits of
//the given number
func toSlice(n int) []int {
	var result []int
	for n != 0 {
		result = append(result, (n % 10))
		n = n / 10
	}
	//reverse
	for i := len(result)/2 - 1; i >= 0; i-- {
		opp := len(result) - 1 - i
		result[i], result[opp] = result[opp], result[i]
	}
	return result
}

//index returns, if present, the index of the first element of the
//given subarray in the given array. Returns -1 if the subarray is
//not present
func index(ints []int, sub []int) int {
	l := len(sub)
	for i := 0; i <= len(ints)-l; i++ {
		if eq(sub, ints[i:i+l]) {
			return i
		}
	}

	return -1
}
