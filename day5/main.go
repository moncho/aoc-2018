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

	var polymer string
	for scanner.Scan() {
		polymer = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", triggerReaction(polymer))
}

func triggerReaction(polymer string) string {
	return polymer
}
