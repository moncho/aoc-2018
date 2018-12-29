package main

import (
	"bufio"
	"fmt"
	"os"
)

type registers []int

type operation struct {
	opcode         string
	inputA, inputB int
	output         int
}
type program struct {
	instructionPointer int
	instructions       []operation
}

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)

	s.Scan()
	var pointer int
	fmt.Sscanf(s.Text(), "#ip %d", &pointer)

	var operations []operation
	for s.Scan() {
		var ins operation
		fmt.Sscanf(s.Text(), "%s %d %d %d", &ins.opcode, &ins.inputA, &ins.inputB, &ins.output)
		operations = append(operations, ins)

	}
	if s.Err() != nil {
		panic(err)
	}
	r := run(program{pointer, operations})
	fmt.Printf("Register values when the background process halts: %v", r)
}

func run(p program) registers {
	reg := make([]int, 6)
	instructionPointer := 0
	for instructionPointer < len(p.instructions) {
		reg[p.instructionPointer] = instructionPointer
		op := p.instructions[instructionPointer]
		ins := instructionSet[op.opcode]
		reg = ins(op)(reg)
		instructionPointer = reg[p.instructionPointer]
		instructionPointer++
	}
	return reg
}
