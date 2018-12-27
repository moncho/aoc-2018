package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	t, x, y, z int
}

func (p point) distanceTo(other point) int {
	return abs(p.t-other.t) + abs(p.x-other.x) + abs(p.y-other.y) + abs(p.z-other.z)
}

func (p point) sortByDistanceTo(points []point) func(i, j int) bool {
	return func(i, j int) bool {
		return p.distanceTo(points[i]) < p.distanceTo(points[j])
	}
}
func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	var points []point
	for s.Scan() {
		var p point
		fmt.Sscanf(s.Text(), "%d,%d,%d,%d", &p.t, &p.x, &p.y, &p.z)
		points = append(points, p)
	}
	if s.Err() != nil {
		panic(err)
	}
	fmt.Printf("Number of constellations: %d\n", constellations(points))
}
func constellations(points []point) int {
	uf := NewWeightedQuickUnion(len(points))
	pointToInt := make(map[point]int)
	for i, p := range points {
		pointToInt[p] = i
		for point, id := range pointToInt {
			if p.distanceTo(point) <= 3 {
				uf.Union(i, id)
			}
		}
	}
	return uf.sets
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
