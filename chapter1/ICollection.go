package chapter1

import (

)

type Item struct {	
    // The value stored with this element.
    Value interface{}
    // contains filtered or unexported fields

	next  	*Item
}

func NewItem(i interface{}) *Item{
	this := &Item{}
	this.Value = i;
	return this;
}

func (this *Item) Next() *Item{
	return this.next
}

type IIterable interface{
	Iterator() *Item
}

type IBag interface{
	IIterable
	Add(interface{})
	IsEmpty() bool
	Size() int
}

type IQueue interface{
	IIterable
	Enqueue(interface{})
	Dequeue() Item
	IsEmpty() bool
	Size() int
}

type IStack interface{
	IIterable
	Push(interface{})
	Pop() Item
	IsEmpty() bool
	Size() int
}