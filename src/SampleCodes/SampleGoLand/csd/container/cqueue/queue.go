package cqueue

type Queue struct {
	//...
}

func New() *Queue {
	return &Queue{}
}

func (stack *Queue) Enqueue(value interface{}) {
	//...
}

func (stack *Queue) Dequeue() interface{} {
	return nil
}

func (stack *Queue) Peek() interface{} {
	return nil
}

func (stack *Queue) Len() int {
	return 0
}
