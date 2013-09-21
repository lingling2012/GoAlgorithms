package algorithms

import (
	"testing"
	"fmt"
)

func TestWeightedQuickUnionWithPathCompressByHalving(t *testing.T) {
	const N = 10
	var cin = [][2]int{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		{8, 9},
		{5, 0},
		{7, 2},
		{6, 1},
		{1, 0},
		{6, 7},
	}
	var uf IUnionFind
	
	uf = NewWeightedQuickUnionWithPathCompressByHalving(N)
	for i := 0; i < len(cin); i++ {
		p := cin[i][0]
		q := cin[i][1]
		
		uf.Summary();
		fmt.Printf("\n");
		
		if uf.Connected(p, q) {
			continue
		}
		uf.Union(p, q)

		uf.Summary();
		fmt.Printf("%d %d\n", p, q)
	}
	fmt.Printf("%d components\n", uf.Count())
}
