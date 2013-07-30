package chapter2

import (
	"testing"
	"bytes"
	"os"
	"strings"
)

//Read strings from standard input, sort them, and print
func TestSort(t *testing.T) {
	var strBuffer bytes.Buffer
	var a []IComparable

	if len(os.Args) < 2{
		println("Too few arguments. Usage: [Selection|Insertion]Sort")
		return;
	}

	strBuffer.ReadFrom(os.Stdin)
	b := strings.Fields(strBuffer.String())

	a = make([]IComparable, len(b))
	for i := 0; i < len(b); i++ {
		a[i] = ComparableString(b[i])
	}

	var sorter ISort
	switch os.Args[1] {
	case "SelectionSort":
		sorter = NewSelectionSort(a)
		break
	case "InsertionSort":
		sorter = NewInsertionSort(a)
		break
	}

	sorter.Sort()
	if sorter.IsSorted() {
		sorter.Show()
	} else {
		println("Sort Failed!")
	}
}
