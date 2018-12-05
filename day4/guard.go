package main

type guard struct {
	id           string
	sleepPattern map[int]int
}

func (g *guard) totalSleeping() int {
	totalSleeping := 0
	for _, sleeps := range g.sleepPattern {
		totalSleeping += sleeps
	}
	return totalSleeping
}
func (g *guard) sleepiestMinute() int {
	sleepiestMinute := 0
	mostSleeps := 0
	for minute, sleeps := range g.sleepPattern {
		if sleeps > mostSleeps || (sleeps == mostSleeps && minute < sleepiestMinute) {
			sleepiestMinute = minute
			mostSleeps = sleeps
		}
	}
	return sleepiestMinute
}
