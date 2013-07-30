package sort

import (

)

type InsertionSort struct{
	SortBase
}

func NewInsertionSort(a []IComparable) *InsertionSort{
	this := &InsertionSort{};
	this.SortBase.Super(a);
	return this;
}

func (this *InsertionSort) Sort(){
	N := len(this.a);
	for i:=1; i<N; i++{
		for j:=i; j>0 && this.Less(j, j-1); j--{
			this.Swap(j, j-1);
		}
	}
}