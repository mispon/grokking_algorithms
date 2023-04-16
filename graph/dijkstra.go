package graph

import (
	"fmt"
	"math"
	"strings"

	"github.com/mispon/grokking_algorithms/internal/types"
)

// Dijkstra is the cheapest path algorith implementation
func Dijkstra[T comparable](graph *types.GraphNode[T], target T) string {
	if graph == nil {
		return ""
	}

	var (
		parents = map[*types.GraphNode[T]]*types.GraphNode[T]{}
		costs   = map[*types.GraphNode[T]]uint{}
		seen    = map[*types.GraphNode[T]]struct{}{}
	)

	calcGraphCosts(graph, costs, 0)
	for _, n := range graph.Neighbors {
		parents[n] = graph
	}

	node := lowerCostNode(costs, seen)
	for node != nil {
		cost := costs[node]
		for _, nn := range node.Neighbors {
			newCost := cost + node.Costs[nn]
			if newCost < costs[nn] {
				costs[nn] = newCost
				parents[nn] = node
			}
		}
		seen[node] = struct{}{}

		node = lowerCostNode(costs, seen)
	}

	return getPath(parents, target)
}

func calcGraphCosts[T comparable](node *types.GraphNode[T], costs map[*types.GraphNode[T]]uint, depth int) {
	for n, c := range node.Costs {
		// we know only neighbors costs at the beginning,
		// other costs are inf.
		if depth == 0 {
			costs[n] = c
		} else {
			costs[n] = math.MaxUint
		}
		calcGraphCosts(n, costs, depth+1)
	}
}

func lowerCostNode[T comparable](
	costs map[*types.GraphNode[T]]uint,
	seen map[*types.GraphNode[T]]struct{},
) *types.GraphNode[T] {
	lowerCost := uint(math.MaxUint)
	var result *types.GraphNode[T]

	for node, cost := range costs {
		if _, ok := seen[node]; ok {
			continue
		}

		if cost < lowerCost {
			lowerCost = cost
			result = node
		}
	}

	return result
}

func getPath[T comparable](
	parents map[*types.GraphNode[T]]*types.GraphNode[T],
	target T,
) string {
	var node *types.GraphNode[T]
	for n := range parents {
		if n.Value == target {
			node = n
		}
	}

	if node == nil {
		return ""
	}

	var (
		steps = []T{node.Value}
		ok    = true
	)
	for {
		node, ok = parents[node]
		if !ok {
			break
		}
		steps = append(steps, node.Value)
	}

	path := ""
	for i := len(steps) - 1; i >= 0; i-- {
		path += fmt.Sprintf("%v -> ", steps[i])
	}

	return strings.TrimRight(path, " -> ")
}
