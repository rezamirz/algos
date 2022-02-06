package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDequeFifo(t *testing.T) {
	q := NewDeque()

	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(3)
	assert.Equal(t, 3, q.Len())

	v := q.PopFront()
	assert.Equal(t, 1, v)
	v = q.PopFront()
	assert.Equal(t, 2, v)
	v = q.PopFront()
	assert.Equal(t, 3, v)
	assert.Equal(t, 0, q.Len())

	for i:=1; i<=100; i++ {
		q.PushBack(i)
	}
	assert.Equal(t, 100, q.Len())

	for i:=1; i<=100; i++ {
		v = q.PopFront()
		assert.Equal(t, i, v)
	}
	assert.Equal(t, 0, q.Len())
}

func TestDequeFilo(t *testing.T) {
	q := NewDeque()

	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(3)
	assert.Equal(t, 3, q.Len())

	v := q.PopBack()
	assert.Equal(t, 3, v)
	v = q.PopBack()
	assert.Equal(t, 2, v)
	v = q.PopBack()
	assert.Equal(t, 1, v)
	assert.Equal(t, 0, q.Len())

	for i:=1; i<=100; i++ {
		q.PushBack(i)
	}
	assert.Equal(t, 100, q.Len())

	for i:=100; i>=1; i-- {
		v = q.PopBack()
		assert.Equal(t, i, v)
	}

	assert.Equal(t, 0, q.Len())
}

func TestDequeFilo2(t *testing.T) {
	q := NewDeque()

	q.PushBack(1)
	assert.Equal(t, 1, q.Len())

	v := q.PopBack()
	assert.Equal(t, 1, v)
	assert.Equal(t, 0, q.Len())
}