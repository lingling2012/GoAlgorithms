package main

import (
	"AlgorithmsInGo/sort"
	"bytes"
	"os"
	"strings"
)

//Read strings from standard input, sort them, and print
func main() {
	var strBuffer bytes.Buffer
	var a []sort.IComparable

	if len(os.Args) < 2{
		println("Too few arguments. Usage: [Selection|Insertion]Sort")
		return;
	}

	strBuffer.ReadFrom(os.Stdin)
	b := strings.Fields(strBuffer.String())

	a = make([]sort.IComparable, len(b))
	for i := 0; i < len(b); i++ {
		a[i] = sort.ComparableString(b[i])
	}

	var sorter sort.ISort
	switch os.Args[1] {
	case "SelectionSort":
		sorter = sort.NewSelectionSort(a)
		break
	case "InsertionSort":
		sorter = sort.NewInsertionSort(a)
		break
	}

	sorter.Sort()
	if sorter.IsSorted() {
		sorter.Show()
	} else {
		println("Sort Failed!")
	}
}
