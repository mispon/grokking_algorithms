package graph

import (
	"testing"

	"github.com/mispon/grokking_algorithms/internal/types"
	"github.com/stretchr/testify/require"
)

func Test_Dijkstra(t *testing.T) {
	nodeD := &types.GraphNode[string]{
		Value:     "D",
		Neighbors: nil,
		Costs:     nil,
	}

	nodeC := &types.GraphNode[string]{
		Value: "C",
		Neighbors: []*types.GraphNode[string]{
			nodeD,
		},
		Costs: map[*types.GraphNode[string]]uint{
			nodeD: 1,
		},
	}

	nodeB := &types.GraphNode[string]{
		Value: "B",
		Neighbors: []*types.GraphNode[string]{
			nodeC,
			nodeD,
		},
		Costs: map[*types.GraphNode[string]]uint{
			nodeC: 3,
			nodeD: 5,
		},
	}

	nodeA := &types.GraphNode[string]{
		Value: "A",
		Neighbors: []*types.GraphNode[string]{
			nodeB,
			nodeC,
		},
		Costs: map[*types.GraphNode[string]]uint{
			nodeB: 2,
			nodeC: 6,
		},
	}

	path := Dijkstra(nodeA, "D")
	require.Equal(t, "A -> B -> C -> D", path)
}
