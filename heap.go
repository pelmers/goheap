package goheap

import "sort" // for Interface

// children of some index i: i*2+1, i*2+2
// parent of some index i: (i-1) / 2  <-- truncating divison

// Modify data such that it becomes a valid min-heap.
// Perform at any time to re-establish heap invariants.
func Heapify(data sort.Interface) {
	for i := data.Len()/2 - 1; i >= 0; i-- {
		trickle(data, i, data.Len())
	}
}

// Return the index of the minimum element in heap data.
// Unfortunately without generics we can't actually return the element.
func Min(data sort.Interface) int {
	return 0
}

// Swap the minimum element to the last position in data.
// This element should then be popped by caller.
func MinToEnd(data sort.Interface) {
	data.Swap(0, data.Len()-1)
	trickle(data, 0, data.Len()-1)
}

// Insert the last element in data to the heap.
// The caller should append the element to be inserted, then call this method.
func InsertFromEnd(data sort.Interface) {
	percolate(data, data.Len()-1)
}

// Return whether data is a valid heap.
func Heaped(data sort.Interface) bool {
	for i := data.Len() - 1; i > 0; i-- {
		// child < parent invalidates heap property
		if data.Less(i, (i-1)/2) {
			return false
		}
	}
	return true
}

// Trickle the element at index down to the correct spot in the heap.
func trickle(data sort.Interface, index, length int) {
	// find the smallest child and swap down if it's bigger
	var min int
	for stop := length / 2; index < stop; index = min {
		min = baby(data, index, length)
		if data.Less(min, index) {
			data.Swap(min, index)
		} else {
			break
		}
	}
}

// Percolate the element at index up to its correct spot.
func percolate(data sort.Interface, index int) {
	// keep finding parent and swap up if it's less
	for index > 0 {
		if data.Less(index, (index-1)/2) {
			data.Swap(index, (index-1)/2)
		} else {
			break
		}
		index = (index - 1) / 2
	}
}

// Return the index of the smallest child of parent in data.
func baby(data sort.Interface, parent, length int) int {
	if parent*2+2 >= length {
		return parent*2 + 1
	}
	if data.Less(parent*2+1, parent*2+2) {
		return parent*2 + 1
	}
	return parent*2 + 2
}
