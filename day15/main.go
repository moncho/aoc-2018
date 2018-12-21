package main

import (
	"bufio"
	"fmt"
	"github.com/beefsack/go-astar"
	"io"
	"math"
	"os"
	"sort"
	"strings"
)

var offSets = [][]int{
	{0, -1},
	{-1, 0},
	{1, 0},
	{0, 1},
}

const (
	wall        = '#'
	open        = '.'
	attackPower = 3
	hitpoints   = 200
)

const (
	goblin = 'G'
	elf    = 'E'
)

type game struct {
	battleMap [][]rune
	elfs      []*unit
	goblins   []*unit
}

func (g *game) isOpen(x, y int) bool {
	return g.battleMap[y][x] == open
}
func (g *game) dimensions() (int, int) {
	if len(g.battleMap) == 0 {
		return 0, 0
	}
	return len(g.battleMap), len(g.battleMap[0])
}
func (g *game) unitOn(x, y int) *unit {
	for _, e := range g.elfs {
		if e.x == x && e.y == y {
			return e
		}
	}
	for _, g := range g.goblins {
		if g.x == x && g.y == y {
			return g
		}
	}
	return nil
}

func (g *game) adjacents(u *unit) []xy {
	var result []xy
	for _, offset := range offSets {
		x := u.x + offset[0]
		y := u.y + offset[1]
		if g.battleMap[y][x] == open {
			result = append(result, xy{x, y})
		}
	}

	return result
}
func (g *game) lockTargetAndMove(u *unit) {
	uxy := xy{u.x, u.y}
	var enemies []*unit
	if u.enemy() == elf {
		enemies = g.elfs
	} else {
		enemies = g.goblins
	}
	sort.Slice(enemies, u.distanceSort(enemies))
	var nearest *unit
	minDistance := math.MaxFloat64
	var shortestPath []astar.Pather
	for _, enemy := range enemies {
		if enemy.alive() {
			xys := g.adjacents(enemy)
			for _, xy := range xys {
				pf := newPathFinder(g.battleMap, uxy, xy)
				if p, distance, found := astar.Path(pf.From(), pf.To()); found && distance < minDistance {
					minDistance = distance
					nearest = enemy
					shortestPath = p
				}
			}
		}
	}
	if nearest != nil {
		//First element on the shortestPath is the destination
		//Last element on the shortestPath is the origin
		toT := shortestPath[len(shortestPath)-2].(*Tile)
		g.move(u, toT.x, toT.y)
	}

}

func (g *game) move(u *unit, x, y int) {
	if g.battleMap[u.y][u.x] != u.unitType {
		panic(
			fmt.Sprintf("trying to move %q at (%d, %d) but there is a %q there ",
				u.unitType, u.x, u.y, g.battleMap[y][x]))
	}
	if g.battleMap[y][x] != open {
		panic(
			fmt.Sprintf("trying to move to (%d, %d) but there is a %q there ",
				x, y, g.battleMap[y][x]))
	}
	g.battleMap[u.y][u.x] = open
	u.x = x
	u.y = y
	g.battleMap[u.y][u.x] = u.unitType

}

func (g *game) canAttack(u *unit) (*unit, bool) {
	battleMap := g.battleMap
	enemyType := u.enemy()
	var enemy *unit
	for _, offset := range offSets {
		x := u.x + offset[0]
		y := u.y + offset[1]
		if battleMap[y][x] == enemyType {
			e := g.unitOn(x, y)
			if e.alive() && (enemy == nil || enemy.hp > e.hp) {
				enemy = e
			}
		}
	}

	return enemy, enemy != nil
}
func (g *game) cleanBattleField() {
	var elfs []*unit
	for _, e := range g.elfs {
		if e.alive() {
			elfs = append(elfs, e)
		} else {
			g.battleMap[e.y][e.x] = open
		}
	}
	g.elfs = elfs

	var goblins []*unit

	for _, gob := range g.goblins {
		if gob.alive() {
			goblins = append(goblins, gob)
		} else {
			g.battleMap[gob.y][gob.x] = open
		}
	}
	g.goblins = goblins

}

