package clist

import "container/list"

type List struct {
	CList list.List
}

func New() *List {
	return &List{}
}

func (lst *List) WalkList(f func(a any)) {
	for e := lst.CList.Front(); e != nil; e = e.Next() {
		f(e.Value)
	}
}

//...
