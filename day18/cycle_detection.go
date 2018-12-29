package main

type f func([][]rune) ([][]rune, int)

//brent is an implementation of Brent cycle detection algorithm
//as described https://en.wikipedia.org/wiki/Cycle_detection#Brent's_algorithm
func brent(f f, x0 [][]rune) (int, int) {
	// main phase: search successive powers of two
	power, lam := 1, 1
	landscape := copyLandscape(x0)
	tortoise, hare := 0, 0
	landscape, hare = f(landscape)
	for tortoise != hare {
		if power == lam { //time to start a new power of two?
			tortoise = hare
			power *= 2
			lam = 0
		}
		landscape, hare = f(landscape)
		lam += 1
	}
	// Find the position of the first repetition of length λ
	mu := 0
	tortoise, hare = 0, 0
	landscape = copyLandscape(x0)
	for i := 0; i < lam; i++ {
		// range(lam) produces a list with the values 0, 1, ... , lam-1
		landscape, hare = f(landscape)
	}
	// The distance between the hare and tortoise is now λ.
	tortoiseLandscape := copyLandscape(x0)
	//# Next, the hare and tortoise move at same speed until they agree
	for tortoise != hare {
		tortoiseLandscape, tortoise = f(tortoiseLandscape)
		landscape, hare = f(landscape)
		mu += 1
	}
	return lam, mu
}
