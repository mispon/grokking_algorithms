package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_FindPath(t *testing.T) {
	graphFactory := func() *Graph {
		d := &Node{
			Key:   "d",
			Value: map[string]any{},
		}
		c := &Node{
			Key: "c",
			Value: map[string]any{
				"d": 2,
			},
			Nodes: []*Node{d},
		}
		b := &Node{
			Key: "b",
			Value: map[string]any{
				"c": 1,
				"d": 5,
			},
			Nodes: []*Node{c, d},
		}
		a := &Node{
			Key: "a",
			Value: map[string]any{
				"b": 4,
				"c": 3,
			},
			Nodes: []*Node{b, c},
		}

		return a.AsGraph()
	}

	t.Run("empty graph", func(t *testing.T) {
		g := &Graph{}

		cost, ok := g.FindPath("foo")

		require.Zero(t, cost)
		require.False(t, ok)
	})

	t.Run("find shortest path", func(t *testing.T) {
		g := graphFactory()

		cost, ok := g.FindPath("d")

		require.Equal(t, 5, cost)
		require.True(t, ok)
	})
}
