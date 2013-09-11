package algorithms

import (
	"errors"
	"math"
)


type Heap struct{
	length int
	heap_size int
	A	[]int
}


func NewHeap(A []int, length int) *Heap{
	this := &Heap{};
	
	this.length = length;
	this.heap_size = length;
	this.A = A;

	return this;
}

func (this *Heap) Parent(i int) int{
	if i==0{
		return 0;
	}else{
		return (i-1)>>1;
	}
}

func (this *Heap) Left(i int) int{
	return i*2+1;
}

func (this *Heap) Right(i int) int{
	return i*2+2;
}

func (this *Heap) MaxHeapify(i int){
	l := this.Left(i);
	r := this.Right(i);
	var largest int;
	if l<this.heap_size && this.A[l]>this.A[i] {
		largest = l;
	}else{
		largest = i;
	}
	if r<this.heap_size && this.A[r]>this.A[largest]{
		largest = r;
	}
	if largest!=i{
		this.A[i], this.A[largest] = this.A[largest], this.A[i];
		this.MaxHeapify(largest);
	}
}

func (this *Heap) MinHeapify(i int){
	l := this.Left(i);
	r := this.Right(i);
	var smallest int;
	if l<this.heap_size && this.A[l]<this.A[i] {
		smallest = l;
	}else{
		smallest = i;
	}
	if r<this.heap_size && this.A[r]<this.A[smallest]{
		smallest = r;
	}
	if smallest!=i{
		this.A[i], this.A[smallest] = this.A[smallest], this.A[i];
		this.MinHeapify(smallest);
	}
}

func (this *Heap) BuildMaxHeap(){
	this.heap_size = this.length;
	for i:= this.Parent(this.length-1); i>=0; i--{
		this.MaxHeapify(i);
	}
}

func (this *Heap) BuildMinHeap(){
	this.heap_size = this.length;
	for i:= this.Parent(this.length-1); i>=0; i--{
		this.MinHeapify(i);
	}
}

func (this *Heap) MaxHeapSort(){
	this.BuildMaxHeap();
	for i:=this.length-1; i>=1; i--{
		this.A[0], this.A[i] = this.A[i], this.A[0];
		this.heap_size--;
		this.MaxHeapify(0);
	}
}

func (this *Heap) MinHeapSort(){
	this.BuildMinHeap();
	for i:=this.length-1; i>=1; i--{
		this.A[0], this.A[i] = this.A[i], this.A[0];
		this.heap_size--;
		this.MinHeapify(0);
	}
}

func (this *Heap) Maximum() int{
	return this.A[0];
}

func (this *Heap) Minimum() int{
	return this.A[0];
}

func (this *Heap) ExtractMax() (int, error){
	if this.heap_size == 0{
		return -1, errors.New("heap underflow");
	}else if this.heap_size==1{
		this.heap_size--;
		return this.A[0], nil;
	}else{
		max := this.A[0];
		this.A[0] = this.A[this.heap_size-1];
		this.heap_size--;
		this.MaxHeapify(0);
		return max, nil;
	}
}

func (this *Heap) ExtractMin() (int, error){
	if this.heap_size == 0{
		return -1, errors.New("heap underflow");
	}else if this.heap_size==1{
		this.heap_size--;
		return this.A[0], nil;
	}else{
		min := this.A[0];
		this.A[0] = this.A[this.heap_size-1];
		this.heap_size--;
		this.MinHeapify(0);
		return min, nil;
	}
}

func (this *Heap) IncreaseKey(i, key int) error{
	if key < this.A[i] {
		return errors.New("new key is smaller than current key");
	}
	this.A[i] = key;
	for i>0 && this.A[this.Parent(i)] < this.A[i] {
		this.A[i], this.A[this.Parent(i)] =  this.A[this.Parent(i)], this.A[i];
		i = this.Parent(i);
	}
	
	return nil;
}

func (this *Heap) DecreaseKey(i, key int) error{
	if key > this.A[i] {
		return errors.New("new key is larger than current key");
	}
	this.A[i] = key;
	for i>0 && this.A[this.Parent(i)] > this.A[i] {
		this.A[i], this.A[this.Parent(i)] =  this.A[this.Parent(i)], this.A[i];
		i = this.Parent(i);
	}
	
	return nil;
}

func (this *Heap) InsertMax(key int) {
	this.heap_size++;
	this.A[this.heap_size-1] = math.MinInt32;
	this.IncreaseKey(this.heap_size-1, key);
}

func (this *Heap) InsertMin(key int) {
	this.heap_size++;
	this.A[this.heap_size-1] = math.MaxInt32;
	this.DecreaseKey(this.heap_size-1, key);
}



