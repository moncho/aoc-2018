package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	open       = '.'
	tree       = '|'
	lumberyard = '#'
)

var offsets = [][]int{
	[]int{-1, -1},
	[]int{-1, 0},
	[]int{-1, 1},
	[]int{0, -1},
	[]int{0, 1},
	[]int{1, -1},
	[]int{1, 0},
	[]int{1, 1},
}

func main() {

	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	landscape := newLandscape(50, f)
	copy := copyLandscape(landscape)
	copy = simulate(10, copy)
	fmt.Printf("Total resource value after 10 min: %d\n", resourcesValue(copy))

	copy = copyLandscape(landscape)

	cycleLength, cycleStart := brent(
		func(landscape [][]rune) ([][]rune, int) {
			landscape = tick(landscape)
			c := resourcesValue(landscape)
			return landscape, c
		},
		copy)
	copy = copyLandscape(landscape)
	var minutes int = 1e9
	//for any duration, knowing when the cycle starts and the cycle length,
	//the simulation can be run just until the moment the cycle starts plus //the position of the duration in the cycle
	ticks := cycleStart + ((minutes - cycleStart) % cycleLength)
	copy = simulate(ticks, copy)
	fmt.Printf("Total resource value after %d min: %d\n",
		minutes,
		resourcesValue(copy))

}

func newLandscape(n int, r io.Reader) [][]rune {
	s := bufio.NewScanner(r)
	landscape := make([][]rune, n)
	lineCount := 0
	for s.Scan() {
		landscape[lineCount] = make([]rune, n)
		for i, t := range s.Text() {
			landscape[lineCount][i] = t
		}
		lineCount++
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return landscape
}

func simulate(duration int, landscape [][]rune) [][]rune {

	for i := 0; i < duration; i++ {
		landscape = tick(landscape)
	}
	return landscape
}

func tick(landscape [][]rune) [][]rune {
	width := len(landscape)
	copy := make([][]rune, width)
	for i, line := range landscape {
		copy[i] = make([]rune, width)
		for j, terrain := range line {
			trees := 0
			lumberyards := 0

			for _, offset := range offsets {
				ni := i + offset[0]
				nj := j + offset[1]
				if ni < 0 || ni >= width || nj < 0 || nj >= width {
					continue
				}
				if landscape[ni][nj] == tree {
					trees++
				} else if landscape[ni][nj] == lumberyard {
					lumberyards++
				}
			}
			newTerrain := terrain
			if terrain == open && trees >= 3 {
				newTerrain = tree
			} else if terrain == tree && lumberyards >= 3 {
				newTerrain = lumberyard
			} else if terrain == lumberyard && (lumberyards < 1 || trees < 1) {
				newTerrain = open
			}
			copy[i][j] = newTerrain
		}
	}

	return copy
}

func resourcesValue(landscape [][]rune) int {
	lumberyards := 0
	trees := 0
	for _, l := range landscape {
		for _, r := range l {
			switch r {
			case lumberyard:
				lumberyards++
			case tree:
				trees++
			}
		}
	}

	return trees * lumberyards
}

func copyLandscape(landscape [][]rune) [][]rune {
	copy := make([][]rune, len(landscape))
	for i, line := range landscape {
		copy[i] = make([]rune, len(landscape[i]))
		for j, r := range line {
			copy[i][j] = r
		}
	}
	return copy

}

func print(landscape [][]rune) {
	for _, line := range landscape {
		for _, r := range line {
			fmt.Printf("%c", r)
		}
		fmt.Print("\n")
	}
}
