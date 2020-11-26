package main

import "fmt"

func main() {
	var s Stacker
	fmt.Printf("The first value in stack is: %v\n", calculate(s, 1, 2, 3))
}

func calculate(stack Stacker, vals ...interface{}) interface{} {
	for _, val := range vals {
		stack.Append(val)
	}

	return stack.Pop()
}
