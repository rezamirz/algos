/*

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
package rbt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type IntComparator struct{}

func (ic *IntComparator) Compare(k1, k2 interface{}) int {
	i1, ok := k1.(*int)
	if !ok {
		return 0
	}

	i2, ok := k2.(*int)
	if !ok {
		return 0
	}
	//fmt.Printf("i1=%d, i2=%d\n", *i1, *i2)

	if *i1 > *i2 {
		return 1
	} else if *i1 < *i2 {
		return -1
	}

	return 0
}

type IntDumper struct{}

func (id *IntDumper) Dump(k, v interface{}) {
	fmt.Printf("(%v, %v)\n", k, v)
}

func TestCreateRBT(t *testing.T) {
	rbt := NewRBT(&IntComparator{}, &IntDumper{})
	assert.Equal(t, true, rbt.IsEmpty())
	assert.Equal(t, uint32(0), rbt.Size())
}

func FillRBT(t *testing.T) *RBT {
	numbers := []int{-10, 23, 17, -23, 210, 222, 1000, -100, 0, -55, 730, 731, -110, 1020, 67, 59, -30, 229, 1300, // min = -110, max = 1300
		-200, -210, 900, 1320, 1400, -1000, 1100, -250, 1402, 1403}

	rbt := NewRBT(&IntComparator{}, &IntDumper{})
	assert.Equal(t, true, rbt.IsEmpty())
	assert.Equal(t, uint32(0), rbt.Size())

	// Add -10
	rbt.Put(&numbers[0], &numbers[0])
	assert.Equal(t, uint32(1), rbt.Size())
	n := rbt.Min().(*int)
	assert.Equal(t, -10, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, -10, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 23
	rbt.Put(&numbers[1], &numbers[1])
	assert.Equal(t, uint32(2), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -10, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 23, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 17
	rbt.Put(&numbers[2], &numbers[2])
	assert.Equal(t, uint32(3), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -10, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 23, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add -23
	rbt.Put(&numbers[3], &numbers[3])
	assert.Equal(t, uint32(4), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -23, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 23, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 210
	rbt.Put(&numbers[4], &numbers[4])
	assert.Equal(t, uint32(5), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -23, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 210, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 222
	rbt.Put(&numbers[5], &numbers[5])
	assert.Equal(t, uint32(6), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -23, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 222, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 1000
	rbt.Put(&numbers[6], &numbers[6])
	assert.Equal(t, uint32(7), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -23, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1000, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add -100
	rbt.Put(&numbers[7], &numbers[7])
	assert.Equal(t, uint32(8), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -100, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1000, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 0
	rbt.Put(&numbers[8], &numbers[8])
	assert.Equal(t, uint32(9), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -100, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1000, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add -55
	rbt.Put(&numbers[9], &numbers[9])
	assert.Equal(t, uint32(10), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -100, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1000, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 730
	rbt.Put(&numbers[10], &numbers[10])
	assert.Equal(t, uint32(11), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -100, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1000, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 731
	rbt.Put(&numbers[11], &numbers[11])
	assert.Equal(t, uint32(12), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -100, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1000, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add -110
	rbt.Put(&numbers[12], &numbers[12])
	assert.Equal(t, uint32(13), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -110, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1000, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 1020
	rbt.Put(&numbers[13], &numbers[13])
	assert.Equal(t, uint32(14), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -110, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1020, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 67
	rbt.Put(&numbers[14], &numbers[14])
	assert.Equal(t, uint32(15), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -110, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1020, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 59
	rbt.Put(&numbers[15], &numbers[15])
	assert.Equal(t, uint32(16), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -110, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1020, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add -30
	rbt.Put(&numbers[16], &numbers[16])
	assert.Equal(t, uint32(17), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -110, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1020, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 229
	rbt.Put(&numbers[17], &numbers[17])
	assert.Equal(t, uint32(18), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -110, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1020, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 1300
	rbt.Put(&numbers[18], &numbers[18])
	assert.Equal(t, uint32(19), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -110, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1300, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add -200
	rbt.Put(&numbers[19], &numbers[19])
	assert.Equal(t, uint32(20), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -200, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1300, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add -210
	rbt.Put(&numbers[20], &numbers[20])
	assert.Equal(t, uint32(21), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -210, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1300, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 900
	rbt.Put(&numbers[21], &numbers[21])
	assert.Equal(t, uint32(22), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -210, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1300, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 1320
	rbt.Put(&numbers[22], &numbers[22])
	assert.Equal(t, uint32(23), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -210, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1320, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 1400
	rbt.Put(&numbers[23], &numbers[23])
	assert.Equal(t, uint32(24), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -210, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1400, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add -1000
	rbt.Put(&numbers[24], &numbers[24])
	assert.Equal(t, uint32(25), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -1000, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1400, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 1100
	rbt.Put(&numbers[25], &numbers[25])
	assert.Equal(t, uint32(26), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -1000, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1400, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add -250
	rbt.Put(&numbers[26], &numbers[26])
	assert.Equal(t, uint32(27), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -1000, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1400, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 1402
	rbt.Put(&numbers[27], &numbers[27])
	assert.Equal(t, uint32(28), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -1000, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1402, *n)
	assert.Equal(t, true, rbt.Is23())

	// Add 1403
	rbt.Put(&numbers[28], &numbers[28])
	assert.Equal(t, uint32(29), rbt.Size())
	n = rbt.Min().(*int)
	assert.Equal(t, -1000, *n)
	n = rbt.Max().(*int)
	assert.Equal(t, 1403, *n)
	assert.Equal(t, true, rbt.Is23())

	return rbt
}

func TestRBT_Insert_DeleteMin(t *testing.T) {

	rbt := FillRBT(t)

	rc, key, value := rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, -1000, *(key.(*int)))
	assert.Equal(t, -1000, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, -250, *(key.(*int)))
	assert.Equal(t, -250, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, -210, *(key.(*int)))
	assert.Equal(t, -210, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, -200, *(key.(*int)))
	assert.Equal(t, -200, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, -110, *(key.(*int)))
	assert.Equal(t, -110, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, -100, *(key.(*int)))
	assert.Equal(t, -100, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, -55, *(key.(*int)))
	assert.Equal(t, -55, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, -30, *(key.(*int)))
	assert.Equal(t, -30, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, -23, *(key.(*int)))
	assert.Equal(t, -23, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, -10, *(key.(*int)))
	assert.Equal(t, -10, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 0, *(key.(*int)))
	assert.Equal(t, 0, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 17, *(key.(*int)))
	assert.Equal(t, 17, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 23, *(key.(*int)))
	assert.Equal(t, 23, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 59, *(key.(*int)))
	assert.Equal(t, 59, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 67, *(key.(*int)))
	assert.Equal(t, 67, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 210, *(key.(*int)))
	assert.Equal(t, 210, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 222, *(key.(*int)))
	assert.Equal(t, 222, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 229, *(key.(*int)))
	assert.Equal(t, 229, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 730, *(key.(*int)))
	assert.Equal(t, 730, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 731, *(key.(*int)))
	assert.Equal(t, 731, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 900, *(key.(*int)))
	assert.Equal(t, 900, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1000, *(key.(*int)))
	assert.Equal(t, 1000, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1020, *(key.(*int)))
	assert.Equal(t, 1020, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1100, *(key.(*int)))
	assert.Equal(t, 1100, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1300, *(key.(*int)))
	assert.Equal(t, 1300, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1320, *(key.(*int)))
	assert.Equal(t, 1320, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1400, *(key.(*int)))
	assert.Equal(t, 1400, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1402, *(key.(*int)))
	assert.Equal(t, 1402, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMin()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1403, *(key.(*int)))
	assert.Equal(t, 1403, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())
}

func TestRBT_Insert_DeleteMax(t *testing.T) {

	rbt := FillRBT(t)

	rc, key, value := rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1403, *(key.(*int)))
	assert.Equal(t, 1403, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1402, *(key.(*int)))
	assert.Equal(t, 1402, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1400, *(key.(*int)))
	assert.Equal(t, 1400, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1320, *(key.(*int)))
	assert.Equal(t, 1320, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1300, *(key.(*int)))
	assert.Equal(t, 1300, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1100, *(key.(*int)))
	assert.Equal(t, 1100, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1020, *(key.(*int)))
	assert.Equal(t, 1020, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 1000, *(key.(*int)))
	assert.Equal(t, 1000, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 900, *(key.(*int)))
	assert.Equal(t, 900, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 731, *(key.(*int)))
	assert.Equal(t, 731, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 730, *(key.(*int)))
	assert.Equal(t, 730, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 229, *(key.(*int)))
	assert.Equal(t, 229, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 222, *(key.(*int)))
	assert.Equal(t, 222, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 210, *(key.(*int)))
	assert.Equal(t, 210, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 67, *(key.(*int)))
	assert.Equal(t, 67, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 59, *(key.(*int)))
	assert.Equal(t, 59, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 23, *(key.(*int)))
	assert.Equal(t, 23, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 17, *(key.(*int)))
	assert.Equal(t, 17, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, 0, *(key.(*int)))
	assert.Equal(t, 0, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, -10, *(key.(*int)))
	assert.Equal(t, -10, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, -23, *(key.(*int)))
	assert.Equal(t, -23, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, -30, *(key.(*int)))
	assert.Equal(t, -30, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, -55, *(key.(*int)))
	assert.Equal(t, -55, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, -100, *(key.(*int)))
	assert.Equal(t, -100, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, -110, *(key.(*int)))
	assert.Equal(t, -110, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, -200, *(key.(*int)))
	assert.Equal(t, -200, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, -210, *(key.(*int)))
	assert.Equal(t, -210, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, -250, *(key.(*int)))
	assert.Equal(t, -250, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())

	rc, key, value = rbt.DeleteMax()
	assert.Equal(t, true, rc)
	assert.Equal(t, -1000, *(key.(*int)))
	assert.Equal(t, -1000, *(value.(*int)))
	assert.Equal(t, true, rbt.Is23())
}

func TestRBT_Insert(t *testing.T) {
	keys := []int{67, 1320, 1000, -200, 59, -100, 731, -23, 900, -1000, 1100, 1020, 23, 0, -250, -55, 210, 1300,
		-30, -10, 1403, -110, 17, 730, -210, 229, 1400, 222, 1402}

	rbt := FillRBT(t)

	for i := 0; i < len(keys); i++ {
		rc := rbt.Delete(&keys[i])
		assert.Equal(t, true, rc)
		assert.Equal(t, true, rbt.Is23())
	}
}

const MaxNumbers = 10000

func TestRBT_Insert_Delete_In_Loop(t *testing.T) {

	numbers := make([]int, MaxNumbers)
	for i := 0; i < MaxNumbers; i++ {
		numbers[i] = i
	}

	rbt := NewRBT(&IntComparator{}, &IntDumper{})
	assert.Equal(t, true, rbt.IsEmpty())
	assert.Equal(t, uint32(0), rbt.Size())
	assert.Equal(t, true, rbt.Is23())

	// Insert numbers from smallest to biggest and delete them
	{
		for i := 0; i < MaxNumbers; i++ {
			rbt.Put(&numbers[i], &numbers[i])
			assert.Equal(t, uint32(i+1), rbt.Size())
			n := rbt.Min().(*int)
			assert.Equal(t, 0, *n)
			n = rbt.Max().(*int)
			assert.Equal(t, i, *n)
			assert.Equal(t, true, rbt.Is23())
		}

		for i := 0; i < MaxNumbers; i++ {
			key := i
			r := rbt.Rank(&key)
			assert.Equal(t, uint32(i), r)
		}

		for i := 0; i < MaxNumbers; i++ {
			key := rbt.Select(uint32(i))
			assert.Equal(t, i, *(key.(*int)))
		}

		for i := 0; i < MaxNumbers; i++ {
			key := i
			rc := rbt.Delete(&key)
			assert.Equal(t, true, rc)
			assert.Equal(t, true, rbt.Is23())

			if i != MaxNumbers-1 {
				key2 := rbt.Ceiling(&key)
				assert.Equal(t, key+1, *(key2.(*int)))
			}
		}
	}

	{
		for i := 0; i < MaxNumbers; i++ {
			rbt.Put(&numbers[i], &numbers[i])
			assert.Equal(t, uint32(i+1), rbt.Size())
			n := rbt.Min().(*int)
			assert.Equal(t, 0, *n)
			n = rbt.Max().(*int)
			assert.Equal(t, i, *n)
			assert.Equal(t, true, rbt.Is23())
		}

		for i := 0; i < MaxNumbers; i++ {
			key := i
			r := rbt.Rank(&key)
			assert.Equal(t, uint32(i), r)
		}

		for i := MaxNumbers - 1; i >= 0; i-- {
			key := i
			rc := rbt.Delete(&key)
			assert.Equal(t, true, rc)
			assert.Equal(t, true, rbt.Is23())

			if i != 0 {
				key2 := rbt.Floor(&key)
				assert.Equal(t, key-1, *(key2.(*int)))

			}
		}
	}
}
