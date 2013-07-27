package sort

import (

)

type SelectionSort struct{
	SortBase
}

func NewSelectionSort(a []IComparable) *SelectionSort{
	this := &SelectionSort{};
	this.SortBase.Super(a);
	return this;
}

func (this *SelectionSort) Sort(){
	N := len(this.a);
	for i:=0; i<N; i++ {
		min := i;
		for j:=i+1; j<N; j++{
			if this.Less(j, min) {
				min = j;
			}
		}
		this.Swap(i, min);
	}
}
