/*

union_find.go

An implementation of compressed union find.

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

import algo_error "github.com/rezamirz/myalgos/error"

type UnionFind interface {
	Count() int                       // Return the number of disjoint sets
	Find(p int) (int, error)          // Return component identifier of p
	Connected(p, q int) (bool, error) // Checks if p and q are connected
	Union(p, q int) error             // Connects p and q
}

type weightedUnionFind struct {
	id    []int
	size  []int
	count int
}

func NewUF(n int) UnionFind {
	uf := &weightedUnionFind{
		id:    make([]int, n),
		size:  make([]int, n),
		count: n,
	}

	for i := 0; i < n; i++ {
		uf.id[i] = i
		uf.size[i] = 1
	}

	return uf
}

func (uf *weightedUnionFind) Count() int {
	return uf.count
}

func (uf *weightedUnionFind) Find(p int) (int, error) {
	if p < 0 || p >= len(uf.id) {
		return -1, algo_error.INVALID_ARGUMENT
	}

	for p != uf.id[p] {
		p = uf.id[p]
	}

	return p, nil
}

func (uf *weightedUnionFind) Connected(p, q int) (bool, error) {
	i, err := uf.Find(p)
	if err != nil {
		return false, err
	}
	j, err := uf.Find(q)
	if err != nil {
		return false, err
	}
	return i == j, nil
}

func (uf *weightedUnionFind) Union(p, q int) error {
	i, err := uf.Find(p)
	if err != nil {
		return err
	}

	j, err := uf.Find(q)
	if err != nil {
		return err
	}

	// They are already connected
	if i == j {
		return nil
	}

	if uf.size[i] < uf.size[j] {
		uf.id[i] = j
		uf.size[j] += uf.size[i]
	} else {
		uf.id[j] = i
		uf.size[i] += uf.size[j]
	}

	uf.count--
	return nil
}
