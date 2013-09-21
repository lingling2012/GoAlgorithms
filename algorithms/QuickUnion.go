package algorithms

import ()

type QuickUnion struct {
	UnionFind
}

func NewQuickUnion(N int) *QuickUnion {
	this := &QuickUnion{}
	this.VUnionFind = this
	this.UnionFind.Super(N)
	return this
}

func (this *QuickUnion) Find(p int) int {
	for p != this.id[p] {
		p = this.id[p]
	}
	return p
}

func (this *QuickUnion) Union(p, q int) {
	pRoot := this.Find(p)
	qRoot := this.Find(q)

	if pRoot == qRoot {
		return
	}

	this.id[pRoot] = qRoot

	this.count--
}
