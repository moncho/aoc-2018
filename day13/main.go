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
	working   bool
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

func runSimulation(grid [][]rune, carts []*cart) ([]*cart, []*cart) {
	var broken []*cart
	wc := workingCarts(carts)
	for len(wc) > 1 {
		sort.Slice(wc, func(i, j int) bool {
			if wc[i].y == wc[j].y {
				return wc[i].x < wc[j].x
			}
			return wc[i].y < wc[j].y
		})
		visited := make(map[string]int)

		for i, cart := range wc {
			id := cart.id()
			if other, ok := visited[id]; ok {
				if wc[other].working {
					broken = append(broken, wc[other])
					wc[other].working = false
				}
				broken = append(broken, cart)
				cart.working = false
				continue
			}
			tick(cart, grid)
			id = cart.id()
			if other, ok := visited[id]; ok {
				if wc[other].working {
					broken = append(broken, wc[other])
					wc[other].working = false
				}
				broken = append(broken, cart)
				cart.working = false
				continue
			}
			visited[id] = i

		}
		wc = workingCarts(wc)

	}
	return broken, wc
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
					working:   true,
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

func workingCarts(carts []*cart) []*cart {
	result := carts[:0]
	for _, c := range carts {
		if c.working {
			result = append(result, c)
		}
	}
	return result
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

	broken, left := runSimulation(grid, carts)
	fmt.Printf("Broken carts: %d\n", len(broken))
	fmt.Printf("First collision at: %d,%d\n", broken[0].x, broken[0].y)

	fmt.Printf("Carts left: %d\n", len(left))
	fmt.Printf("Last cart standing at: %d,%d\n", left[0].x, left[0].y)

	for _, c := range carts {
		if isCart(grid[c.y][c.x]) {
			grid[c.y][c.x] = 'X'
		} else {
			grid[c.y][c.x] = c.direction
		}
	}
	//	print(grid)
}
