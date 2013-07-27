package p1_2

import (
	"fmt"
)

func FastUnion(cin [][2]int, N int) {
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
		var i, j int;
		for i=p; i!=id[i]; i=id[i] {
		}
		for j=q; j!=id[j]; j=id[j] {
		}
		if i!=j {
			id[i] = j;
		}
//////////////////////////////////////

		fmt.Printf("(%d %d):  ", p, q)
		for k := 0; k < N; k++ {
			fmt.Printf("%d ", id[k])
		}
		
		if i != j {
			fmt.Printf(":  (%d %d) \n", p, q);
		}else{
			fmt.Printf("\n");
		}
	}
}

