package algorithms

import (

)

import (
	"errors"
)

type Queue struct{
	Q 	[]int
	head  int
	tail  int
	size  int
}

func NewQueue() *Queue{
	this := &Queue{};
	this.head = 0;
	this.tail = 0;
	this.size = 1;
	this.Q = make([]int, this.size);
	return this;
}

func (this *Queue) Empty() bool{
	return this.head==this.tail;
}

func (this *Queue) Push(x int){
	this.Q[this.tail]= x;
	if this.head==(this.tail+1)%this.size {
		A := make([]int, this.size*2);
		for i:=0; i<this.size; i++ {
			A[i] = this.Q[(this.head+i)%this.size];
		}
		this.head = 0;
		this.tail = this.size; 
		this.size = this.size*2;
		this.Q    = A;
	}else{
		this.tail = (this.tail+1)%this.size;
	}
}

func (this *Queue) Pop() (int, error){
	if this.Empty() {
		return -1, errors.New("underflow");
	}else{
		x := this.Q[this.head];
		this.head = (this.head+1)%this.size;
		return x, nil;
	}
}