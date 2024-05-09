package main

import "strconv"

func calPoints(ops []string) int {
	var stack []int
	for _, op := range ops {
		switch op {
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		case "D":
			stack = append(stack, stack[len(stack)-1]*2)
		case "C":
			stack = stack[:len(stack)-1]
		default:
			score, _ := strconv.Atoi(op)
			stack = append(stack, score)
		}
	}
	var count int
	for _, i := range stack {
		count += i
	}
	return count
}
