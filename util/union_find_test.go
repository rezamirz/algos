package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUF(t *testing.T) {
	uf := NewUF(11)

	p, err := uf.Find(10)
	assert.Equal(t, nil, err)
	assert.Equal(t, 10, p)

	isConnected, err := uf.Connected(9, 10)
	assert.NoError(t, err)
	assert.Equal(t, false, isConnected)
	assert.Equal(t, 11, uf.Count())

	uf.Union(9, 10)
	isConnected, err = uf.Connected(9, 10)
	assert.NoError(t, err)
	assert.Equal(t, true, isConnected)
	assert.Equal(t, 10, uf.Count())

	p, err = uf.Find(9)
	assert.NoError(t, err)
	q, err := uf.Find(10)
	assert.NoError(t, err)
	assert.Equal(t, true, p == q)

	uf.Union(1, 2)
	uf.Union(4, 7)
	uf.Union(1, 7)
	p, err = uf.Find(2)
	assert.NoError(t, err)
	q, err = uf.Find(7)
	assert.NoError(t, err)
	assert.Equal(t, true, p == q)
	assert.Equal(t, 7, uf.Count())
}
