package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var changes []int
	sum := 0
	firstRepeated := 0
	found := false
	freqSeen := make(map[int]bool)
	freqSeen[0] = true

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		changes = append(changes, i)
		sum += i
		if !freqSeen[sum] {
			freqSeen[sum] = true
		} else if !found {
			firstRepeated = sum
			found = true
		}
	}
	firstSol := sum
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for !found {
		for _, c := range changes {
			sum += c
			if !freqSeen[sum] {
				freqSeen[sum] = true
			} else if !found {
				firstRepeated = sum
				found = true
				break
			}
		}
	}

	fmt.Printf("Part1 sol: %d\n", firstSol)
	if found {
		fmt.Printf("Part2 sol: %d\n", firstRepeated)
	} else {
		fmt.Print("Part2 sol: not repeated\n")
	}
}
