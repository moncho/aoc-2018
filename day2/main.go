package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	insCost = 1
	delCost = 1
	subCost = 1
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	twoCount := 0
	threeCount := 0
	var ids []string
	for scanner.Scan() {
		id := scanner.Text()
		ids = append(ids, id)
		two, three := repetitions(id)
		if two {
			twoCount++
		}
		if three {
			threeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	box1, box2 := findTwoClosest(ids)

	fmt.Printf("Checksum: %d\n", twoCount*threeCount)
	fmt.Printf("Box 1 ID: %s\nBox 2 ID: %s \n", box1, box2)
	fmt.Printf("Shared characters:  %s \n", common(box1, box2))

}

func repetitions(s string) (bool, bool) {
	twoRep, threeRep := false, false
	count := make(map[rune]int)
	for _, c := range s {
		count[c] += 1
	}
	for _, v := range count {
		if v == 2 {
			twoRep = true
		} else if v == 3 {
			threeRep = true
		}
	}

	return twoRep, threeRep
}

//common returns the characters that are present in both string
//in the same position, so: "ab", "aa" -> "a", "ab", "ba" -> ""
func common(s1, s2 string) string {
	var result []byte
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			result = append(result, s1[i])
		}
	}
	return string(result)
}

//findTwoClosest returns from the given list of strings the first two
//with just one different character.
func findTwoClosest(ids []string) (string, string) {
	for i, id := range ids {
		for j := i; j < len(ids); j++ {
			if levenshteinDistance(id, ids[j]) == 1 {
				return id, ids[j]
			}
		}
	}
	return "", ""
}

// levenshteinDistance returns the edit distance between source and target.
//
// It has a runtime proportional to len(source) * len(target) and memory use
// proportional to len(target).
func levenshteinDistance(source, target string) int {
	height := len(source) + 1
	width := len(target) + 1
	matrix := make([][]int, 2)

	// Initialize trivial distances (from/to empty string). That is, fill
	// the left column and the top row with row/column indices.
	for i := 0; i < 2; i++ {
		matrix[i] = make([]int, width)
		matrix[i][0] = i
	}
	for j := 1; j < width; j++ {
		matrix[0][j] = j
	}

	// Fill in the remaining cells: for each prefix pair, choose the
	// (edit history, operation) pair with the lowest cost.
	for i := 1; i < height; i++ {
		cur := matrix[i%2]
		prev := matrix[(i-1)%2]
		cur[0] = i
		for j := 1; j < width; j++ {
			delCost := prev[j] + delCost
			matchSubCost := prev[j-1]
			if source[i-1] != target[j-1] {
				matchSubCost += subCost
			}
			insCost := cur[j-1] + insCost
			cur[j] = min(delCost, min(matchSubCost, insCost))
		}
	}
	return matrix[(height-1)%2][width-1]
}

func min(a int, b int) int {
	if b < a {
		return b
	}
	return a
}
