package algorithms

import (
    "testing"
    //"fmt"
)

func TestLinkedList(t *testing.T) {
	var A =[]int{1, 4, 16, 9, 25};
	
	l := NewLinkedList();
	
	for i:=0; i<len(A); i++{
		x := &LinkedElement{A[i], A[i], nil, nil};
		l.Insert(x);
		l.Summary();
	}
	
	for i:=0; i<len(A); i++{
		x := l.Search(A[i]);
		l.Delete(x);
		l.Summary();
	}
	
	x := l.Search(1);
		l.Delete(x);
		l.Summary();
}