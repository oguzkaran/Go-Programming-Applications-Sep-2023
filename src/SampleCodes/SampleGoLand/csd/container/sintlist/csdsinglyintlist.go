package sintlist

type SinglyIntNode struct {
	Value int
	pNext *SinglyIntNode
}

type CSDSinglyIntList struct {
	pHead *SinglyIntNode
	pTail *SinglyIntNode
	size  int
}

func New() *CSDSinglyIntList {
	return &CSDSinglyIntList{}
}

func (lst *CSDSinglyIntList) Clear() {
	//...
}

func (lst *CSDSinglyIntList) FindFirst(pred func(a int) bool) *SinglyIntNode {
	return nil
}

func (lst *CSDSinglyIntList) Front() *SinglyIntNode {
	return lst.pHead
}

func (lst *CSDSinglyIntList) Back() *SinglyIntNode {
	return lst.pTail
}

func (lst *CSDSinglyIntList) IsListEmpty() bool {
	return lst.size == 0
}

func (lst *CSDSinglyIntList) IsListNotEmpty() bool {
	return !lst.IsListEmpty()
}

func (lst *CSDSinglyIntList) InsertAfter(value int, node *SinglyIntNode) *SinglyIntNode {
	return nil
}

func (lst *CSDSinglyIntList) RemoveFront() (bool, int) {
	return false, 0
}

func (lst *CSDSinglyIntList) RemoveBack() (bool, int) {
	return false, 0
}

func (lst *CSDSinglyIntList) RemoveByValue(value int) bool {
	return false
}

func (lst *CSDSinglyIntList) RemoveByNode(node *SinglyIntNode) bool {
	return false
}

func (lst *CSDSinglyIntList) PushBack(value int) *SinglyIntNode {
	return nil
}

func (lst *CSDSinglyIntList) PushFront(value int) *SinglyIntNode {
	return nil
}

func (lst *CSDSinglyIntList) Size() int {
	return lst.size
}

func (lst *CSDSinglyIntList) WalkList(f func(int)) {

}

func (node *SinglyIntNode) Next() *SinglyIntNode {
	return node.pNext
}
