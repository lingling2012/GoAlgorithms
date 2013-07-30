package chapter2

import (

)

type IComparable interface{
	CompareTo(that interface{}) int;
	String() string;
}