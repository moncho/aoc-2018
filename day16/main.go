package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type registers []int

type operation struct {
	opcode         int
	inputA, inputB int
	output         int
}

type sample struct {
	before    registers
	operation operation
	after     registers
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	var samples []sample
	var operations []operation

	var curi *sample
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "Before") {
			r := make([]int, 4)
			fmt.Sscanf(line, "Before: [%d, %d, %d, %d]", &r[0], &r[1], &r[2], &r[3])
			curi = &sample{
				before: r,
			}
			continue
		} else if strings.HasPrefix(line, "After") {
			r := make([]int, 4)
			fmt.Sscanf(line, "After: [%d, %d, %d, %d]", &r[0], &r[1], &r[2], &r[3])
			curi.after = r
			samples = append(samples, *curi)
			curi = nil
			continue
		}
		var op operation
		fmt.Sscanf(line, "%d %d %d %d", &op.opcode, &op.inputA, &op.inputB, &op.output)
		if curi != nil {
			curi.operation = op
		} else {
			operations = append(operations, op)
		}
	}
	if s.Err() != nil {
		panic(err)
	}

	count := countThreeOpCodes(samples)

	fmt.Printf("Samples behaving like three opcodes: %d\n", count)
}

func countThreeOpCodes(samples []sample) int {
	count := 0

	for _, s := range samples {
		c := 0
		for _, i := range instructionSet {
			before := make([]int, 4)
			copy(before, s.before)
			got := i(s.operation)(before)
			if reflect.DeepEqual(got, s.after) {
				c++
			}
		}
		if c >= 3 {
			count++
		}
	}

	return count
}
