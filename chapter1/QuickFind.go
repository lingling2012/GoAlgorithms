package chapter1

import (

)

type QuickFind struct{
	UnionFind
}

func NewQuickFind(N int) *QuickFind{
	this := &QuickFind{};
	this.VUnionFind = this;
	this.UnionFind.Super(N);
	return this;
}

func (this *QuickFind) Find(p int) int{
	return this.id[p];
}

func (this *QuickFind) Union(p, q int){
	pID := this.VUnionFind.Find(p);
	qID := this.VUnionFind.Find(q);
	
	if pID==qID {
		return
	}
	
	for i:=0; i<len(this.id); i++ {
		if this.id[i]==pID {
			this.id[i] = qID;
		}
	}
	
	this.count--;
}