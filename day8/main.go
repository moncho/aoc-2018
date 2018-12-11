package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	id       string
	header   header
	children []*node
	metadata []int
}

func (n *node) newChild(node *node) {
	n.children = append(n.children, node)
}
func (n *node) metadataSum() int {
	sum := 0
	for _, m := range n.metadata {
		sum += m
	}

	for _, c := range n.children {
		sum += c.metadataSum()
	}

	return sum
}

type header struct {
	childCount    int
	metadataCount int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var n int
	var numbers []int
	for scanner.Scan() {
		_, err = fmt.Sscanf(scanner.Text(), "%d", &n)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	root, _ := treeNode(numbers)
	fmt.Printf("Metadata sum: %d\n", root.metadataSum())
}

func treeNode(numbers []int) (*node, []int) {
	l := len(numbers)
	if l == 0 {
		return nil, numbers
	}
	cc := numbers[0]
	mc := numbers[1]
	n := node{
		header: header{
			childCount:    cc,
			metadataCount: mc,
		},
	}
	numbers = numbers[2:]

	var child *node
	for i := 0; i < cc; i++ {
		child, numbers = treeNode(numbers)
		n.children = append(n.children, child)
	}

	n.metadata = numbers[:mc]

	return &n, numbers[mc:]
}
