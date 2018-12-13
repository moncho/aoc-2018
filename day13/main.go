package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

const (
	faceUp    = '^' //4
	faceDown  = 'v' //4
	faceLeft  = '<' //2
	faceRight = '>' //7

	horizontal    = '-'
	vertical      = '|'
	forwardSlah   = '/'
	backwardSlash = '\\'
	crossPath     = '+'
)

type cart struct {
	x, y      int
	nextCross int
	direction rune
}

func (c *cart) move() {
	switch c.direction {
	case faceUp:
		c.y--
	case faceDown:
		c.y++
	case faceLeft:
		c.x--
	case faceRight:
		c.x++
	}
}
func (c *cart) leftTurn() {
	switch c.direction {
	case faceUp:
		c.direction = faceLeft
	case faceDown:
		c.direction = faceRight
	case faceLeft:
		c.direction = faceDown
	case faceRight:
		c.direction = faceUp
	}

}
func (c *cart) rightTurn() {
	switch c.direction {
	case faceUp:
		c.direction = faceRight
	case faceDown:
		c.direction = faceLeft
	case faceLeft:
		c.direction = faceUp
	case faceRight:
		c.direction = faceDown
	}

}
func (c *cart) cross() {
	switch c.nextCross {
	case 0:
		c.nextCross = 1
		c.leftTurn()
	case 1:
		c.nextCross = 2
	case 2:
		c.nextCross = 0
		c.rightTurn()
	}
}

func (c *cart) id() string {
	return strconv.Itoa(c.x) + "," + strconv.Itoa(c.y)
}

func runSimulation(grid [][]rune, carts []*cart) (int, int) {
	collision := false
	collisionX, collisionY := -1, -1
	for !collision {
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y == carts[j].y {
				return carts[i].x < carts[j].x
			}
			return carts[i].y < carts[j].y
		})
		visited := make(map[string]bool)
		for _, cart := range carts {
			id := cart.id()
			if visited[id] {
				collision = true
				collisionX = cart.x
				collisionY = cart.y
				break
			}
			visited[id] = true
			tick(cart, grid)
			id = cart.id()
			if visited[id] {
				collision = true
				collisionX = cart.x
				collisionY = cart.y
				break
			}
			visited[id] = true
		}
	}
	return collisionX, collisionY
}
func newGrid(r io.Reader) ([][]rune, []*cart) {
	scanner := bufio.NewScanner(r)
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
					x:         pos,
					y:         lc,
					direction: r,
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
	return grid, carts
}

func print(runes [][]rune) {
	for _, line := range runes {
		fmt.Printf("%s\n", string(line))
	}
}

func isCart(r rune) bool {
	return r == faceUp || r == faceDown || r == faceLeft || r == faceRight
}

func tick(c *cart, grid [][]rune) {
	track := grid[c.y][c.x]
	switch track {
	case horizontal, vertical:
	case forwardSlah:
		if c.direction == faceLeft || c.direction == faceRight {
			c.leftTurn()
		} else {
			c.rightTurn()
		}
	case backwardSlash:
		if c.direction == faceLeft || c.direction == faceRight {
			c.rightTurn()
		} else {
			c.leftTurn()
		}
	case crossPath:
		c.cross()
	default:
		panic(fmt.Sprintf("Unexpected coords: %d, %d", c.x, c.y))
	}
	c.move()
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	grid, carts := newGrid(file)
	defer file.Close()

	x, y := runSimulation(grid, carts)
	fmt.Printf("Collision at %d,%d\n", x, y)

	for _, c := range carts {
		if isCart(grid[c.y][c.x]) {
			grid[c.y][c.x] = 'X'
		} else {
			grid[c.y][c.x] = c.direction
		}
	}
	//	print(grid)
}
