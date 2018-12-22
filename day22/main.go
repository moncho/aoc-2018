package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const (
	rocky regionType = iota
	wet
	narrow
)

type regionType int

type xy struct {
	x, y int
}
type cave struct {
	depth            int
	targetX, targetY int
	geologicIndices  map[xy]int
}

func (c cave) geologicIndex(x, y int) int {
	xy := xy{x, y}
	if index, ok := c.geologicIndices[xy]; ok {
		return index
	}
	index := 0
	if x == 0 && y == 0 {
	} else if x == c.targetX && y == c.targetY {
	} else if x == 0 {
		index = y * 48271
	} else if y == 0 {
		index = x * 16807
	} else {
		index = c.erosionLevel(x-1, y) * c.erosionLevel(x, y-1)
	}
	c.geologicIndices[xy] = index
	return index
}

func (c cave) erosionLevel(x, y int) int {
	return (c.geologicIndex(x, y) + c.depth) % 20183
}
func (c cave) regionType(x, y int) regionType {
	switch c.erosionLevel(x, y) % 3 {
	case 0:
		return rocky
	case 1:
		return wet
	case 2:
		return narrow
	}
	return -1
}

func (c cave) riskLevel() int {
	risk := 0
	for x := 0; x <= c.targetX; x++ {
		for y := 0; y <= c.targetY; y++ {
			risk += int(c.regionType(x, y))
		}
	}
	return risk
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	cave := newCave(f)

	fmt.Printf("Risk level: %d\n", cave.riskLevel())
}

func newCave(r io.Reader) *cave {
	c := cave{
		geologicIndices: make(map[xy]int),
	}
	s, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(s), "\n")

	var depth int
	var targetX, targetY int

	fmt.Sscanf(input[0], "depth: %d", &depth)
	fmt.Sscanf(input[1], "target: %d,%d", &targetX, &targetY)
	c.depth = depth
	c.targetX = targetX
	c.targetY = targetY
	return &c
}
