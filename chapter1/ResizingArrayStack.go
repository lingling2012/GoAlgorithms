package chapter1

import (

)

type ResizingArrayStack struct{
	a []*Item
	N int
}

func NewResizingArrayStack() *ResizingArrayStack{
	this := &ResizingArrayStack{}
	this.N = 0;
	this.a = make([]*Item, 1);
	this.a[0] = nil;
	return this;
}

func (this *ResizingArrayStack) IsEmpty() bool{
	return this.N==0
}

func (this *ResizingArrayStack) Size() int{
	return this.N
}

func (this *ResizingArrayStack) Resize(max int){
	temp := make([]*Item, max)
	for i:=0; i<this.N; i++{
		temp[i] = this.a[i]
	}
	
	this.a = temp;
}

func (this *ResizingArrayStack) Push(value interface{}){
	if this.N == len(this.a) {
		this.Resize(2*len(this.a));
	}
	
	this.a[this.N] = NewItem(value);
	this.a[this.N].next = nil;
	if this.N-1>0{
		this.a[this.N-1].next = this.a[this.N];
	}
	this.N++;
}

func (this *ResizingArrayStack) Pop() *Item{
	this.N--;
	item := this.a[this.N];
	this.a[this.N] = nil;
	if this.N-1>0{
		this.a[this.N-1].next = nil;
	}
	if this.N>0 && this.N == len(this.a)/4 {
		this.Resize(len(this.a)/2);
	}
	return item;
}

func (this *ResizingArrayStack) Iterator() *Item{
	return this.a[0];
}
