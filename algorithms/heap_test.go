package ch06

import (
    "testing"
    "fmt"
)

func TestHeap(t *testing.T) {
	var A =[]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7};
    
	h := NewHeap(A, len(A));
	
 	h.BuildMaxHeap();
 	
 	for i:=0; i<len(A); i++{
 		fmt.Printf("%d ", A[i]);
 	}
 	fmt.Printf("\n");
 	
 	//h.MaxHeapSort();
	h.IncreaseKey(8, 15);
	
 	for i:=0; i<len(A); i++{
 		fmt.Printf("%d ", A[i]);
 	}
 	fmt.Printf("\n");

 	for i:=0; i<len(A); i++{
 		max,_ := h.ExtractMax();
 		fmt.Printf("%d ", max);
 	}
 	fmt.Printf("\n"); 	
}

