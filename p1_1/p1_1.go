package p1_1

import (
	"fmt"
)

func FastFind(cin [][2]int, N int) {	
	id := make([]int, N)
	
	fmt.Printf("(p q):  ")
	for i := 0; i < N; i++ {
		id[i] = i
		fmt.Printf("%d ", id[i])
	}
	fmt.Printf(": (id_r id_w)\n\n")

	for l := 0; l < len(cin); l++ {
		p := cin[l][0]
		q := cin[l][1]

		id_r := 0;
		id_w := 0;
//////////////////////////////////////
		t := id[p]
		id_r++;
		if t != id[q] {
			for i := 0; i < N; i++ {
				if id[i] == t {
					id[i] = id[q]
					id_r++;
					id_w++;
				}
				id_r++;
			}
		}
		id_r++;
//////////////////////////////////////
		
		fmt.Printf("(%d %d):  ", p, q)
		for k := 0; k < N; k++ {
			fmt.Printf("%d ", id[k])
		}
		
		fmt.Printf(": (%4d %4d) ", id_r, id_w);
		
		if t != id[q] {
			fmt.Printf(":  (%d %d) \n", p, q);
		}else{
			fmt.Printf("\n");
		}
	}
}
