package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()

	s.Push(1)
	v := s.Top()
	assert.Equal(t, 1, v)
	s.Push(2)
	v = s.Top()
	assert.Equal(t, 2, v)
	s.Push(3)
	v = s.Top()
	assert.Equal(t, 3, v)
	assert.Equal(t, 3, s.Len())

	v = s.Pop()
	assert.Equal(t, 3, v)
	v = s.Pop()
	assert.Equal(t, 2, v)
	v = s.Pop()
	assert.Equal(t, 1, v)
	assert.Equal(t, 0, s.Len())

	for i:=1; i<=100; i++ {
		s.Push(i)
		v = s.Top()
		assert.Equal(t, i, v)
	}
	assert.Equal(t, 100, s.Len())

	for i:=1; i<=100; i++ {
		v = s.Pop()
		assert.Equal(t, 101-i, v)
	}
	assert.Equal(t, 0, s.Len())
}

func TestStack2(t *testing.T) {
	s := NewStack()

	s.Push(1)
	assert.Equal(t, 1, s.Len())

	v := s.Pop()
	assert.Equal(t, 1, v)
	assert.Equal(t, 0, s.Len())
}