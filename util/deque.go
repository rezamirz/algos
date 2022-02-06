/*

deque.go

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
	"sync"
)

type Deque struct {
	list  *list.List
	cond  *sync.Cond
	mutex *sync.RWMutex
}

func NewDequeu() *Deque {
	q := &Deque{
		list:  list.New(),
		mutex: &sync.RWMutex{},
	}

	q.cond = sync.NewCond(q.mutex)

	return q
}

func (q *Deque) PushFront(v interface{}) *list.Element {
	q.mutex.Lock()
	e := q.list.PushFront(v)
	q.cond.Broadcast()
	q.mutex.Unlock()

	return e
}

func (q *Deque) PopFront() interface{} {
	var e *list.Element

	q.mutex.Lock()
	for {
		e = q.list.Front()
		if e != nil {
			break
		}

		q.cond.Wait()
	}

	q.list.Remove(e)
	q.cond.Broadcast()
	q.mutex.Unlock()

	return e.Value
}

func (q *Deque) PushBack(v interface{}) *list.Element {
	q.mutex.Lock()
	e := q.list.PushBack(v)
	q.cond.Broadcast()
	q.mutex.Unlock()

	return e
}

func (q *Deque) PopBack() interface{} {
	var e *list.Element

	q.mutex.Lock()
	for {
		e = q.list.Back()
		if e != nil {
			break
		}

		q.cond.Wait()
	}

	q.list.Remove(e)
	q.cond.Broadcast()
	q.mutex.Unlock()

	return e.Value
}

func (q *Deque) Len() int {
	return q.list.Len()
}
