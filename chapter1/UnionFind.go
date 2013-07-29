package chapter1

import ()

type UnionFind struct {
	VUnionFind

	id    []int
	count int
}

func NewUnionFind(N int) *UnionFind {
	this := &UnionFind{}
	this.VUnionFind = this
	this.count = N
	this.id = make([]int, N)
	for i := 0; i < N; i++ {
		this.id[i] = i
	}
	return this
}

func (this *UnionFind) Super(N int) {
	this.count = N
	this.id = make([]int, N)
	for i := 0; i < N; i++ {
		this.id[i] = i
	}
}

func (this *UnionFind) Count() int {
	return this.count
}

func (this *UnionFind) Connected(p, q int) bool {
	return this.VUnionFind.Find(p) == this.VUnionFind.Find(q)
}

func (this *UnionFind) Find(p int) int {
	return 0
}

func (this *UnionFind) Union(p, q int) {
	return
}
