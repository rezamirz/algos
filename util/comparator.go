/*

comparator.go

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

import "fmt"

type Comparator interface {
	/*
	 * Compares two keys and returns -1, 0, or 1.
	 * The return value is used to sort the keys.
	 */
	Compare(k1, k2 interface{}) int
}

type IntComparator struct{}

func (ic *IntComparator) Compare(k1, k2 interface{}) int {
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

type Int64Greater struct{}

func (g *Int64Greater) Compare(k1, k2 interface{}) int {
	p1, ok := k1.(int64)
	if !ok {
		panic(fmt.Errorf("Type assersion failed in Int64Geater %T", k1))
	}

	p2, ok := k2.(int64)
	if !ok {
		panic(fmt.Errorf("Type assersion failed in Int64Greater %T", k2))
	}

	if p1 > p2 {
		return 1
	}

	if p1 == p2 {
		return 0
	}

	return -1
}

type Int64Smaller struct{}

func (s *Int64Smaller) Compare(k1, k2 interface{}) int {
	p1, ok := k1.(int64)
	if !ok {
		panic(fmt.Errorf("Type assersion failed in Int64Smaller %T", k1))
	}

	p2, ok := k2.(int64)
	if !ok {
		panic(fmt.Errorf("Type assersion failed in Int64Smaller %T", k2))
	}

	if p1 < p2 {
		return 1
	}

	if p1 == p2 {
		return 0
	}

	return -1
}
