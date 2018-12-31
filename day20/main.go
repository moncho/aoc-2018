package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type route struct {
	path      []rune
	branches  []*route
	parent    *route
	skippable bool
}

func (r *route) longestRoute() int {
	l := len(r.path)
	longest := 0
	for _, b := range r.branches {
		if b.skippable {
			continue
		}
		bl := b.longestRoute()
		if bl > longest {
			longest = bl
		}
	}
	return l + longest
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	root := buildRoute(f)

	fmt.Printf("What is the largest number of doors you would be required to pass through to reach a room? %d\n", root.longestRoute())
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
			branch := &route{
				parent: cur,
			}
			cur.branches = append(cur.branches, branch)
			cur = branch
		case '|':
			branch := &route{
				parent: cur.parent,
			}
			cur.parent.branches = append(cur.parent.branches, branch)
			cur = branch

		case ')':
			if prev == '|' {
				cur.parent.branches = cur.parent.branches[:len(cur.parent.branches)-1]
				for _, b := range cur.parent.branches {
					b.skippable = true
				}
			}
			cur = cur.parent

		default:
			cur.path = append(cur.path, rune(r))
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
