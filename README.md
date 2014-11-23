goheap implements a binary heap.
The standard lib provides its own heap container, but that relies on its own interface.
This builds upon sort.Interface instead to allow for convenient access to its structures as well.

Example:
```go
data := sort.IntSlice{5, 8, 3, 2, 4, 9, 7, 6}
Heapify(data)
Heaped(data) // true
// insert a 1 into the heap
data = append(data, 1)
InsertFromEnd(data)
// data should contain 1-9
for i := 1; i <= 9; i++ {
    MinToEnd(data)
    val := data[len(data)-1]
    data = data[:len(data)-1]
    val == i // true
}
```
