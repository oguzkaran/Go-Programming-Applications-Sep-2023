package cqueue

type Queue struct {
	//...
}

func New() *Queue {
	return &Queue{}
}

func (stack *Queue) Enqueue(value any) {
	//...
}

func (stack *Queue) Dequeue() any {
	return nil
}

func (stack *Queue) Peek() any {
	return nil
}

func (stack *Queue) Len() int {
	return 0
}
