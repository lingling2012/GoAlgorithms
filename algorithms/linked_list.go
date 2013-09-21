package algorithms

import (
	"fmt"
)

type LinkedElement struct{
	key  int;
	data interface{}
	
	prev *LinkedElement;
	next *LinkedElement;
}

type LinkedList struct{
	sentinel *LinkedElement;
}

func NewLinkedList() *LinkedList{
	this := &LinkedList{};
	this.sentinel = &LinkedElement{};
	this.sentinel.prev = this.sentinel;
	this.sentinel.next = this.sentinel;
	return this;
}

func (this *LinkedList) Insert(x *LinkedElement){
	x.next = this.sentinel.next;
	this.sentinel.next.prev = x;
	this.sentinel.next = x;
	x.prev = this.sentinel;
}

func (this *LinkedList) Delete(x *LinkedElement){
	if x!=this.sentinel && x!=nil{
		x.prev.next = x.next;
		x.next.prev = x.prev;
	}	
}

func (this *LinkedList) Search(k int) *LinkedElement{
	x := this.sentinel.next;
	for x!=this.sentinel && x.key != k{
		x = x.next;
	}
	if x==this.sentinel{
		return nil;
	}else{
		return x;
	}
}

func (this *LinkedList) Summary(){
	x := this.sentinel.next;
	fmt.Printf("nil");
	for x!=this.sentinel{
		fmt.Printf("->%d", x.key);
		x = x.next;
	}
	fmt.Printf("\n");
}