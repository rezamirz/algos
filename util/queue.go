/*

queue.go

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
	"container/list"
)

type Queue struct {
	list  *list.List
}

func NewQueue() *Queue {
	q := &Queue{
		list:  list.New(),
	}

	return q
}

func (q *Queue) Front() interface{} {
	e := q.list.Front()
	if e == nil {
		return nil
	}

	return e.Value
}

func (q *Queue) Pop() interface{} {
	e := q.list.Front()
	if e == nil {
		return nil
	}

	q.list.Remove(e)
	return e.Value
}

func (q *Queue) Push(v interface{}) {
	q.list.PushBack(v)
}

func (q *Queue) Len() int {
	return q.list.Len()
}
