package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

const (
	plant        = '#'
	noplant      = '.'
	fiftyBillion = 50000000000
)

type pots struct {
	pots          []bool
	zeroPotOffset int
}

func (p *pots) plantAt(n int) bool {
	if n < 0 || n >= len(p.pots) {
		return false
	}
	return p.pots[n]
}

func (p *pots) len() int {
	l := len(p.pots)
	if p.zeroPotOffset > 0 {
		l = l + p.zeroPotOffset
	}
	return l
}

func (p *pots) grow(notes []note) {
	var newGen []bool

	leftmostPot := 0
	firstLeftFound := false
	rightmostPot := 0
	offset := 3
	for i := -offset; i <= p.len()+offset; i++ {
		pattern := []bool{p.plantAt(i - 2), p.plantAt(i - 1), p.plantAt(i), p.plantAt(i + 1), p.plantAt(i + 2)}
		grows := matchInNotes(notes, pattern)
		if grows {
			if !firstLeftFound {
				leftmostPot = i
				firstLeftFound = true
			}
			rightmostPot = i
		}
		newGen = append(newGen, grows)
	}
	p.zeroPotOffset += leftmostPot
	p.pots = newGen[offset+leftmostPot : offset+rightmostPot+1]
}
func (p *pots) potNumSum() int {
	sum := 0
	for i, pot := range p.pots {
		if pot {
			sum += i + p.zeroPotOffset
		}
	}
	return sum
}

type note struct {
	state []bool
	plant bool
}

func matchInNotes(notes []note, pattern []bool) bool {
	for _, n := range notes {
		if reflect.DeepEqual(n.state, pattern) {
			return n.plant
		}
	}
	return false
}

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()

	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	s.Scan()
	initialState := parse(strings.TrimLeft(s.Text(), "initial state: "))
	p := &pots{
		pots:          initialState,
		zeroPotOffset: 0,
	}

	var notes []note
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		notes = append(notes,
			note{
				state: parse(line[:5]),
				plant: parse(line[9:])[0],
			})

	}
	if s.Err() != nil {
		panic(err)
	}
	simulate(20, p, notes)

	fmt.Printf("After 20 generations, what is the sum of the numbers of all pots which contain a plant? %d\n", p.potNumSum())

	p = &pots{
		pots:          initialState,
		zeroPotOffset: 0,
	}
	gen, growthPerGen := simulate(fiftyBillion, p, notes)
	sum := p.potNumSum() + ((fiftyBillion - gen) * growthPerGen)
	fmt.Printf("After fifty billion generations, what is the sum of the numbers of all pots which contain a plant? %d\n", sum)

}

//simulate grows the given pots until the given number of generations
//is reached or the generation growth stabilizes
func simulate(n int, pots *pots, notes []note) (int, int) {
	var lastSum int
	var diff int
	for i := 0; i < n; i++ {
		pots.grow(notes)
		sum := pots.potNumSum()
		if sum-lastSum == diff {
			return i + 1, diff
		}
		diff = sum - lastSum
		lastSum = sum

	}
	return n, diff
}

func printPots(pots []bool) {
	var s strings.Builder

	for _, p := range pots {
		if p {
			s.WriteRune(plant)
		} else {
			s.WriteRune(noplant)
		}
	}
	fmt.Println(s.String())

}
func parse(s string) []bool {
	var plants []bool
	for _, c := range s {
		if c == plant {
			plants = append(plants, true)
		} else {
			plants = append(plants, false)
		}
	}
	return plants
}
