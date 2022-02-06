/*

search.go

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

package graph

import (
	"github.com/rezamirz/myalgos/util"
)

type Search interface {
	DoSearch(g Graph, start, end int) // Search from start to end
	Count() int                       // Number of connected nodes to start
	PathTo(v int) []int               // Path to node v from start
}

type SearchType int

const (
	Uknown SearchType = iota
	DepthFirstSearch
	BreathFirstSearch
)

func NewSearch(searchType SearchType) Search {
	switch searchType {
	case DepthFirstSearch:
		return &DFS{}
	case BreathFirstSearch:
		return &BFS{}
	}

	return nil
}

type DFS struct {
	start  int    // Start node for the search
	count  int    // Number of nodes connected to the start
	marked []bool // is the node already marked / visited
	pathTo []int  // Path to node v from start
}

func (dfs *DFS) DoSearch(g Graph, start, end int) {
	dfs.marked = make([]bool, g.GetNumVertices())
	dfs.pathTo = make([]int, g.GetNumVertices())
	dfs.start = start
	dfs.doSearch(g, start, end)
}

func (dfs *DFS) doSearch(g Graph, start, end int) {
	dfs.count++
	dfs.marked[start] = true

	for _, v := range g.GetNeighbors(start) {
		if !dfs.marked[v] {
			dfs.pathTo[v] = start
			dfs.doSearch(g, v, end)
		}
	}
}

func (dfs *DFS) Count() int {
	return dfs.count
}

func (dfs *DFS) PathTo(v int) []int {
	if !dfs.marked[v] {
		return nil
	}

	path := []int{}
	for {
		path = append(path, v)
		if v == dfs.start {
			break
		}
		v = dfs.pathTo[v]
	}

	reverse(path)
	return path
}

type BFS struct {
	q *util.Queue
	start  int    // Start node for the search
	count  int    // Number of nodes connected to the start
	marked []bool // is the node already marked / visited
	pathTo []int  // Path to node v from start
}

func (bfs *BFS) DoSearch(g Graph, start, end int) {
	bfs.start = start
	bfs.q = util.NewQueue()
	bfs.marked = make([]bool, g.GetNumVertices())
	bfs.pathTo = make([]int, g.GetNumVertices())

	bfs.q.Push(start)
	for bfs.q.Len() > 0 {
		v := bfs.q.Pop().(int)
		neighbors := g.GetNeighbors(v)
		for _, w := range neighbors {
			if bfs.marked[w] {
				continue
			}
			bfs.count++
			bfs.marked[w] = true
			bfs.pathTo[w] = v
			bfs.q.Push(w)
		}
	}
}

func (bfs *BFS) Count() int {
	return bfs.count
}

func (bfs *BFS) PathTo(v int) []int {
	if !bfs.marked[v] {
		return nil
	}

	path := []int{}
	for {
		path = append(path, v)
		if v == bfs.start {
			break
		}
		v = bfs.pathTo[v]
	}

	reverse(path)
	return path
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
