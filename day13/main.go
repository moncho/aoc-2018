package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	faceUp    = '^'
	faceDown  = 'v'
	faceLeft  = '<'
	faceRight = '>'

	horizontal = '-'
	vertical   = '|'
	leftCurve  = '/'
	rightCurve = '\\'
	cross      = '+'
)

type cart struct {
	x, y  int
	turns int
	face  rune
}

func (c *cart) move(x, y int) {
	c.x = c
	c.y = y
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]rune
	var carts []*cart
	var line string
	lc := 0
	for scanner.Scan() {
		line = scanner.Text()
		var row []rune
		for pos, r := range line {
			if isCart(r) {
				carts = append(carts, &cart{
					x: pos,
					y: lc,
				})
				if r == faceUp || r == faceDown {
					r = vertical
				} else {
					r = horizontal
				}
			}
			row = append(row, r)
		}
		grid = append(grid, row)
		lc++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	//print(grid)

}

func print(runes [][]rune) {
	for _, line := range runes {
		fmt.Printf("%s\n", string(line))
	}
}

func isCart(r rune) bool {
	return r == faceUp || r == faceDown || r == faceLeft || r == faceRight
}
