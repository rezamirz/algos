/*

graph.go

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

type Graph interface {
	// Returns number of vertices
	GetNumVertices() int

	// Returns number of edges
	GetNumEdges() int

	// Returns the index of newly added vertex
	AddVertex() int

	// Returns true if the graph has the vertex v
	HasVertex(v int) bool

	// Add an edge to the graph
	AddEdge(v, w int)

	// Returns an array containing all the neighbors of v
	GetNeighbors(v int) []int
}

// Directed graph
type DGraph struct {
	// Adjacency map of the graph nodes
	adjMap map[int][]int

	nVertices int
	nEdges    int
}

func NewDGraph(nVertices int) Graph {
	g := &DGraph{
		adjMap:    map[int][]int{},
		nVertices: nVertices,
	}

	return g
}

func (dg *DGraph) GetNumVertices() int {
	return dg.nVertices
}

func (dg *DGraph) GetNumEdges() int {
	return dg.nEdges
}

func (dg *DGraph) AddVertex() int {
	dg.nVertices++
	return dg.nVertices
}

func (dg *DGraph) HasVertex(v int) bool {
	return v > 0 && v < dg.nVertices
}

func (dg *DGraph) AddEdge(v, w int) {
	neighbors, _ := dg.adjMap[v]
	neighbors = append(neighbors, w)
	//fmt.Printf("AddEdge v=%d, neighbors=%v\n", v, neighbors)
	dg.adjMap[v] = neighbors
	dg.nEdges++
}

func (dg *DGraph) GetNeighbors(v int) []int {
	neighbors, ok := dg.adjMap[v]
	if !ok {
		return nil
	}

	copyNeighbors := make([]int, len(neighbors))
	copy(copyNeighbors, neighbors)
	return copyNeighbors
}
