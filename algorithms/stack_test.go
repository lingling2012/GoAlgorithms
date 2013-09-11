package algorithms

import (
    "testing"
	"fmt"
)

func TestStack(t *testing.T) {
	var A =[]int{15, 6, 2, 9, 17, 3};
	
	s := NewStack();
	
	for i:=0; i<len(A); i++{
		s.Push(A[i]);
		for j:=0; j<len(s.S); j++{
			fmt.Printf("%d ", s.S[j]);
		}
		fmt.Printf("\n");
	}

	for i:=len(A)-1; i>=-1; i--{
		a, err := s.Pop(); 
		if err!=nil{
			fmt.Printf("%s\n", err.Error());
		}else{
			fmt.Printf("%d :", a);
			for j:=0; j<len(s.S); j++{
				fmt.Printf("%d ", s.S[j]);
			}
			fmt.Printf("\n");
		}
	}
}

