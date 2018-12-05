package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	pattern = "2006-01-02 15:04"
)

type recordEntry struct {
	guardID    string
	timestamp  time.Time
	minute     int
	recordType recordType
}

type recordType int

const (
	NewShift recordType = iota
	Sleep
	Awake
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var entries []recordEntry
	for scanner.Scan() {
		entries = append(entries, parse(scanner.Text()))
	}

	guardLog := buildGuardLog(entries)
	sleepiestGuard := mostSleepingMinutesGuard(guardLog)

	id, _ := strconv.Atoi(sleepiestGuard.id)
	fmt.Printf("GuardId: %s, minute: %d, sol1: %d\n",
		sleepiestGuard.id,
		sleepiestGuard.sleepiestMinute(),
		id*sleepiestGuard.sleepiestMinute())
	sleepiestMinuteGuard := sleepiestMinuteGuard(guardLog)
	id, _ = strconv.Atoi(sleepiestMinuteGuard.id)
	fmt.Printf("GuardId: %s, minute: %d, sol2: %d\n",
		sleepiestMinuteGuard.id,
		sleepiestMinuteGuard.sleepiestMinute(),
		id*sleepiestMinuteGuard.sleepiestMinute())
}

func buildGuardLog(entries []recordEntry) []*guard {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].timestamp.Before(entries[j].timestamp)
	})
	guardLog := make(map[string]*guard)
	var onDuty *guard
	var startSleep int
	for _, entry := range entries {
		switch entry.recordType {
		case NewShift:
			g, ok := guardLog[entry.guardID]
			if !ok {
				g = &guard{
					id:           entry.guardID,
					sleepPattern: make(map[int]int),
				}
				guardLog[g.id] = g
			}
			onDuty = g
		case Sleep:
			if onDuty != nil {
				startSleep = entry.minute
			}
		case Awake:
			if onDuty != nil {
				for i := startSleep; i < entry.minute; i++ {
					onDuty.sleepPattern[i]++
				}
			}
		}
	}
	result := make([]*guard, 0, len(guardLog))

	for _, guard := range guardLog {
		result = append(result, guard)
	}
	return result
}
func mostSleepingMinutesGuard(guardLog []*guard) *guard {
	var candidate *guard
	mostSleepingMinutes := 0
	for _, guard := range guardLog {
		total := guard.totalSleeping()
		if total > mostSleepingMinutes {
			candidate = guard
			mostSleepingMinutes = total
		}
	}

	return candidate
}

func sleepiestMinuteGuard(guardLog []*guard) *guard {
	var candidate *guard
	sleepiestMinute := 0
	for _, guard := range guardLog {
		sleepiest := guard.sleepPattern[guard.sleepiestMinute()]
		if sleepiest > sleepiestMinute {
			candidate = guard
			sleepiestMinute = sleepiest
		}
	}

	return candidate
}

func parse(logEntry string) recordEntry {
	timestamp := parseTimeStamp(logEntry)

	re := recordEntry{
		timestamp: timestamp,
		minute:    timestamp.Minute(),
	}

	if strings.Contains(logEntry, "begins") {
		re.recordType = NewShift
		i := strings.IndexRune(logEntry, '#')
		re.guardID = strings.Fields(logEntry[i:])[0][1:]

	} else if strings.Contains(logEntry, "asleep") {
		re.recordType = Sleep
	} else {
		re.recordType = Awake
	}
	return re
}

func parseTimeStamp(logEntry string) time.Time {
	time, err := time.Parse(pattern, logEntry[1:17])
	if err != nil {
		panic(err)
	}
	return time
}
