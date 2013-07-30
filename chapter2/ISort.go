package chapter2

import (
	"os"
)

type ISortBase interface{
	Less(i, j int) bool;
 	Swap(i, j int);
 	Show()
 	IsSorted() bool;
}

type ISort interface{
	ISortBase
	Sort()
}


type SortBase struct{
	a	[]IComparable
}

func NewSortBase(a []IComparable) *SortBase{
	this := &SortBase{}
	this.a = a;
	return this;
}

func (this *SortBase) Super(a []IComparable){
	this.a = a;
}

func (this *SortBase) Less(i, j int) bool{
	return this.a[i].CompareTo(this.a[j]) < 0;
}

func (this *SortBase) Swap(i, j int){
	t := this.a[i];
	this.a[i] = this.a[j];
	this.a[j] = t;
}

func (this *SortBase) Show(){
	for i:=0; i<len(this.a); i++ {
		os.Stdout.WriteString(this.a[i].String()+" ");
	}
	os.Stdout.WriteString("\n");
}

func (this *SortBase) IsSorted() bool{
	for i:=1; i<len(this.a); i++ {
		if this.Less(i, i-1) {
			return false;
		}
	}
	return true;
}