package Heap

import "fmt"

//Heap Data Struct
type Heap struct {
	Elements    []int
	Capacity    int
	CurrentSize int
}

//InitHeap initilizes heap
func InitHeap(size int) *Heap {
	return &Heap{Capacity: size, Elements: make([]int, size)}
}

//Left finds left child
func (a *Heap) Left(i int) int {
	return i*2 + 1
}

//Right finds right child
func (a *Heap) Right(i int) int {
	return i*2 + 2
}

//Parent finds parent
func (a *Heap) Parent(i int) int {
	return (i - 1) / 2
}

//Insert inserts element to heap
func (a *Heap) Insert(v int) {
	if a.CurrentSize == a.Capacity {
		fmt.Println("Heap is full")
		return
	}
	a.Elements[a.CurrentSize] = v
	a.CurrentSize++
	index := a.CurrentSize - 1
	for index != 0 && a.Elements[a.Parent(index)] > a.Elements[index] {
		a.Elements[a.Parent(index)], a.Elements[index] = a.Elements[index], a.Elements[a.Parent(index)]
		index = a.Parent(index)
	}
}

//Delete index from heap
func (a *Heap) Delete(i int) {
	a.DecreaseKey(i, -10*6)
	a.ExtractMin()
	a.Heapify(0)
}

//ExtractMin extracts minimum element
func (a *Heap) ExtractMin() int {
	if a.CurrentSize == 1 {
		a.CurrentSize--
		return a.Elements[0]
	}
	v := a.Elements[0]
	a.Elements[0] = a.Elements[a.CurrentSize-1]
	a.CurrentSize--
	a.Heapify(0)
	return v
}

//GetMin retursn minimum element value
func (a *Heap) GetMin() int {
	return a.Elements[0]
}

//DecreaseKey decreases value of key
func (a *Heap) DecreaseKey(key int, nval int) {
	a.Elements[key] = nval
	index := key
	for index != 0 && a.Elements[a.Parent(index)] > a.Elements[index] {
		a.Elements[a.Parent(index)], a.Elements[index] = a.Elements[index], a.Elements[a.Parent(index)]
		index = a.Parent(index)
	}
}

//Heapify for heap
func (a *Heap) Heapify(i int) {
	l := a.Left(i)
	r := a.Right(i)
	smallest := i
	if l < a.CurrentSize && a.Elements[l] < a.Elements[i] {
		smallest = l
	}

	if r < a.CurrentSize && a.Elements[r] < a.Elements[smallest] {
		smallest = r
	}

	if smallest != i {
		a.Elements[smallest], a.Elements[i] = a.Elements[i], a.Elements[smallest]
		a.Heapify(smallest)
	}

}
