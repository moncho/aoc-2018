package main

import (
	"bufio"
	"io"
)

type route struct {
	directionBlocks map[int][]rune
	branchBlocks    map[int][]*route
	parent          *route
	skippable       bool
	branchCount     int
}

func (r *route) direction(dir rune) {
	if r.directionBlocks == nil {
		r.directionBlocks = make(map[int][]rune)
	}

	if b, ok := r.directionBlocks[r.branchCount]; ok {
		b = append(b, dir)
		r.directionBlocks[r.branchCount] = b
	} else {
		r.directionBlocks[r.branchCount] = []rune{dir}
	}
}
func (r *route) addBranch(branch *route) {
	if r.branchBlocks == nil {
		r.branchBlocks = make(map[int][]*route)
	}

	if b, ok := r.branchBlocks[r.branchCount]; ok {
		b = append(b, branch)
		r.branchBlocks[r.branchCount] = b
	} else {
		r.branchBlocks[r.branchCount] = []*route{branch}
	}
}
func (r *route) closeBranch() {
	l := len(r.branchBlocks[r.branchCount])
	//If the last route of the current branch block can
	//be skipped then all routes of this block can be skipped
	if r.branchBlocks[r.branchCount][l-1].skippable {
		for _, b := range r.branchBlocks[r.branchCount] {
			b.skippable = true
		}
	}
	r.branchCount++
}
func (r *route) longestRoute() int {
	l := 0
	for _, block := range r.directionBlocks {
		l += len(block)
	}
	longest := 0
	for _, block := range r.branchBlocks {
		for _, b := range block {
			if b.skippable {
				continue
			}
			bl := b.longestRoute()
			if bl > longest {
				longest = bl
			}
		}
	}
	return l + longest
}

//allRoutes returns all the distances starting from this route
//TODO implementation is not correct, possibly because it is
//revisiting distances
func (r *route) allDistances() []int {
	routes := []int{0}
	curMax := 0
	for blockIndex := 0; blockIndex < len(r.directionBlocks); blockIndex++ {
		block := r.directionBlocks[blockIndex]
		for i := 1; i <= len(block); i++ {
			routes = append(routes, i+curMax)
		}
		curMax += len(block)
		for _, branch := range r.branchBlocks[blockIndex] {
			if branch.skippable {
				continue
			}
			for _, br := range branch.allDistances() {
				if br == 0 {
					continue
				}
				routes = append(routes, br+curMax)
			}
		}
	}

	return routes
}

func buildRoute(r io.Reader) *route {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)

	cur := &route{}
	var prev rune
	for s.Scan() {
		r := s.Text()[0]
		switch r {
		case '^', '$', '\n':
			break
		case '(':
			r := &route{
				parent: cur,
			}
			cur.addBranch(r)
			cur = r
		case '|':
			r := &route{
				parent: cur.parent,
			}
			cur.parent.addBranch(r)
			cur = r

		case ')':
			if prev == '|' {
				cur.skippable = true
			}
			cur = cur.parent
			cur.closeBranch()
		default:
			cur.direction(rune(r))
		}
		prev = rune(r)
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	root := cur
	for root.parent != nil {
		root = root.parent
	}
	return root
}
