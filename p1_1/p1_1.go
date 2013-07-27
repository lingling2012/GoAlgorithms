package p1_1

import (
	"fmt"
)

func FastFind(cin [][2]int, N int) {
	var p, q int
	var id []int

	fmt.Printf("(p q):  ")

	id = make([]int, N)
	for i := 0; i < N; i++ {
		id[i] = i
		fmt.Printf("%d ", id[i])
	}
	fmt.Printf("\n\n")

	for l := 0; l < len(cin); l++ {
		p = cin[l][0]
		q = cin[l][1]

//////////////////////////////////////
		t := id[p]
		if t != id[q] {
			for i := 0; i < N; i++ {
				if id[i] == t {
					id[i] = id[q]
				}
			}
		}
//////////////////////////////////////
		
		fmt.Printf("(%d %d):  ", p, q)
		for k := 0; k < N; k++ {
			fmt.Printf("%d ", id[k])
		}
		
		if t != id[q] {
			fmt.Printf(":  (%d %d) \n", p, q);
		}else{
			fmt.Printf("\n");
		}
	}
}
