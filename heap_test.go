package goheap

import (
	"sort"
	"testing"
)

func isHeap(data sort.Interface, t *testing.T) {
	if !Heaped(data) {
		t.Errorf("Not correctly heapified: %v", data)
	}
}

func TestHeapify(t *testing.T) {
	data1 := sort.IntSlice{1, 2, 3, 4, 5, 6, 7, 8}
	Heapify(data1)
	data2 := sort.Float64Slice{1.1, 139.4, 3.4, 0.5, 0.1, 10.6}
	Heapify(data2)
	data3 := sort.IntSlice{1, 1, 1, 1, 1, 1, 1, 1, 1}
	Heapify(data3)
	isHeap(data1, t)
	isHeap(data2, t)
	isHeap(data3, t)
}

func TestUsage(t *testing.T) {
	data := sort.IntSlice{5, 8, 3, 2, 4, 9, 7, 6}
	Heapify(data)
	isHeap(data, t)
	data = append(data, 1)
	InsertFromEnd(data)
	// data should contain 1-9
	for i := 1; i <= 9; i++ {
		MinToEnd(data)
		val := data[len(data)-1]
		data = data[:len(data)-1]
		if val != i {
			t.Errorf("Heap should pop things in sorted order")
		}
	}
}
