package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type xyz struct {
	x, y, z int
}

func (me xyz) distanceTo(xyz xyz) int {
	return abs(me.x-xyz.x) + abs(me.y-xyz.y) + abs(me.z-xyz.z)
}

type nanobot struct {
	xyz xyz
	r   int
}

func (n nanobot) nanoInRange(nano nanobot) bool {
	return n.inRange(nano.xyz)
}
func (n nanobot) inRange(xyz xyz) bool {
	return n.xyz.distanceTo(xyz) <= n.r
}

func (n nanobot) distanceToNano(nano nanobot) int {
	return n.xyz.distanceTo(nano.xyz)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var nanos []nanobot
	for s.Scan() {
		nano := nanobot{
			xyz: xyz{},
		}
		fmt.Sscanf(s.Text(), "pos=<%d,%d,%d>, r=%d", &nano.xyz.x, &nano.xyz.y, &nano.xyz.z, &nano.r)
		nanos = append(nanos, nano)

	}

	sort.Slice(nanos, func(i, j int) bool {
		return nanos[i].r < nanos[j].r
	})

	strongesNano := nanos[len(nanos)-1]
	fmt.Printf("Nanos: %d\n", len(nanos))
	fmt.Printf("Strongest signal: %d\n", strongesNano.r)

	inRange := filterInRange(strongesNano, nanos)

	fmt.Printf("Nanobots in range of strongest signal: %d\n", len(inRange))

	best := mostInRange(nanos)
	fmt.Printf("Distance between %v and (0,0,0)?: %d\n", best, best.distanceTo(xyz{0, 0, 0}))
}

func mostInRange(nanos []nanobot) xyz {
	var cur, min, max xyz
	//The zoom level can be adjusted, current value works for
	//the problem input
	zoomLevel := 1 << 32
	zoomedBots := make([]nanobot, len(nanos))

	for {
		for i, n := range nanos {
			zc := nanobot{
				xyz{
					n.xyz.x / zoomLevel, n.xyz.y / zoomLevel, n.xyz.z / zoomLevel,
				}, n.r / zoomLevel}
			zoomedBots[i] = zc
		}

		best := struct {
			pos   xyz
			count int
		}{}

		for cur.x = min.x; cur.x <= max.x; cur.x++ {
			for cur.y = min.y; cur.y <= max.y; cur.y++ {
				for cur.z = min.z; cur.z <= max.z; cur.z++ {
					count := 0
					for _, nano := range zoomedBots {
						if nano.inRange(cur) {
							count++
						}
					}

					if count < best.count {
						continue
					}

					best.pos, best.count = cur, count
				}
			}
		}

		min.x, min.y, min.z = (best.pos.x-1)<<1, (best.pos.y-1)<<1, (best.pos.z-1)<<1
		max.x, max.y, max.z = (best.pos.x+1)<<1, (best.pos.y+1)<<1, (best.pos.z+1)<<1
		zoomLevel >>= 1

		if zoomLevel == 0 {
			return best.pos
		}
	}
}

func filterInRange(nano nanobot, nanos []nanobot) []nanobot {

	var inRange []nanobot
	for _, n := range nanos {
		if nano.nanoInRange(n) {
			inRange = append(inRange, n)
		}
	}

	return inRange
}
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
