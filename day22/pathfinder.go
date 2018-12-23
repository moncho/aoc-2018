package main

import (
	"github.com/beefsack/go-astar"
)

var offsets = [][]int{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
}

func (e explorer) PathNeighbors() []astar.Pather {
	var neighbors []astar.Pather
	cave := e.cave

	for _, offset := range offsets {
		x := e.x + offset[0]
		y := e.y + offset[1]
		if x < 0 || y < 0 {
			continue
		}
		regionType := cave.regionType(x, y)
		ft := forbiddenTools[regionType]
		for _, t := range tools {
			if t != ft {
				n := explorer{
					cave: e.cave,
					tool: t,
					x:    x,
					y:    y,
				}
				neighbors = append(neighbors, n)
			}
		}

	}
	return neighbors
}

// PathNeighborCost returns the movement cost of the directly neighboring tile.
func (e explorer) PathNeighborCost(to astar.Pather) float64 {
	n := to.(explorer)
	cost := 1.0
	if e.tool != n.tool {
		cost += 7.0
	}
	return cost
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
// between non-adjacent nodes.
func (e explorer) PathEstimatedCost(to astar.Pather) float64 {
	toE := to.(explorer)
	absX := toE.x - e.x
	if absX < 0 {
		absX = -absX
	}
	absY := toE.y - e.y
	if absY < 0 {
		absY = -absY
	}
	distance := float64(absX + absY)
	if e.tool != toE.tool {
		distance += 7.0
	}
	return distance
}

type explorer struct {
	x, y int
	cave *cave
	tool tool
}

func rescue(cave *cave) explorer {
	explorer := explorer{
		x:    0,
		y:    0,
		cave: cave,
		tool: torch,
	}
	return explorer
}
