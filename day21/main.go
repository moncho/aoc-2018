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
	reg := make([]int, 6)
	//Guessed by running the program an observing the state of register 5
	//on the first comparison with register 0. run stops the program on
	//this comparison
	reg[0] = 15690445
	r := run(program{pointer, operations}, reg)
	fmt.Printf("Register values when the background process halts: %v\n", r)

	reg2 := []int{0, 0, 0, 0, 0, 0}

	reg2 = part2Run(program{pointer, operations}, reg2)
	fmt.Printf("Register values when the background process halts: %v\n", reg2)

}

func run(p program, reg registers) registers {
	instructionPointer := 0
	for instructionPointer < len(p.instructions) {
		reg[p.instructionPointer] = instructionPointer
		op := p.instructions[instructionPointer]
		reg = instructionSet[op.opcode](op)(reg)
		instructionPointer = reg[p.instructionPointer] + 1
		//This opcode is the only operation from input
		//that uses register 0
		if op.opcode == "eqrr" {
			break
		}
	}
	return reg
}

func part2Run(p program, reg registers) registers {
	instructionPointer := 0
	seen := make(map[int]bool)
	last := 0
	for instructionPointer < len(p.instructions) {
		reg[p.instructionPointer] = instructionPointer
		op := p.instructions[instructionPointer]
		reg = instructionSet[op.opcode](op)(reg)
		instructionPointer = reg[p.instructionPointer] + 1
		//This opcode is the only operation from input
		//that uses register 0
		if op.opcode == "eqrr" {
			if seen[reg[op.inputA]] {
				fmt.Printf("Last value on reg used on the comparision, before a duplicate: %d\n", last)
				break
			}
			seen[reg[op.inputA]] = true
			last = reg[op.inputA]
		}

	}
	return reg
}
