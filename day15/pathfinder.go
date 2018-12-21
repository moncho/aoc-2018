package main

import (
	"github.com/beefsack/go-astar"
)

const (
	KindPlain = iota
	KindBlocker
	KindElf
	KindGoblin
	KindFrom
	KindTo
)

// RuneKinds map input runes to tile kinds.
var RuneKinds = map[rune]int{
	'.': KindPlain,
	'#': KindBlocker,
	'E': KindElf,
	'G': KindGoblin,
	'F': KindFrom,
	'T': KindTo,
}

// A Tile is a tile in a grid which implements Pather.
type Tile struct {
	kind       int
	x, y       int
	pathFinder *pathFinder
}

// PathNeighbors returns the neighbors of the tile, excluding blockers and
// tiles off the edge of the board.
func (t *Tile) PathNeighbors() []astar.Pather {
	var neighbors []astar.Pather
	for _, offset := range [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	} {
		x := t.x + offset[0]
		y := t.y + offset[1]
		w := t.pathFinder.w
		n := t.pathFinder.At(x, y)

		kind := RuneKinds[w[y][x]]
		if n != nil {
			kind = n.kind
		}
		if kind == KindPlain || kind == KindTo {
			if n == nil {
				n = &Tile{
					x:          x,
					y:          y,
					pathFinder: t.pathFinder,
					kind:       kind,
				}
				t.pathFinder.addTile(n)
			}
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

// PathNeighborCost returns the movement cost of the directly neighboring tile.
func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	return 1.0
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
// between non-adjacent nodes.
func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	absX := toT.x - t.x
	if absX < 0 {
		absX = -absX
	}
	absY := toT.y - t.y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

type pathFinder struct {
	tiles    map[xy]*Tile
	w        [][]rune
	from, to *Tile
}

func newPathFinder(w [][]rune, from, to xy) *pathFinder {
	pf := &pathFinder{
		w:     w,
		tiles: make(map[xy]*Tile),
	}
	fromT := &Tile{
		x:          from.x,
		y:          from.y,
		kind:       KindFrom,
		pathFinder: pf,
	}
	pf.from = fromT
	pf.addTile(fromT)

	toT := &Tile{
		x:          to.x,
		y:          to.y,
		kind:       KindTo,
		pathFinder: pf,
	}
	pf.to = toT
	pf.addTile(toT)

	return pf
}

func (pf *pathFinder) At(x, y int) (tile *Tile) {
	n, _ := pf.tiles[xy{
		x: x,
		y: y,
	}]
	return n
}

func (pf *pathFinder) addTile(tile *Tile) {
	pf.tiles[xy{
		tile.x, tile.y,
	}] = tile
}

func (pf *pathFinder) From() *Tile {
	return pf.from
}

func (pf *pathFinder) To() *Tile {
	return pf.to
}
