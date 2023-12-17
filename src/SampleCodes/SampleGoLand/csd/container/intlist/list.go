package intlist

type Node[T any] struct {
	Value T
	pNext *Node[T]
	pPrev *Node[T]
}

type List[T any] struct {
	pHead *Node[T]
	pTail *Node[T]
	size  int
}

func New[T any]() *List[T] {
	return &List[T]{}
}

func (lst *List[T]) Clear() {
	for pNode := lst.pHead; pNode.pNext != nil; pNode = pNode.pNext {
		pNode.pNext.pPrev = nil
	}

	lst.size = 0
	lst.pHead = nil
	lst.pTail = nil
}

func (lst *List[T]) FindFirst(pred func(a T) bool) *Node[T] {
	for pNode := lst.pHead; pNode != nil; pNode = pNode.pNext {
		if pred(pNode.Value) {
			return pNode
		}
	}
	return nil
}

func (lst *List[T]) FindLast(pred func(a T) bool) *Node[T] {
	for pNode := lst.pTail; pNode != nil; pNode = pNode.pPrev {
		if pred(pNode.Value) {
			return pNode
		}
	}

	return nil
}

func (lst *List[T]) Front() *Node[T] {
	return lst.pHead
}

func (lst *List[T]) Back() *Node[T] {
	return lst.pTail
}

func (lst *List[T]) IsEmpty() bool {
	return lst.size == 0
}

func (lst *List[T]) IsNotEmpty() bool {
	return !lst.IsEmpty()
}

func (lst *List[T]) InsertAfter(value int, node *Node[T]) *Node[T] {
	return nil
}

func (lst *List[T]) InsertBefore(value int, node *Node[T]) *Node[T] {
	return nil
}

func (lst *List[T]) RemoveFront() (bool, int) {
	return false, 0
}

func (lst *List[T]) RemoveBack() (bool, int) {
	return false, 0
}

func (lst *List[T]) RemoveByValue(value T) bool {
	return false
}

func (lst *List[T]) RemoveByNode(node *Node[T]) bool {
	return false
}

func (lst *List[T]) PushBack(value T) *Node[T] {
	pNewNode := &Node[T]{value, nil, nil}

	if lst.pHead != nil {
		lst.pTail.pNext = pNewNode
	} else {
		lst.pHead = pNewNode
	}

	pNewNode.pPrev = lst.pTail
	lst.pTail = pNewNode
	lst.size++

	return pNewNode
}

func (lst *List[T]) PushFront(value T) *Node[T] {
	pNewNode := &Node[T]{value, nil, nil}

	if lst.pHead != nil {
		lst.pHead.pPrev = pNewNode
	} else {
		lst.pTail = pNewNode
	}

	pNewNode.pNext = lst.pHead
	lst.pHead = pNewNode
	lst.size++

	return pNewNode
}

func (lst *List[T]) Size() int {
	return lst.size
}

func (lst *List[T]) WalkList(f func(T)) {
	for pNode := lst.pHead; pNode != nil; pNode = pNode.pNext {
		f(pNode.Value)
	}
}

func (node *Node[T]) Next() *Node[T] {
	return node.pNext
}

func (node *Node[T]) Prev() *Node[T] {
	return node.pPrev
}
