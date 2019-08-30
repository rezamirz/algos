/*

heap.go

An implementation of Min/Max heap.
In order to build a heap one interface must be implemented:
  - A comparator interface to compare two elements in the heap.
  - Based on return value of comparator the heap can be min or max heap.

MIT License

Copyright (c) 2018 rezamirz

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package pq

type IHeap interface {
	IsEmpty() bool
	Size() uint32
	Put(interface{})
	Top() interface{}       // Returns min or max element of the heap if it is minheap or maxheap
	DeleteTop() interface{} // Deletes min or max element of the heap if it is minheap or maxheap
}

type Comparator interface {
	/*
	 * Compares two keys and returns -1, 0, or 1.
	 * The return value is used to sort the keys.
	 */
	Compare(k1, k2 interface{}) int
}

type Heap struct {
	maxCapacity uint32        /* The size of data array */
	size        uint32        /* The number of entries in heap */
	data        []interface{} /* Array of data pointers */
	comparator  Comparator
}

func NewHeap(capacity uint32, comparator Comparator) IHeap {

	heap := &Heap{}
	heap.data = make([]interface{}, capacity)
	heap.comparator = comparator
	heap.maxCapacity = capacity

	return heap
}

func (heap *Heap) IsEmpty() bool {
	return heap.size == 0
}

func (heap *Heap) Size() uint32 {
	return heap.size
}

func (heap *Heap) Top() interface{} {
	return heap.data[1]
}

func (heap *Heap) resize(newCapacity uint32) {
	data := make([]interface{}, newCapacity)

	copy(data, heap.data[:heap.size+1])
	heap.data = data
	heap.maxCapacity = newCapacity
}

func (heap *Heap) exch(i, j uint32) {
	tmp := heap.data[i]
	heap.data[i] = heap.data[j]
	heap.data[j] = tmp
}

func (heap *Heap) swim() {
	n := heap.size

	for n > 1 && heap.comparator.Compare(heap.data[n/2], heap.data[n]) > 0 {
		heap.exch(n/2, n)
		n = n / 2
	}
}

func (heap *Heap) sink(k uint32) {
	for 2*k <= heap.size {
		j := 2 * k
		if j < heap.size && heap.comparator.Compare(heap.data[j], heap.data[j+1]) > 0 {
			j++
		}
		if heap.comparator.Compare(heap.data[k], heap.data[j]) <= 0 {
			break
		}
		heap.exch(k, j)
		k = j
	}
}

func (heap *Heap) Put(data interface{}) {
	if heap.size+1 == heap.maxCapacity {
		heap.resize(heap.maxCapacity * 2)
	}

	heap.size++
	heap.data[heap.size] = data
	heap.swim()
}

func (heap *Heap) DeleteTop() interface{} {
	if heap.size == 0 {
		return nil
	}

	data := heap.data[1]
	heap.exch(1, heap.size)
	heap.size--
	heap.sink(1)
	heap.data[heap.size+1] = nil
	if heap.size > 0 && heap.size == (heap.maxCapacity-1)/4 {
		heap.resize(heap.maxCapacity / 2)
	}

	return data
}
