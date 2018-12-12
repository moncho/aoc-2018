package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type point struct {
	x, y int
	xvel int
	yvel int
}

func (p *point) tick(time int) {
	p.x = p.x + p.xvel*time
	p.y = p.y + p.yvel*time
}

func (p *point) visible() bool {
	return p.x > 0 && p.y > 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var points []*point

	for scanner.Scan() {
		p := point{}
		_, err = fmt.Sscanf(scanner.Text(), "position=<%d,  %d> velocity=<%d,  %d>",
			&p.x, &p.y, &p.xvel, &p.yvel)
		if err != nil {
			panic(err)
		}
		points = append(points, &p)

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	future, time := fastForward(points)
	draw(future)
	fmt.Printf("Took %d seconds to reach the future.\n", time)

}

func fastForward(points []*point) ([]*point, int) {
	keepMoving := true
	i := 0
	//TODO figure out a way to stop going forward without using magic
	//numbers
	//15 was the result of manually checking  and tweaking values until
	// the result was good.
	magicNumber := 15

	//This get us closer by is not enough
	for keepMoving {
		i++
		keepMoving = false
		for _, p := range points {
			p.tick(magicNumber)
			if !p.visible() {
				keepMoving = true
			}
		}
	}
	//15 more ticks are needed
	for _, p := range points {
		p.tick(magicNumber)
	}

	return points, i*magicNumber + magicNumber
}

func draw(points []*point) {
	minX, minY, maxX, maxY := minMax(points)

	message := make([][]rune, maxY-minY+1)
	for y := range message {
		message[y] = make([]rune, maxX-minX+1)
		for x := range message[y] {
			message[y][x] = '.'
		}
	}

	var normY, normX int
	for _, p := range points {
		normY = (p.y - minY)
		normX = (p.x - minX)
		message[normY][normX] = '#'
	}
	for _, line := range message {
		fmt.Println(string(line))
	}
}

func minMax(points []*point) (int, int, int, int) {

	maxX, maxY := 0, 0
	minX, minY := math.MaxInt32, math.MaxInt32

	for _, p := range points {

		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}

	}

	return minX, minY, maxX, maxY

}
