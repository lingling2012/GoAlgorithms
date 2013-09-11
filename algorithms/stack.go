package algorithms

import (
	"errors"
)

type Stack struct{
	S 	[]int
	top	  int
	size  int
}

func NewStack() *Stack{
	this := &Stack{};
	this.top = -1;
	this.size = 1;
	this.S = make([]int, this.size);
	return this;
}

func (this *Stack) Empty() bool{
	return this.top==-1;
}

func (this *Stack) Push(x int){
	this.top++;
	if this.top>=this.size-1 {
		A := make([]int, this.size*2);
		for i:=0; i<this.size; i++ {
			A[i] = this.S[i];
		}
		this.size = this.size*2;
		this.S    = A;
	}
	
	this.S[this.top] = x;
}

func (this *Stack) Pop() (int, error){
	if this.Empty() {
		return -1, errors.New("underflow");
	}else{
		this.top--;
		return this.S[this.top+1], nil;
	}
}