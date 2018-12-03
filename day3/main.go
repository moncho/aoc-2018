package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	id                        string
	leftDistance, topDistance int
	width, height             int
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var claims []claim
	fabric := make([][]int, 1000)

	//Build the list of claims and at the same time count the
	//number of claims on each square of the fabric
	for scanner.Scan() {
		c := parse(scanner.Text())
		claims = append(claims, c)
		for i := c.topDistance; i < (c.topDistance + c.height); i++ {
			row := fabric[i]
			if row == nil {
				row = make([]int, 1000)
				fabric[i] = row
			}
			for j := c.leftDistance; j < (c.leftDistance + c.width); j++ {
				row[j]++
			}
		}
	}
	targetClaim := ""
	for _, c := range claims {
		overlaps := false
		for i := c.topDistance; i < (c.topDistance+c.height) && !overlaps; i++ {
			row := fabric[i]
			for j := c.leftDistance; j < (c.leftDistance + c.width); j++ {
				if row[j] > 1 {
					overlaps = true
					break
				}
			}
		}

		if !overlaps {
			targetClaim = c.id
		}
	}

	count := overlapCount(fabric)
	fmt.Printf("Overlap inches: %d\n", count)
	fmt.Printf("ID of the only claim that doesn't overlap: %s\n", targetClaim)
}

//overlapCount counts the number of elements on the given matrix
//with a value higher than 1. Since the matrix is built by increasing
//an element value every time that a claim touches it, it serves as a way
//of counting the inches with overlapping claims.
func overlapCount(fabric [][]int) int {
	count := 0
	for _, rows := range fabric {
		for _, cols := range rows {
			if cols > 1 {
				count++
			}
		}
	}
	return count

}
func parse(s string) claim {
	ff := strings.Fields(s)
	if len(ff) != 4 {
		panic(
			fmt.Sprintf(
				"Unexpected claim %s", s))
	}
	claim := claim{}
	claim.id = ff[0][1:]
	distances := strings.Split(ff[2], ",")
	ld, err := strconv.Atoi(distances[0])

	if err != nil {
		panic(err)
	}

	claim.leftDistance = ld
	td, err := strconv.Atoi(strings.Replace(distances[1], ":", "", -1))
	if err != nil {
		panic(err)
	}
	claim.topDistance = td

	sizes := strings.Split(ff[3], "x")

	width, err := strconv.Atoi(sizes[0])

	if err != nil {
		panic(err)
	}

	claim.width = width
	height, err := strconv.Atoi(sizes[1])

	if err != nil {
		panic(err)
	}

	claim.height = height

	return claim
}
