package chapter1

import (

)


type WeightedQuickUnionWithPathCompressByHalving struct {
	UnionFind
	sz []int
}

func NewWeightedQuickUnionWithPathCompressByHalving(N int) *WeightedQuickUnionWithPathCompressByHalving {
	this := &WeightedQuickUnionWithPathCompressByHalving{}
	this.VUnionFind = this
	this.UnionFind.Super(N)
	this.sz = make([]int, N)
	for i:=0; i<N; i++{
		this.sz[i] = 1;
	}
	return this
}

func (this *WeightedQuickUnionWithPathCompressByHalving) Find(p int) int {
	for p != this.id[p] {
		this.id[p] = this.id[this.id[p]]
		p = this.id[p]
	}
	return p
}

func (this *WeightedQuickUnionWithPathCompressByHalving) Union(p, q int) {
	pRoot := this.Find(p)
	qRoot := this.Find(q)

	if pRoot == qRoot {
		return
	}
	
	if this.sz[pRoot] < this.sz[qRoot] {
		this.id[pRoot] = qRoot; 
		this.sz[qRoot] += this.sz[pRoot];
	}else{
		this.id[qRoot] = pRoot
		this.sz[pRoot] += this.sz[qRoot];
	}

	this.count--
}