package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	values []interface{}
}

func (stack Stack) Size() int {
	return len(stack.values)
}

func (stack Stack) Empty() bool {
	return stack.Size() == 0
}

func (stack *Stack) StackUp(value interface{}) {
	stack.values = append(stack.values, value)
}

func (stack *Stack) Unstack() (interface{}, error) {
	if stack.Empty() {
		return nil, errors.New("Empty Stack")
	}
	value := stack.values[stack.Size()-1]
	stack.values = stack.values[:stack.Size()-1]
	return value, nil
}

func main() {
	stack := Stack{}
	stack.StackUp("Go")
	stack.StackUp(2009)
	stack.StackUp(3.14)
	stack.StackUp("End")
	fmt.Println("Size: ", stack.Size())
	fmt.Println("Empty? ", stack.Empty())
	for !stack.Empty() {
		v, _ := stack.Unstack()
		fmt.Println(v, " removed")
		fmt.Println("Size: ", stack.Size())
		fmt.Println("Empty? ", stack.Empty())
	}
}
