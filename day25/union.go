package main

type weightedQuickUnion struct {
	roots   []int
	weights []int
	sets    int
}

func NewWeightedQuickUnion(elements int) *weightedQuickUnion {
	roots := make([]int, elements)
	weights := make([]int, elements)

	for i := range roots {
		roots[i] = i
		weights[i] = 1
	}
	return &weightedQuickUnion{
		roots:   roots,
		weights: weights,
		sets:    elements,
	}
}

func (qu *weightedQuickUnion) root(a int) int {
	root := qu.roots[a]
	for root != qu.roots[root] {
		//flatten tree
		qu.roots[root] = qu.roots[qu.roots[root]]
		root = qu.roots[root]
	}
	return root
}

func (qu *weightedQuickUnion) Union(a, b int) {
	rootB := qu.root(b)
	rootA := qu.root(a)
	if rootB == rootA {
		return
	}
	if qu.weights[rootA] > qu.weights[rootB] {
		qu.weights[rootA] += qu.weights[rootB]
		qu.roots[rootB] = rootA
	} else {
		qu.weights[rootB] += qu.weights[rootA]
		qu.roots[rootA] = rootB
	}
	qu.sets--
}

func (qu *weightedQuickUnion) Connected(a, b int) bool {
	return qu.root(a) == qu.root(b)
}
