package graph

import "github.com/mispon/grokking_algorithms/internal/types"

// DFS is Depth-First Search implementation
func DFS[T comparable](graph *types.GraphNode[T], target T) *types.GraphNode[T] {
	if graph == nil {
		return nil
	}

	if graph.Value == target {
		return graph
	}

	for _, node := range graph.Neighbors {
		if n := DFS(node, target); n != nil {
			return n
		}
	}

	return nil
}
