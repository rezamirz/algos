package pq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	MAX_NUMBERS = 10000
)

type MyIntComparator struct{}

func (ic *MyIntComparator) Compare(k1, k2 interface{}) int {
	i1, ok := k1.(*int)
	if !ok {
		panic(fmt.Errorf("Type assersion failed in IntComparator"))
	}

	i2, ok := k2.(*int)
	if !ok {
		panic(fmt.Errorf("Type assersion failed in IntComparator"))
	}

	if *i1 > *i2 {
		return 1
	} else if *i1 < *i2 {
		return -1
	}

	return 0
}

func TestCreateTopheap(t *testing.T) {
	heap := NewHeap(10, &MyIntComparator{})
	assert.Equal(t, true, heap.IsEmpty())
	assert.Equal(t, uint32(0), heap.Size())
}

func TestInsertDeleteTopheap_1(t *testing.T) {
	numbers := []int{-10, 23, 17, -23, 210, 222, 1000, -100, 0, -55}

	heap := NewHeap(10, &MyIntComparator{})
	assert.Equal(t, true, heap.IsEmpty())
	assert.Equal(t, uint32(0), heap.Size())

	heap.Put(&numbers[0])
	assert.Equal(t, uint32(1), heap.Size())
	n := heap.Top().(*int)
	assert.Equal(t, -10, *n)

	heap.Put(&numbers[1])
	assert.Equal(t, uint32(2), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, -10, *n)

	heap.Put(&numbers[2])
	assert.Equal(t, uint32(3), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, -10, *n)

	heap.Put(&numbers[3])
	assert.Equal(t, uint32(4), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, -23, *n)

	heap.Put(&numbers[4])
	assert.Equal(t, uint32(5), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, -23, *n)

	heap.Put(&numbers[5])
	assert.Equal(t, uint32(6), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, -23, *n)

	heap.Put(&numbers[6])
	assert.Equal(t, uint32(7), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, -23, *n)

	heap.Put(&numbers[7])
	assert.Equal(t, uint32(8), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, -100, *n)

	heap.Put(&numbers[8])
	assert.Equal(t, uint32(9), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, -100, *n)

	heap.Put(&numbers[9])
	assert.Equal(t, uint32(10), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, -100, *n)

	n = heap.DeleteTop().(*int)
	assert.Equal(t, -100, *n)
	assert.Equal(t, uint32(9), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, -55, *n)

	n = heap.DeleteTop().(*int)
	assert.Equal(t, -55, *n)
	assert.Equal(t, uint32(8), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, -23, *n)

	n = heap.DeleteTop().(*int)
	assert.Equal(t, -23, *n)
	assert.Equal(t, uint32(7), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, -10, *n)

	n = heap.DeleteTop().(*int)
	assert.Equal(t, -10, *n)
	assert.Equal(t, uint32(6), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, 0, *n)

	n = heap.DeleteTop().(*int)
	assert.Equal(t, 0, *n)
	assert.Equal(t, uint32(5), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, 17, *n)

	n = heap.DeleteTop().(*int)
	assert.Equal(t, 17, *n)
	assert.Equal(t, uint32(4), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, 23, *n)

	n = heap.DeleteTop().(*int)
	assert.Equal(t, 23, *n)
	assert.Equal(t, uint32(3), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, 210, *n)

	n = heap.DeleteTop().(*int)
	assert.Equal(t, 210, *n)
	assert.Equal(t, uint32(2), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, 222, *n)

	n = heap.DeleteTop().(*int)
	assert.Equal(t, 222, *n)
	assert.Equal(t, uint32(1), heap.Size())
	n = heap.Top().(*int)
	assert.Equal(t, 1000, *n)

	n = heap.DeleteTop().(*int)
	assert.Equal(t, 1000, *n)
	assert.Equal(t, uint32(0), heap.Size())
	last := heap.Top()
	assert.Equal(t, nil, last)
}

func TestInsertDeleteTopheap_2(t *testing.T) {
	heap := NewHeap(10, &MyIntComparator{})
	assert.Equal(t, true, heap.IsEmpty())
	assert.Equal(t, uint32(0), heap.Size())
	assert.Equal(t, true, heap.IsEmpty())

	numbers := make([]int, MAX_NUMBERS)
	for i := 0; i < MAX_NUMBERS; i++ {
		if i%2 == 0 {
			numbers[i] = i
		} else {
			numbers[i] = -i
		}
		heap.Put(&numbers[i])
		assert.Equal(t, uint32(i+1), heap.Size())
	}

	for i := 0; i < MAX_NUMBERS; i++ {
		n := heap.DeleteTop().(*int)
		if i < MAX_NUMBERS/2 {
			assert.Equal(t, -MAX_NUMBERS+2*i+1, *n)
		} else {
			assert.Equal(t, -MAX_NUMBERS+2*i, *n)
		}
		assert.Equal(t, uint32(MAX_NUMBERS-i-1), heap.Size())
	}

}
