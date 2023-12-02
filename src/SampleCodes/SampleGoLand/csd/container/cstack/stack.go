package cstack

type Stack struct {
	//...
}

func New() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(value interface{}) {
	//...
}

func (stack *Stack) Peek() interface{} {
	return nil
}

func (stack *Stack) Pop() interface{} {
	return nil
}

func (stack *Stack) Len() int {
	return 0
}
