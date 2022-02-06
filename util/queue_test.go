package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFifo(t *testing.T) {
	q := New()

	q.Push(1)
	q.Push(2)
	q.Push(3)
	assert.Equal(t, 3, q.Len())

	v := q.Pop()
	assert.Equal(t, 1, v)
	v = q.Pop()
	assert.Equal(t, 2, v)
	v = q.Pop()
	assert.Equal(t, 3, v)
	assert.Equal(t, 0, q.Len())

	for i:=1; i<=100; i++ {
		q.Push(i)
	}
	assert.Equal(t, 100, q.Len())

	for i:=1; i<=100; i++ {
		v = q.Pop()
		assert.Equal(t, i, v)
	}
	assert.Equal(t, 0, q.Len())
}

func TestFifo2(t *testing.T) {
	q := New()

	q.Push(1)
	assert.Equal(t, 1, q.Len())

	v := q.Pop()
	assert.Equal(t, 1, v)
	assert.Equal(t, 0, q.Len())
}
