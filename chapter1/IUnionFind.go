package chapter1

import ()

type VUnionFind interface {
	Union(p, q int)
	Find(p int) int
}

type IUnionFind interface {
	VUnionFind
	Connected(p, q int) bool
	Count() int
}

