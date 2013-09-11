package algorithms

import (
    "testing"
    "fmt"
)

func TestQueue(t *testing.T) {
	var A =[]int{15, 6, 2, 9, 17, 3, 1};
	
	s := NewQueue();
	
	for i:=0; i<len(A); i++{
		s.Push(A[i]);
		for j:=0; j<len(s.Q); j++{
			fmt.Printf("%d ", s.Q[j]);
		}
		fmt.Printf("\n");
	}

	for i:=len(A)-1; i>=-1; i--{
		a, err := s.Pop(); 
		if err!=nil{
			fmt.Printf("%s\n", err.Error());
		}else{
			fmt.Printf("%d :", a);
			for j:=0; j<len(s.Q); j++{
				fmt.Printf("%d ", s.Q[j]);
			}
			fmt.Printf("\n");
		}
	}
	
	for i:=0; i<len(A); i++{
		s.Push(A[i]);
		for j:=0; j<len(s.Q); j++{
			fmt.Printf("%d ", s.Q[j]);
		}
		fmt.Printf("\n");
	}
	for i:=0; i<len(A); i++{
		s.Push(A[i]);
		for j:=0; j<len(s.Q); j++{
			fmt.Printf("%d ", s.Q[j]);
		}
		fmt.Printf("\n");
	}
}

