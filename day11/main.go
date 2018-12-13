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
	for y := range grid {
		row := make([]*cell, gridSize)
		grid[y] = row
		for x := range row {
			c := cell{
				x: x + 1,
				y: y + 1,
			}
			c.powerLevel(serNo)
			grid[y][x] = &c

		}
	}

	maxX, maxY := 0, 0
	maxPower := 0
	maxSize := 0
	for gs := 0; gs < gridSize; gs++ {
		for y := 0; y < gridSize-gs; y++ {
			for x := 0; x < gridSize-gs; x++ {
				power := 0
				for difx := gs; difx >= 0; difx-- {
					for dify := gs; dify >= 0; dify-- {
						power += grid[y+dify][x+difx].powerLevel(serNo)
					}
				}
				if power > maxPower {
					maxX = x
					maxY = y
					maxPower = power
					maxSize = gs
				}
			}
		}
	}
	fmt.Printf("Top-left fuel cell found at: %dx%d, grid side: %d\n", maxX+1, maxY+1, maxSize+1)

}
