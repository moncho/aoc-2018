package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type coord struct {
	id   string
	x, y int
}

func (c coord) manhattanDistance(x, y int) int {
	return abs(c.x-x) + abs(c.y-y)
}

type distance struct {
	coordID     string
	distance    int
	distanceSum int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var coords []coord
	maxX, maxY := 0, 0
	counter := 0
	for scanner.Scan() {
		c := coord{
			id: strconv.Itoa(counter),
		}
		counter++
		_, err = fmt.Sscanf(scanner.Text(), "%d, %d\n", &c.x, &c.y)
		if err != nil {
			panic(err)
		}
		coords = append(coords, c)
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	grid := distancesGrid(coords, maxX+1, maxY+1)
	id, largest := largestFiniteArea(grid)
	c := find(coords, id)
	fmt.Printf("Location at (%d, %d) has the largest finite area: %d\n", c.x, c.y, largest)

	p2 := part2(grid)
	fmt.Printf("Size of the region with total distance less than 10.000: %d\n", p2)

}

func largestFiniteArea(distances [][]*distance) (string, int) {
	count := make(map[string]int)
	width, height := len(distances), len(distances[0])
	escapingCoords := map[string]struct{}{}
	for y, row := range distances {
		for x, d := range row {
			if d.coordID != "." {
				count[d.coordID]++
			}
			//if the coordinate is on the border of the grid, then
			//it extends forever and must be ignored
			if x == 0 || x == width || y == 0 || y == height {
				escapingCoords[d.coordID] = struct{}{}
			}

		}
	}
	largestCoord := ""
	largest := 0
	for k, c := range count {
		if _, ok := escapingCoords[k]; !ok && c > largest {
			largest = c
			largestCoord = k
		}
	}

	return largestCoord, largest
}

//distancesGrid generates a grid with the given height and width.
//For each position of the grid, the distance to the closest coordinate from the
//given slice of coordinates is calculated
func distancesGrid(coords []coord, width, height int) [][]*distance {
	grid := make([][]*distance, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]*distance, width)
	}

	for _, coord := range coords {
		for y, row := range grid {
			for x, d := range row {
				md := coord.manhattanDistance(x, y)
				if d == nil {
					grid[y][x] = &distance{coord.id, md, 0}
				} else if d.distance == md {
					d.coordID = "."
				} else if d.distance > md {
					d.coordID = coord.id
					d.distance = md
				}
				grid[y][x].distanceSum += md
			}
		}

	}

	return grid
}

//part2 returns ithe size of the region containing all locations which have a total distance to all given coordinates of less than 10000
func part2(distances [][]*distance) int {
	total := 0
	for _, row := range distances {
		for _, d := range row {
			if d.distanceSum < 10000 {
				total++
			}
		}
	}

	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func find(coords []coord, id string) coord {
	for _, c := range coords {
		if c.id == id {
			return c
		}
	}
	return coord{}
}
