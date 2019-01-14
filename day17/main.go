package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
)

const (
	clay      = '#'
	sand      = '.'
	spring    = '+'
	water     = '~'
	waterFlow = '|'
	springX   = 500
	springY   = 0
)

type xy struct {
	x, y int
}
type scanLine struct {
	fromX, fromY int
	toX, toY     int
}

type groundMap struct {
	terrain    [][]rune
	minY, maxY int
	minX, maxX int
}

func (g groundMap) depth() int {
	//The last level is below the clay level
	//and it is not of interest
	return len(g.terrain) - 1
}

func (g groundMap) canWaterFlow(x, y int) bool {
	if x >= g.width() {
		return false
	}
	if y >= g.depth() {
		return false
	}
	return g.terrain[y][x] == sand || g.terrain[y][x] == waterFlow
}
func (g groundMap) width() int {
	return len(g.terrain[0])
}
func (g groundMap) spring() (int, int) {
	return springX - g.minX + 1, springY
}
func (g groundMap) waterReach() (int, int) {
	waterCount := 0
	flowCount := 0
	for y := 0; y < g.depth(); y++ {
		for x := 0; x < g.width(); x++ {
			if g.terrain[y][x] == water {
				waterCount++
			} else if g.terrain[y][x] == waterFlow {
				flowCount++
			}
		}
	}

	return waterCount, flowCount
}
func main() {

	f, err := os.Open("input.txt")
	defer f.Close()

	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)

	var clay []xy
	lowestY, maxY := math.MaxInt32, 0
	lowestX, maxX := math.MaxInt32, 0
	for s.Scan() {
		line := s.Text()
		var s scanLine
		if strings.HasPrefix(line, "x=") {
			fmt.Sscanf(line, "x=%d, y=%d..%d", &s.fromX, &s.fromY, &s.toY)
			s.toX = s.fromX
		} else {
			fmt.Sscanf(line, "y=%d, x=%d..%d", &s.fromY, &s.fromX, &s.toX)
			s.toY = s.fromY
		}
		if s.fromY < lowestY {
			lowestY = s.fromY
		}
		if s.toY > maxY {
			maxY = s.toY
		}
		if s.fromX < lowestX {
			lowestX = s.fromX
		}
		if s.toX > maxX {
			maxX = s.toX
		}
		clay = append(clay, scanToXy(s)...)
	}
	if s.Err() != nil {
		panic(err)
	}
	sort.Slice(clay, func(i, j int) bool {
		if clay[i].y == clay[j].y {
			return clay[i].x < clay[j].x
		}
		return clay[i].y < clay[j].y
	})

	groundMap := newGroundMap(clay, lowestX, lowestY, maxX, maxY)

	x, y := groundMap.spring()
	simulateWaterFlow(x, y, groundMap)
	wc, fc := groundMap.waterReach()
	fmt.Printf("How many tiles can the water reach within the range of y values in your scan? %d %d, sum: %d\n", wc, fc, wc+fc)
}
func newGroundMap(clayScan []xy, minX, minY, maxX, maxY int) groundMap {
	ground := groundMap{
		minY: minY,
		maxY: maxY,
		minX: minX,
		maxX: maxX,
	}
	//Map starts at ground level(Y = 0), where the spring is located and goes
	//down infinitely. But since levels below the last level where there is clay
	//are not of interest the depth is set to be one level below the last level
	//with clay.
	depth := maxY - minY + 3
	//Water might spread to the right or left of the clay
	width := maxX - minX + 3
	lines := make([][]rune, depth)

	for line := 0; line < depth; line++ {
		lines[line] = make([]rune, width)
		for i := 0; i < width; i++ {
			lines[line][i] = sand
		}
	}

	lines[springY][springX-minX+1] = spring
	for _, xy := range clayScan {
		lines[xy.y-minY+1][xy.x-minX+1] = clay
	}
	ground.terrain = lines
	return ground
}
func simulateWaterFlow(springX, springY int, gm groundMap) {
	//Sources of water flow, the initial source of water is the
	//spring. Additional sources might appear as water overflows
	//clay layers.
	sources := []xy{xy{springX, springY}}
	for len(sources) > 0 {
		sourceCount := len(sources)
		source := sources[sourceCount-1]
		sources = sources[:sourceCount-1]
		x, y := source.x, source.y
		//If still water is found there is nothing left to do
		//for this source
		if gm.terrain[y][x] == water {
			continue
		}

		//The water flows down, creating new sources of water
		//where it overflows
		y++
		for y < gm.depth() {
			terrain := gm.terrain[y][x]
			//Sand found, a flow is created on this level, explore the next level
			if terrain == sand {
				gm.terrain[y][x] = waterFlow
				y++
			} else if terrain == waterFlow {
				//A water flow is already present, simulation is back from the level below
				//Current source has filled everything with water as much as it could,
				//simulation will continue with other sources, if any.
				break
			} else {
				//Clay layer, left and right positions are explored to check
				//whether the current flow overflows (and creates new sources)
				//or not and it creates a layer of still water
				y--
				left := x
				overFlows := false
				for gm.canWaterFlow(left, y) {
					if gm.terrain[y+1][left] == sand {
						sources = append(sources, xy{left, y})
						overFlows = true
						left--
						break
					}
					left--
				}
				right := x + 1
				for gm.canWaterFlow(right, y) {
					if gm.terrain[y+1][right] == sand {
						sources = append(sources, xy{right, y})
						overFlows = true
						right++
						break
					}
					right++
				}

				for x := left + 1; x < right; x++ {
					if overFlows {
						gm.terrain[y][x] = waterFlow
					} else {
						gm.terrain[y][x] = water
					}
				}
			}
		}
	}
}

func scanToXy(s scanLine) []xy {
	var result []xy

	for x := s.fromX; x <= s.toX; x++ {
		for y := s.fromY; y <= s.toY; y++ {
			result = append(result, xy{x, y})
		}
	}

	return result
}

func print(w io.Writer, ground groundMap) {
	for y := 0; y < ground.depth(); y++ {
		for x := 0; x < ground.width(); x++ {
			fmt.Fprintf(w, "%c", ground.terrain[y][x])
		}
		fmt.Fprint(w, "\n")
	}
}
