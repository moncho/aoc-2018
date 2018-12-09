package main

import (
	"sort"
)

// crossProduct is the 2D cross product of OA and OB vectors.
// Returns a positive value, if OAB makes a counter-clockwise turn,
// negative for clockwise turn, and zero if the points are collinear.
func crossProduct(o, a, b coord) int {
	return (a.x-o.x)*(b.y-o.y) - (a.y-o.y)*(b.x-o.x)
}

//convexHull computes the convex hull of the given slice of coords.
//Based on https://www.wikibooks.org/wiki/Algorithm_Implementation/Geometry/Convex_hull/Monotone_chain
func convexHull(coords []coord) []coord {
	n := len(coords)
	if n <= 3 {
		return coords
	}
	// sort our coords by x and if equal by y
	sort.Slice(coords, func(i, j int) bool {
		if coords[i].x == coords[j].x {
			return coords[i].y < coords[j].y
		}
		return coords[i].x < coords[j].x
	})

	result := make([]coord, n*2)
	k := 0
	// Build lower hull
	for i := 0; i < n; i++ {
		for k >= 2 && crossProduct(result[k-2], result[k-1], coords[i]) <= 0 {
			k--
		}
		result[k] = coords[i]
		k++
	}

	// Build upper hull
	for i, t := n-2, k+1; i >= 0; i-- {
		for k >= t && crossProduct(result[k-2], result[k-1], coords[i]) <= 0 {
			k--
		}
		result[k] = coords[i]
		k++
	}

	//	remove k - 1 which is a duplicate
	return result[:k-1]
}
