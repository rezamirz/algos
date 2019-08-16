package pq

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
	Get(key string) (int64, error)

	// If the key doesn't exist in the PQ, it inserts new key with priority.
	// If the key exist in PQ then it updates the key priority.
	// The key might move to new place in the PQ based on new priority.
	// Returns true if the PQ contains the key, otherwise on new insertions returns false.
	Put(key string, priority int64) (bool, error)

	Delete(key string) error

	// Returns first key (with min/max priority), it's priority and error if there is any
	First() (string, int64, error)

	// Pops first key (with min/max priority) from the queue, and returns the key it's priority, and error if there is any
	Dequeue() (string, int64, error)
}

type Comparator interface {
	/*
	 * Compares two keys and returns -1, 0, or 1.
	 * The return value is used to sort the keys.
	 */
	Compare(i, j int64) int
}

type greater struct{}

func (g *greater) Compare(p1, p2 int64) int {
	if p1 > p2 {
		return 1
	}

	if p1 == p2 {
		return 0
	}

	return -1
}

type smaller struct{}

func (s *smaller) Compare(p1, p2 int64) int {
	if p1 < p2 {
		return 1
	}

	if p1 == p2 {
		return 0
	}

	return -1
}

// The elements inside a HashedPQ
type heapElement struct {
	index int // Index in heap array
	key   string
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
	table    map[string]*heapElement
	n        int // Number of elements in PQ
	capacity int

	comparator Comparator
}

func NewHashedPQ(pqType PQType, capacity int) PriorityQueue {
	heap := &HashedPQ{
		pqType:   pqType,
		a:        make([]*heapElement, capacity),
		table:    map[string]*heapElement{},
		n:        0,
		capacity: capacity,
	}

	if pqType == MinPQ {
		heap.comparator = &greater{}
	} else {
		heap.comparator = &smaller{}
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
func (pq *HashedPQ) Get(key string) (int64, error) {
	element, ok := pq.table[key]
	if !ok {
		return 0, fmt.Errorf("GET key '%s', not found", key)
	}

	if element == nil {
		return 0, fmt.Errorf("GET key '%s', nil heap element", key)
	}

	return element.p, nil
}

func (pq *HashedPQ) Put(key string, priority int64) (bool, error) {
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

func (pq *HashedPQ) Delete(key string) error {
	element, ok := pq.table[key]
	if !ok {
		return fmt.Errorf("UPDATE key %s, not found", key)
	}

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

func (pq *HashedPQ) First() (string, int64, error) {
	if pq.n < 1 {
		return "", 0, fmt.Errorf("FIRST empty heap")
	}

	element := pq.a[1]
	return element.key, element.p, nil
}

func (pq *HashedPQ) Dequeue() (string, int64, error) {
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

	pq.dump("DEQUEUE", element)
	return element.key, element.p, nil

}

func (pq *HashedPQ) dump(op string, element *heapElement) {
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
