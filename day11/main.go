package main

import ()
import "fmt"

type cell struct {
	x, y int
	pow  int
}

func (c *cell) rackId() int {
	return c.x + 10
}

func (c *cell) powerLevel(serNo int) int {
	if c.pow == 0 {
		c.pow = powerLevel(c.x, c.y, serNo)
	}
	return c.pow
}

func powerLevel(x, y, serNo int) int {
	rackId := x + 10
	pl := rackId * y
	pl += serNo
	pl *= rackId
	pl = (pl / 100) % 10
	pl -= 5
	return pl
}

func main() {
	serNo := 1308
	gridSize := 300
	fmt.Printf("Grid %dx%d serial no: %d\n", gridSize, gridSize, serNo)

	grid := make([][]*cell, gridSize)
	var c cell
	for y := range grid {
		row := make([]*cell, gridSize)
		grid[y] = row
		for x := range row {
			c = cell{
				x: x + 1,
				y: y + y,
			}
			c.powerLevel(serNo)
			grid[y][x] = &c

		}
	}

	maxX, maxY := 0, 0
	maxPower := 0
	for gs := 0; gs < gridSize; gs++ {
		for y := 1; y <= gridSize-3; y++ {
			for x := 1; x <= gridSize-3; x++ {
				power := powerLevel(x, y, serNo) +
					powerLevel(x+1, y, serNo) +
					powerLevel(x+2, y, serNo) +
					powerLevel(x, y+1, serNo) +
					powerLevel(x+1, y+1, serNo) +
					powerLevel(x+2, y+1, serNo) +
					powerLevel(x, y+2, serNo) +
					powerLevel(x+1, y+2, serNo) +
					powerLevel(x+2, y+2, serNo)
				if power > maxPower {
					maxX = x
					maxY = y
					maxPower = power
				}
			}
		}
	}
	fmt.Printf("Top-left fuel cell found at: %dx%d \n", maxX, maxY)

}
