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

func (s step) duration() int {
	return 61 + int(s.id[0]) - 'A'
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
	req := newRequirements(instructions)

	order := req.completionOrder()

	fmt.Printf("Instruction steps order: %s\n", order)

	req.reset()
	order, time := req.parallelCompletionOrder(5)
	fmt.Printf("Instruction steps order: %s\n", order)
	fmt.Printf("Steps were completed in %d seconds\n", time)
}

type requirements map[string]*step

func (i requirements) completionOrder() string {

	root := i.root()
	var stack []*step
	stack = append(stack, root)
	result := ""

	for n := len(stack); n > 0; n = len(stack) {

		sort.Slice(stack, func(i, j int) bool {
			return stack[i].id < stack[j].id
		})
		current := stack[0]
		if !current.finished {
			current.finished = true
			result += current.id
			stack = append(stack[1:], i.completeStep(current.id)...)
		} else {
			stack = stack[1:]
		}
	}
	return result
}
func (i requirements) parallelCompletionOrder(wc int) (string, int) {
	root := i.root()
	var stack []*step

	result := ""
	workers := make(map[*step]int)
	workers[root] = root.duration()
	totalTime := 0
	for len(workers) > 0 {
		totalTime++
		var keys []*step
		for k := range workers {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].id < keys[j].id
		})
		for _, step := range keys {
			if workers[step] == 1 {
				delete(workers, step)
				step.finished = true
				result += step.id
				stack = append(stack, i.completeStep(step.id)...)
				sort.Slice(stack, func(i, j int) bool {
					return stack[i].id < stack[j].id
				})
				n := len(stack)
				index := 0
				for ; index < n && len(workers) < wc; index++ {
					bla := stack[index]
					workers[bla] = bla.duration()
				}
				stack = stack[index:]
			} else {
				workers[step]--
			}
		}

	}

	return result, totalTime
}

//completeStep returns the list of steps that can be processed
//after the completion of the given step.
func (i requirements) completeStep(s string) []*step {
	var result []*step
	for _, step := range i {
		if !step.finished && contains(s, step.dependsOn) && i.allTriggered(step.dependsOn) {
			result = append(result, step)
		}
	}
	return result
}
func (i requirements) root() *step {
	for _, step := range i {
		if len(step.dependsOn) == 0 {
			return step
		}
	}
	return nil
}

func (i requirements) reset() {
	for _, step := range i {
		step.finished = false
	}
}
func (i requirements) allTriggered(steps []string) bool {
	for _, id := range steps {
		if s := i[id]; !s.finished {
			return false
		}
	}
	return true
}

func newRequirements(instructions []instruction) requirements {
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

func contains(s string, ss []string) bool {
	for _, e := range ss {
		if e == s {
			return true
		}
	}
	return false
}
