package slist

type Node[T any] struct {
	Value int
	pNext *Node[T]
}

type CSDSList[T any] struct {
	pHead *Node[T]
	pTail *Node[T]
	size  int
}

func New[T any]() *CSDSList[T] {
	return &CSDSList[T]{}
}

func (lst *CSDSList[T]) Clear() {
	//...
}

func (lst *CSDSList[T]) FindFirst(pred func(a int) bool) *Node[T] {
	return nil
}

func (lst *CSDSList[T]) Front() *Node[T] {
	return lst.pHead
}

func (lst *CSDSList[T]) Back() *Node[T] {
	return lst.pTail
}

func (lst *CSDSList[T]) IsListEmpty() bool {
	return lst.size == 0
}

func (lst *CSDSList[T]) IsListNotEmpty() bool {
	return !lst.IsListEmpty()
}

func (lst *CSDSList[T]) InsertAfter(value int, node *Node[T]) *Node[T] {
	return nil
}

func (lst *CSDSList[T]) RemoveFront() (bool, int) {
	return false, 0
}

func (lst *CSDSList[T]) RemoveBack() (bool, int) {
	return false, 0
}

func (lst *CSDSList[T]) RemoveByValue(value int) bool {
	return false
}

func (lst *CSDSList[T]) RemoveByNode(node *Node[T]) bool {
	return false
}

func (lst *CSDSList[T]) PushBack(value int) *Node[T] {
	return nil
}

func (lst *CSDSList[T]) PushFront(value int) *Node[T] {
	return nil
}

func (lst *CSDSList[T]) Size() int {
	return lst.size
}

func (lst *CSDSList[T]) WalkList(f func(int)) {

}

func (node *Node[T]) Next() *Node[T] {
	return node.pNext
}
