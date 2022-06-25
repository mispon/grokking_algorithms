package graph

import (
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
