package graph

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Load(filename string) (Graph, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	lineNo := 0
	nVertices := 0

	var g Graph

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		s := scanner.Text()
		lineNo++
		switch lineNo {
		case 1:
			nVertices, err = strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			g = NewDGraph(nVertices)
		case 2:
			// Read nEdges
			_, err = strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
		default:
			edges := strings.Split(s, " ")
			v, err := strconv.Atoi(edges[0])
			if err != nil {
				return nil, err
			}

			w, err := strconv.Atoi(edges[1])
			if err != nil {
				return nil, err
			}

			g.AddEdge(v, w)
		}
	}
	//fmt.Printf("Load V=%d g.V=%d, E=%d, g.E=%d\n", nVertices, g.GetNumVertices(), nEdges, g.GetNumEdges())
	return g, nil
}

func TestDFS_SmallGraph(t *testing.T) {
	g, err := Load("data/tinyG.txt")
	assert.NoError(t, err)
	assert.Equal(t, 6, g.GetNumVertices())
	assert.Equal(t, 8, g.GetNumEdges())

	dfs := NewSearch(DepthFirstSearch)
	assert.NotEqual(t, nil, dfs)
	dfs.DoSearch(g, 0, 3)
	path := dfs.PathTo(3)
	assert.Equal(t, 4, len(path))
}


func TestBFS_SmallGraph(t *testing.T) {
	g, err := Load("data/tinyG.txt")
	assert.NoError(t, err)
	assert.Equal(t, 6, g.GetNumVertices())
	assert.Equal(t, 8, g.GetNumEdges())

	bfs := NewSearch(BreathFirstSearch)
	assert.NotEqual(t, nil, bfs)
	bfs.DoSearch(g, 0, 3)
	path := bfs.PathTo(3)
	assert.Equal(t, 3, len(path))
}
