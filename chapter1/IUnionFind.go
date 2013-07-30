package chapter1

import ()

type VUnionFind interface { //Virtual Part of UnionFind Interface
	Union(p, q int)
	Find(p int) int
}

type IUnionFind interface {
	VUnionFind
	Connected(p, q int) bool
	Count() int
	Summary()
}
