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
				"d": 1,
			},
			Nodes: []*Node{d},
		}
		b := &Node{
			Key: "b",
			Value: map[string]any{
				"c": 3,
				"d": 5,
			},
			Nodes: []*Node{c, d},
		}
		a := &Node{
			Key: "a",
			Value: map[string]any{
				"b": 2,
				"c": 6,
			},
			Nodes: []*Node{b, c},
		}

		return a.AsGraph()
	}

	t.Run("empty graph", func(t *testing.T) {
		g := &Graph{}

		path, cost, ok := g.FindPath("foo")

		require.Empty(t, path)
		require.Zero(t, cost)
		require.False(t, ok)
	})

	t.Run("find shortest path", func(t *testing.T) {
		g := graphFactory()

		path, cost, ok := g.FindPath("d")

		require.Equal(t, "a -> b -> c -> d", path)
		require.Equal(t, 6, cost)
		require.True(t, ok)
	})
}
