package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type step struct {
	id        string
	dependsOn []string
	finished  bool
}

type instruction struct {
	step    string
	precond string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var instructions []instruction
	var i instruction
	for scanner.Scan() {
		_, err = fmt.Sscanf(scanner.Text(), "Step %s must be finished before step %s can begin.\n", &i.precond, &i.step)
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, i)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	steps := buildGraph(instructions)
	//root := root(steps)
	fmt.Printf("%s\n", complete("B", steps))
}

func buildGraph(instructions []instruction) map[string]*step {
	steps := make(map[string]*step)
	for _, i := range instructions {
		s, ok := steps[i.step]
		if ok {
			s.dependsOn = append(s.dependsOn, i.precond)
		} else {
			steps[i.step] = &step{
				id:        i.step,
				dependsOn: []string{i.precond},
			}
		}
		if _, ok := steps[i.precond]; !ok {
			steps[i.precond] = &step{
				id:        i.precond,
				dependsOn: []string{},
			}
		}
	}
	return steps
}

func complete(s string, steps map[string]*step) string {
	step := steps[s]

	if step == nil || step.finished {
		return ""
	}

	sort.Strings(step.dependsOn)
	var result string
	for _, pre := range step.dependsOn {
		result += complete(pre, steps)
	}
	step.finished = true
	return result + s
}
