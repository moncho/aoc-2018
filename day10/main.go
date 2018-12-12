package main

import (
	"bufio"
	"fmt"
	"image"
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
	future := fastForward(points)
	draw(future)
}

func fastForward(points []*point) []*point {
	keepMoving := true
	for keepMoving {
		keepMoving = false
		for _, p := range points {
			p.tick(100)
			if !p.visible() {
				keepMoving = true
			}
		}
	}
	return points
}

func draw(points []*point) {
	b := bounds(points)
	s := b.Size()
	arr := make([][]rune, s.Y)
	for iy := range arr {
		arr[iy] = make([]rune, s.X)
		for ix := range arr[iy] {
			arr[iy][ix] = '.'
		}
	}

	for _, p := range points {
		arr[p.y][p.x] = '#'
	}
	for _, line := range arr {
		fmt.Println(string(line))
	}
}

func bounds(points []*point) image.Rectangle {

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

	return image.Rect(minX, minY, maxX, maxY)

}
