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
	var marble int
	for scanner.Scan() {
		_, err = fmt.Sscanf(scanner.Text(), "%d players; last marble is worth %d points", &players, &marble)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("%d players; last marble is worth %d points: high score is %d\n", players, marble, highScore(players, marble))
}

func highScore(player, marble int) int {
	return 0
}
