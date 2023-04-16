package graph

import (
	"testing"

	"github.com/mispon/grokking_algorithms/internal/types"
	"github.com/stretchr/testify/require"
)

func Test_Dijkstra(t *testing.T) {
	t.Run("ABCD graph", func(t *testing.T) {
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

		t.Run("path to B", func(t *testing.T) {
			path := Dijkstra(nodeA, "B")
			require.Equal(t, "A -> B", path)
		})

		t.Run("path to C", func(t *testing.T) {
			path := Dijkstra(nodeA, "C")
			require.Equal(t, "A -> B -> C", path)
		})

		t.Run("path to D", func(t *testing.T) {
			path := Dijkstra(nodeA, "D")
			require.Equal(t, "A -> B -> C -> D", path)
		})
	})

	t.Run("ABCDEG graph", func(t *testing.T) {
		nodeG := &types.GraphNode[string]{
			Value:     "G",
			Neighbors: nil,
			Costs:     nil,
		}

		nodeE := &types.GraphNode[string]{
			Value: "E",
			Neighbors: []*types.GraphNode[string]{
				nodeG,
			},
			Costs: map[*types.GraphNode[string]]uint{
				nodeG: 1,
			},
		}

		nodeD := &types.GraphNode[string]{
			Value: "D",
			Neighbors: []*types.GraphNode[string]{
				nodeE,
				nodeG,
			},
			Costs: map[*types.GraphNode[string]]uint{
				nodeE: 6,
				nodeG: 3,
			},
		}

		nodeC := &types.GraphNode[string]{
			Value: "C",
			Neighbors: []*types.GraphNode[string]{
				nodeD,
				nodeE,
			},
			Costs: map[*types.GraphNode[string]]uint{
				nodeD: 4,
				nodeE: 2,
			},
		}

		nodeB := &types.GraphNode[string]{
			Value: "B",
			Neighbors: []*types.GraphNode[string]{
				nodeC,
				nodeE,
			},
			Costs: map[*types.GraphNode[string]]uint{
				nodeC: 8,
				nodeE: 7,
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
				nodeC: 5,
			},
		}

		t.Run("path to G", func(t *testing.T) {
			path := Dijkstra(nodeA, "G")
			require.Equal(t, "A -> C -> E -> G", path)
		})

		t.Run("path to D", func(t *testing.T) {
			path := Dijkstra(nodeA, "D")
			require.Equal(t, "A -> C -> D", path)
		})
	})
}
