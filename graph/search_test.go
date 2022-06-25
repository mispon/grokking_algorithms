package graph

import (
	"fmt"
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

func newGraphMock() *Graph {
	root := &Node{
		Key: "root",
		Value: map[string]any{
			"label": "root",
			"foo":   "bar",
		},
		Parent: nil,
	}

	for i := 0; i < 10; i++ {
		nodeName := fmt.Sprintf("node-%d", i+1)
		node := &Node{
			Key: nodeName,
			Value: map[string]any{
				"label": nodeName,
			},
			Parent: root,
		}
		for j := 0; j < 10; j++ {
			subNodeName := fmt.Sprintf("%s-subnode-%d", nodeName, j+1)
			subNode := &Node{
				Key: subNodeName,
				Value: map[string]any{
					"label": subNodeName,
				},
				Parent: node,
			}
			node.Nodes = append(node.Nodes, subNode)
		}
		root.Nodes = append(root.Nodes, node)
	}

	return root.AsGraph()
}
