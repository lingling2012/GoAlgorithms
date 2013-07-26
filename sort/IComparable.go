package sort

import (

)

type IComparable interface{
	CompareTo(that interface{}) int;
	String() string;
}