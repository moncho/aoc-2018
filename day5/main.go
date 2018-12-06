package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	units = "abcdefghijklmnopqrstuvwxyz"
)

type reacts func(s string, i, j int) bool

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var polymer string
	for scanner.Scan() {
		polymer = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	collapsedPol := triggerReaction(polymer, reactionByPolarity())
	fmt.Printf("Unit count on reacted polymer: %d\n", len(collapsedPol))

	lowestLen := math.MaxInt32
	var lowestUnit rune
	for _, unit := range units {
		pol := cleanPolymer(polymer, unit)

		l := len(triggerReaction(pol, reactionByPolarity()))
		if l < lowestLen {
			lowestUnit = unit
			lowestLen = l
		}
	}
	fmt.Printf("Unit count on reacted polymer removing %q: %d\n", lowestUnit, lowestLen)
}

func triggerReaction(polymer string, reacts reacts) string {
	triggered := true
	for triggered {
		triggered = false
		for i := 0; i < len(polymer)-1; i++ {
			if reacts(polymer, i, i+1) {
				polymer = polymer[:i] + polymer[i+2:]
				triggered = true
			}
		}
	}
	return polymer
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func reactionByPolarity() reacts {
	return func(s string, i, j int) bool {
		return abs(int(s[i])-int(s[j])) == 32
	}
}

func cleanPolymer(polymer string, unit rune) string {
	s := string(unit)
	pol := strings.Replace(polymer, s, "", -1)
	return strings.Replace(pol, strings.ToUpper(s), "", -1)
}
