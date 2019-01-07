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
	//Guessed by s
	reg[0] = 15690445
	r := run(program{pointer, operations}, reg)
	fmt.Printf("Register values when the background process halts: %v\n", r)

}

func run(p program, reg registers) registers {
	instructionPointer := 0
	counter := 0

	for instructionPointer < len(p.instructions) {
		reg[p.instructionPointer] = instructionPointer
		op := p.instructions[instructionPointer]
		ins := instructionSet[op.opcode]
		reg = ins(op)(reg)
		instructionPointer = reg[p.instructionPointer]
		instructionPointer++
		counter++
		//Uncomment to check the expected value on register 0 to stop the process
		/*if op.opcode == "eqrr" && (op.inputA == 0 || op.inputB == 0) {
			fmt.Printf("Stopped on %d\n", counter)
			break
		}*/

	}
	return reg
}
