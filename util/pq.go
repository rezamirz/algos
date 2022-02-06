/*

pq.go

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

package util

import (
	"fmt"
	"os"
)

const debug = false

type PriorityQueue interface {
	IsEmpty() bool
	Size() int

	Capacity() int

	// Returns priority of the key and error
	Get(key interface{}) (int64, error)

	// If the key doesn't exist in the PQ, it inserts new key with priority.
	// If the key exist in PQ then it updates the key priority.
	// The key might move to new place in the PQ based on new priority.
	// Returns true if the PQ contains the key, otherwise on new insertions returns false.
	Put(key interface{}, priority int64) (bool, error)

	Delete(key interface{}) error

	// Returns first key (with min/max priority), it's priority and error if there is any
	First() (interface{}, int64, error)

	// Pops first key (with min/max priority) from the queue, and returns the key it's priority, and error if there is any
	Dequeue() (interface{}, int64, error)
}

// The elements inside a HashedPQ
type heapElement struct {
	index int // Index in heap array
	key   interface{}
	p     int64 // Priority
}

type PQType int

const (
	MinPQ PQType = 1
	MaxPQ PQType = 2
)

// An implementation of PriorityQueue interface with heap and hash
type HashedPQ struct {
	pqType   PQType
	a        []*heapElement
	table    map[interface{}]*heapElement
	n        int // Number of elements in PQ
	capacity int

	comparator Comparator
}

func NewHashedPQ(pqType PQType, capacity int) PriorityQueue {
	heap := &HashedPQ{
		pqType:   pqType,
		a:        make([]*heapElement, capacity),
		table:    map[interface{}]*heapElement{},
		n:        0,
		capacity: capacity,
	}

	if pqType == MinPQ {
		heap.comparator = &Int64Greater{}
	} else {
		heap.comparator = &Int64Smaller{}
	}

	return heap
}

func (pq *HashedPQ) IsEmpty() bool {
	return pq.n <= 0
}

func (pq *HashedPQ) Size() int {
	return pq.n
}

func (pq *HashedPQ) Capacity() int {
	return pq.capacity
}

// Returns priority and error if there is any
func (pq *HashedPQ) Get(key interface{}) (int64, error) {
	element, ok := pq.table[key]
	if !ok {
		return 0, fmt.Errorf("GET key '%s', not found", key)
	}

	if element == nil {
		return 0, fmt.Errorf("GET key '%s', nil heap element", key)
	}

	return element.p, nil
}

func (pq *HashedPQ) Put(key interface{}, priority int64) (bool, error) {
	element, ok := pq.table[key]
	if ok {
		pq.update(element, priority)
		return true, nil
	}

	pq.n++
	element = &heapElement{
		index: pq.n,
		key:   key,
		p:     priority,
	}

	pq.table[key] = element

	if pq.capacity == pq.n {
		pq.a = append(pq.a, element)
		pq.capacity++
	} else {
		pq.a[pq.n] = element
	}

	pq.swim(pq.n)
	pq.dump("PUT", element)
	return false, nil
}

func (pq *HashedPQ) swim(k int) {
	for k > 1 && pq.comparator.Compare(pq.a[k/2].p, pq.a[k].p) > 0 {
		pq.exch(k, k/2)
		k = k / 2
	}
}

func (pq *HashedPQ) sink(k int) {
	for 2*k <= pq.n {
		j := 2 * k

		if j < pq.n && pq.comparator.Compare(pq.a[j].p, pq.a[j+1].p) > 0 {
			j++
		}

		if pq.comparator.Compare(pq.a[k].p, pq.a[j].p) <= 0 {
			break
		}

		pq.exch(k, j)
		k = j
	}
}

func (pq *HashedPQ) exch(i, j int) {
	tmp := pq.a[i]
	pq.a[i] = pq.a[j]
	pq.a[i].index = i
	pq.a[j] = tmp
	pq.a[j].index = j
}

func (pq *HashedPQ) update(element *heapElement, priority int64) {

	oldPriority := element.p
	element.p = priority
	if priority > oldPriority {
		pq.sink(element.index)
	} else if priority < oldPriority {
		pq.swim(element.index)
	}
}

func (pq *HashedPQ) Delete(key interface{}) error {
	element, ok := pq.table[key]
	if !ok {
		return fmt.Errorf("UPDATE key %s, not found", key)
	}

	delete(pq.table, key)
	oldPriority := element.p
	index := element.index
	lastElement := pq.a[pq.n]
	pq.a[index] = lastElement
	pq.n--
	lastElement.index = index
	if lastElement.p > oldPriority {
		pq.sink(index)
	} else if lastElement.p < oldPriority {
		pq.swim(index)
	}

	pq.dump("DELETE", element)
	return nil
}

func (pq *HashedPQ) First() (interface{}, int64, error) {
	if pq.n < 1 {
		return "", 0, fmt.Errorf("FIRST empty heap")
	}

	element := pq.a[1]
	return element.key, element.p, nil
}

func (pq *HashedPQ) Dequeue() (interface{}, int64, error) {
	if pq.n <= 0 {
		return "", 0, fmt.Errorf("DEQUEUE empty heap")
	}

	element := pq.a[1]
	index := element.index
	if index != 1 {
		panic("DEQUEUE root index != 1")
	}
	lastElement := pq.a[pq.n]

	pq.a[index] = lastElement
	pq.n--
	lastElement.index = index
	pq.sink(index)

	delete(pq.table, element.key)

	pq.dump("DEQUEUE", element)
	return element.key, element.p, nil

}

func (pq *HashedPQ) dump(op interface{}, element *heapElement) {
	if !debug {
		return
	}

	fmt.Fprintf(os.Stdout, "Dump %s, key=%s, p=%d, capacity=%d, n=%d, index=%d\n",
		op, element.key, element.p, pq.capacity, pq.n, element.index)
	for i := 1; i <= int(pq.n); i++ {
		element := pq.a[i]
		fmt.Fprintf(os.Stdout, "pq[%d]=%s, priority=%d\n",
			i, element.key, element.p)
	}
}
