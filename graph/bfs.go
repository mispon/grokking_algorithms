package graph

import "github.com/mispon/grokking_algorithms/internal/types"

// BFS is Breadth-First Search implementation
func BFS[T comparable](graph *types.GraphNode[T], target T) *types.GraphNode[T] {
	if graph == nil {
		return nil
	}

	var (
		queue = types.NewQueue[*types.GraphNode[T]](len(graph.Neighbors))
		seen  = map[*types.GraphNode[T]]struct{}{}
	)

	queue.Enqueue(graph)

	for {
		node, ok := queue.Dequeue()
		if !ok {
			break
		}

		if _, ok = seen[node]; ok {
			continue
		}
		seen[node] = struct{}{}

		switch node.Value {
		case target:
			return node
		default:
			for _, n := range node.Neighbors {
				queue.Enqueue(n)
			}
		}
	}

	return nil
}
