package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Search(t *testing.T) {
	t.Run("empty graph", func(t *testing.T) {
		graph := Graph{}

		node, ok := graph.Search(func(n *Node) bool {
			return n.Key == "any_key"
		})

		require.Nil(t, node)
		require.False(t, ok)
	})

	t.Run("search node by key", func(t *testing.T) {
		tests := []string{
			"root",
			"node-6",
			"node-2-subnode-4",
		}

		for _, key := range tests {
			graph := newGraphMock()

			node, ok := graph.Search(func(n *Node) bool {
				return n.Key == key
			})

			require.NotNil(t, node)
			require.True(t, ok)
		}
	})

	t.Run("search node by value", func(t *testing.T) {
		tests := []string{
			"root",
			"node-6",
			"node-2-subnode-4",
		}

		for _, key := range tests {
			graph := newGraphMock()

			node, ok := graph.Search(func(n *Node) bool {
				return n.Value["label"] == key
			})

			require.NotNil(t, node)
			require.True(t, ok)
		}
	})

	t.Run("try search not existing sub node", func(t *testing.T) {
		graph := newGraphMock()

		node, ok := graph.Search(func(n *Node) bool {
			return n.Key == "node-777-subnode-123"
		})

		require.Nil(t, node)
		require.False(t, ok)
	})

	t.Run("search subnode in node", func(t *testing.T) {
		graph := newGraphMock()

		node, ok := graph.Search(func(n *Node) bool {
			return n.Key == "node-3"
		})
		require.NotNil(t, node)
		require.True(t, ok)

		subGraph := node.AsGraph()
		require.NotNil(t, subGraph)
		require.IsType(t, &Graph{}, subGraph)

		subNode, ok := graph.Search(func(n *Node) bool {
			return n.Key == "node-3-subnode-7"
		})
		require.NotNil(t, subNode)
		require.True(t, ok)
	})
}
