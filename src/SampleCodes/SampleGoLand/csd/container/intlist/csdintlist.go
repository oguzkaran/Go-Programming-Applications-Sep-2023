package intlist

type IntNode struct {
	Value int
	pNext *IntNode
	pPrev *IntNode
}

type CSDIntList struct {
	pHead *IntNode
	pTail *IntNode
	size  int
}

func New() *CSDIntList {
	return &CSDIntList{}
}

func (lst *CSDIntList) Clear() {
	for pNode := lst.pHead; pNode.pNext != nil; pNode = pNode.pNext {
		pNode.pNext.pPrev = nil
	}

	lst.size = 0
	lst.pHead = nil
	lst.pTail = nil
}

func (lst *CSDIntList) FindFirst(pred func(a int) bool) *IntNode {
	for pNode := lst.pHead; pNode != nil; pNode = pNode.pNext {
		if pred(pNode.Value) {
			return pNode
		}
	}
	return nil
}

func (lst *CSDIntList) FindLast(pred func(a int) bool) *IntNode {
	for pNode := lst.pTail; pNode != nil; pNode = pNode.pPrev {
		if pred(pNode.Value) {
			return pNode
		}
	}

	return nil
}

func (lst *CSDIntList) Front() *IntNode {
	return lst.pHead
}

func (lst *CSDIntList) Back() *IntNode {
	return lst.pTail
}

func (lst *CSDIntList) IsEmpty() bool {
	return lst.size == 0
}

func (lst *CSDIntList) IsNotEmpty() bool {
	return !lst.IsEmpty()
}

func (lst *CSDIntList) InsertAfter(value int, node *IntNode) *IntNode {
	return nil
}

func (lst *CSDIntList) InsertBefore(value int, node *IntNode) *IntNode {
	return nil
}

func (lst *CSDIntList) RemoveFront() (bool, int) {
	return false, 0
}

func (lst *CSDIntList) RemoveBack() (bool, int) {
	return false, 0
}

func (lst *CSDIntList) RemoveByValue(value int) bool {
	return false
}

func (lst *CSDIntList) RemoveByNode(node *IntNode) bool {
	return false
}

func (lst *CSDIntList) PushBack(value int) *IntNode {
	pNewNode := &IntNode{value, nil, nil}

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

func (lst *CSDIntList) PushFront(value int) *IntNode {
	pNewNode := &IntNode{value, nil, nil}

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

func (lst *CSDIntList) Size() int {
	return lst.size
}

func (lst *CSDIntList) WalkList(f func(int)) {
	for pNode := lst.pHead; pNode != nil; pNode = pNode.pNext {
		f(pNode.Value)
	}
}

func (node *IntNode) Next() *IntNode {
	return node.pNext
}

func (node *IntNode) Prev() *IntNode {
	return node.pPrev
}
