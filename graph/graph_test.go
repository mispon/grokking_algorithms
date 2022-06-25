package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AsGraph(t *testing.T) {
	n := &Node{}
	g := n.AsGraph()
	require.IsType(t, &Graph{}, g)
}

func Test_AsNode(t *testing.T) {
	g := &Graph{}
	n := g.AsNode()
	require.IsType(t, &Node{}, n)
}

func Test_Print(t *testing.T) {
	t.SkipNow()
	g := newGraphMock()
	g.Print()
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
