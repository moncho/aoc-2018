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
	r := run(program{pointer, operations}, reg)
	fmt.Printf("Register values when the background process halts: %v\n", r)

	reg = make([]int, 6)
	reg[0] = 1
	reg = optimizedRun(program{pointer, operations}, reg)

	fmt.Printf("Register values when the background process halts: %v\n", reg)

}

func run(p program, reg registers) registers {
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

func optimizedRun(p program, reg registers) registers {
	instructionPointer := 0
	ticks := 0
	for instructionPointer < len(p.instructions) {
		ticks++
		reg[p.instructionPointer] = instructionPointer
		//This replaces the tight loop
		//The loop searches for prime factors of the number on reg[5]
		//starting from 0 until reg[5] and then backwards.
		//Every prime is sum to an accumulator on reg[0]
		if instructionPointer == 2 {
			pfs := primeFactors(reg[5])
			//Commented code iterates prime factors forwards
			//and then backwards as the original program would do.
			//But once the prime factors are found the expected result
			//can be achieved in one pass.
			for _, pf := range pfs {
				//reg[1] = pf
				//reg[0] = reg[1] + reg[0]
				//reg[4] = reg[5] / reg[1]
				reg[0] += pf + (reg[5] / pf)
			}
			/*for i := len(pfs) - 1; i >= 0; i-- {
				/*pf := pfs[i]
				reg[4] = pf
				reg[1] = reg[5] / reg[4]
				reg[0] = reg[1] + reg[0]

			}*/
			//Register are set to the breaking condition of the
			//original loop
			reg[1] = reg[5]
			reg[2] = 7
			reg[3] = 1
			reg[4] = 1
			instructionPointer = 8
			continue
		}

		op := p.instructions[instructionPointer]
		ins := instructionSet[op.opcode]
		reg = ins(op)(reg)
		instructionPointer = reg[p.instructionPointer]
		instructionPointer++
	}

	return reg
}
