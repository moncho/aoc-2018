package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var movements = map[byte]xy{
	'N': xy{0, -1},
	'E': xy{1, 0},
	'S': xy{0, 1},
	'W': xy{-1, 0},
}

type xy struct {
	x, y int
}

func (this xy) move(other xy) xy {
	return xy{this.x + other.x, this.y + other.y}
}
func allDistances(r io.Reader) []int {
	current := xy{0, 0}
	type branch struct {
		xy       xy
		distance int
	}
	var branches []branch
	distance := 0
	distances := map[xy]int{
		current: 0,
	}

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)

	for s.Scan() {
		r := s.Text()[0]
		switch r {
		case 'N', 'E', 'S', 'W':
			distance++
			current = current.move(movements[r])
			if dist, ok := distances[current]; !ok || distance < dist {
				distances[current] = distance
			}
		case '(':
			branches = append(branches, branch{current, distance})
		case '|':
			//peek
			g := branches[len(branches)-1]
			distance, current = g.distance, g.xy
		case ')':
			//pop
			branches = branches[:len(branches)-1]
		}
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	result := make([]int, len(distances))
	i := 0
	for _, v := range distances {
		result[i] = v
		i++
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	routes := allDistances(f)

	//root := buildRoute(f)

	fmt.Printf("What is the largest number of doors you would be required to pass through to reach a room? %d\n", routes[len(routes)-1])

	//routes := root.allRoutes()

	routes = filter(routes, 1000)
	fmt.Printf("How many rooms have a shortest path from your current location that pass through at least 1000 doors? %d\n", len(routes))

}

func filter(routes []int, n int) []int {
	var res []int

	for _, r := range routes {
		if r >= n {
			res = append(res, r)
		}
	}

	return res
}
