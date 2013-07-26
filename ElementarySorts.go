package main

import (
	"os"	
	"bytes"
	"strings"
	"AlgorithmsInGo/sort"
)

//Read strings from standard input, sort them, and print
func main(){
	var strBuffer bytes.Buffer;
	var a []sort.IComparable;

	strBuffer.ReadFrom(os.Stdin)
	b := strings.Fields(strBuffer.String())
	
	a = make([]sort.IComparable, len(b));
	for i:=0; i<len(b); i++ {
		a[i] = sort.ComparableString(b[i]);
	}

	selectionSorter := sort.NewSelectionSort(a);
	selectionSorter.Sort();
	if selectionSorter.IsSorted() {
		selectionSorter.Show();
	}else{
		println("Sort Failed!");
	}
}




