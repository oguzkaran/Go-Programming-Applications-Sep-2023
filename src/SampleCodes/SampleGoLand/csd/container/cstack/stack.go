package cstack

type Stack struct {
	//...
}

func New() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(value any) {
	//...
}

func (stack *Stack) Peek() any {
	return nil
}

func (stack *Stack) Pop() any {
	return nil
}

func (stack *Stack) Len() int {
	return 0
}
