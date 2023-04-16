package search

import "github.com/mispon/grokking_algorithms/internal/types"

// BFS is Breadth-First Search implementation
func BFS[T comparable](graph *types.GraphNode[T], target T) *types.GraphNode[T] {
	if graph == nil {
		return nil
	}

	if graph.Value == target {
		return graph
	}

	var (
		q    = types.NewQueue[*types.GraphNode[T]](len(graph.Neighbors))
		seen = map[*types.GraphNode[T]]struct{}{}
	)

	for _, n := range graph.Neighbors {
		q.Enqueue(n)
	}

	for {
		node, ok := q.Dequeue()
		if !ok {
			break
		}

		if _, ok = seen[node]; ok {
			continue
		}
		seen[node] = struct{}{}

		if node.Value == target {
			return node
		}
		
		for _, n := range node.Neighbors {
			q.Enqueue(n)
		}
	}

	return nil
}
