package sort

import (
	"bytes"
)

type ComparableString string

func (this ComparableString) CompareTo(that interface{}) int{
	a := []byte(this);
	b := []byte(that.(ComparableString));
	return bytes.Compare(a, b);
}

func (this ComparableString) String() string{
	return string(this);
}