func (g *game) over() bool {
	return len(g.elfs) == 0 || len(g.goblins) == 0
}

func (g *game) score() int {
	score := 0
	for _, e := range g.elfs {
		score += e.hp
	}

	for _, gob := range g.goblins {
		score += gob.hp
	}

	return score
}

type xy struct {
	x, y int
}
type unit struct {
	x, y     int
	hp       int
	power    int
	unitType rune
}

func (u *unit) alive() bool {
	return u.hp > 0
}

func (u *unit) enemy() rune {
	if u.unitType == goblin {
		return elf
	}
	return goblin
}

func (u *unit) attack(enemy *unit) {
	enemy.hp -= u.power
}

func (u *unit) distanceTo(other *unit) int {
	return abs(u.x-other.x) + abs(u.y-other.y)
}

func (u *unit) compareTo(that *unit) int {
	if u.y == that.y {
		return u.x - that.x
	}
	return u.y - that.y

}
func (u *unit) distanceSort(units []*unit) func(i, j int) bool {
	return func(i, j int) bool {
		iSlope := u.distanceTo(units[i])
		jSlope := u.distanceTo(units[j])
		if iSlope == jSlope {
			return units[i].compareTo(units[j]) < 0
		}
		return iSlope < jSlope
	}
}

func (u *unit) string() string {
	return fmt.Sprintf("%c(%d)", u.unitType, u.hp)
}
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	game := newGame(f)

	print(game)
	moves := play(game)
	fmt.Printf("Final map\n")
	print(game)
	fmt.Printf("Final score: %dx%d=%d\n", game.score(), moves, game.score()*moves)

}
func newGame(r io.Reader) *game {
	s := bufio.NewScanner(r)
	var bm [][]rune
	var elfs []*unit
	var goblins []*unit
	row := 0
	for s.Scan() {
		s := s.Text()
		var line []rune
		for col, r := range s {
			switch r {
			case goblin:
				goblins = append(goblins, &unit{
					x:        col,
					y:        row,
					hp:       hitpoints,
					power:    attackPower,
					unitType: goblin,
				})
			case elf:
				elfs = append(elfs, &unit{
					x:        col,
					y:        row,
					hp:       hitpoints,
					power:    attackPower,
					unitType: elf,
				})
			}
			line = append(line, r)
		}
		bm = append(bm, line)
		row++
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return &game{
		battleMap: bm,
		elfs:      elfs,
		goblins:   goblins,
	}
}
func play(game *game) int {
	rounds := 0
	for {
		var units []*unit
		units = append(units, game.elfs...)
		units = append(units, game.goblins...)
		sortUnits(units)
		for _, u := range units {
			if u.alive() {
				if enemy, ok := game.canAttack(u); ok {
					u.attack(enemy)
					continue
				}
				game.lockTargetAndMove(u)
				if enemy, ok := game.canAttack(u); ok {
					u.attack(enemy)
				}
			}
		}
		game.cleanBattleField()
		if game.over() {
			break
		}
		rounds++
		fmt.Printf("Round %d\n", rounds)
		print(game)

	}
	return rounds
}

func print(game *game) {
	bm := game.battleMap
	for y, line := range bm {
		var hps []string
		for x, r := range line {
			if r == elf || r == goblin {
				u := game.unitOn(x, y)
				hps = append(hps, u.string())
			}
		}
		fmt.Printf("%s   %s\n", string(line), strings.Join(hps, ", "))
	}
}

func sortUnits(units []*unit) {
	sort.Slice(units, func(i, j int) bool {
		if units[i].y == units[j].y {
			return units[i].x < units[j].x
		}
		return units[i].y < units[j].y
	})
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
