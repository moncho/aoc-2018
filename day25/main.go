package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		fmt.Sscanf(s.Text(), "%d,%d,%d,%d", p.t, p.x, p.y, p.z)
		points = append(points, p)
	}
	if s.Err() != nil {
		panic(err)
	}
}

func constellations(points []point) int {
	included := make(map[point]bool)
	count := 0
	sort.Slice(points, func(i, j int) bool {
		pi := points[i]
		pj := points[j]
		if pi.t != pj.t {
			return pi.t < pj.t
		}
		if pi.t != pj.t {
			return pi.t < pj.t
		}
		if pi.x != pj.x {
			return pi.x < pj.x
		}
		if pi.y != pj.y {
			return pi.y < pj.y
		}
		return pi.z < pj.z
	})
	for i, p := range points {
		if _, ok := included[p]; ok {
			continue
		}

		found := false
		for j := i + 1; j < len(points); j++ {
			if p.distanceTo(points[j]) <= 3 {
				found = true
				included[p] = true
				included[points[j]] = true
			}
		}
		if found {
			count++
		}
	}

	return count
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
