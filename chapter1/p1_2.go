package p1_2

import (
	"fmt"
)

func FastUnion(cin [][2]int, N int) {
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
		var i, j int;
		for i=p; i!=id[i]; i, id_r =id[i], id_r+1 {
			id_r++;
		}
		id_r++;
		for j=q; j!=id[j]; j, id_r =id[j], id_r+1 {
			id_r++;
		}
		id_r++
		if i!=j {
			id[i] = j;
			id_w++;
		}
//////////////////////////////////////

		fmt.Printf("(%d %d):  ", p, q)
		for k := 0; k < N; k++ {
			fmt.Printf("%d ", id[k])
		}
		
		fmt.Printf(": (%4d %4d) ", id_r, id_w);
		
		if i != j {
			fmt.Printf(":  (%d %d) \n", p, q);
		}else{
			fmt.Printf("\n");
		}
	}
}

