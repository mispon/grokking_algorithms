package graph

import (
	"testing"

	"github.com/mispon/grokking_algorithms/internal/types"
	"github.com/stretchr/testify/require"
)

func Test_BFS(t *testing.T) {
	tests := []struct {
		graph  *types.GraphNode[int]
		target int
		want   *types.GraphNode[int]
	}{
		{
			graph:  nil,
			target: 4,
			want:   nil,
		},
		{
			graph: &types.GraphNode[int]{
				Value:     4,
				Neighbors: nil,
			},
			target: 4,
			want: &types.GraphNode[int]{
				Value:     4,
				Neighbors: nil,
			},
		},
		{
			graph: &types.GraphNode[int]{
				Value: 1,
				Neighbors: []*types.GraphNode[int]{
					{
						Value:     2,
						Neighbors: nil,
					},
					{
						Value: 3,
						Neighbors: []*types.GraphNode[int]{
							{
								Value:     4,
								Neighbors: nil,
							},
							{
								Value:     11,
								Neighbors: nil,
							},
							{
								Value:     12,
								Neighbors: nil,
							},
						},
					},
					{
						Value:     4,
						Neighbors: nil,
					},
				},
			},
			target: 11,
			want: &types.GraphNode[int]{
				Value:     11,
				Neighbors: nil,
			},
		},
		{
			graph: &types.GraphNode[int]{
				Value: 1,
				Neighbors: []*types.GraphNode[int]{
					{
						Value:     2,
						Neighbors: nil,
					},
					{
						Value: 3,
						Neighbors: []*types.GraphNode[int]{
							{
								Value: 4,
								Neighbors: []*types.GraphNode[int]{
									{
										Value:     31,
										Neighbors: nil,
									},
								},
							},
							{
								Value:     11,
								Neighbors: nil,
							},
							{
								Value:     12,
								Neighbors: nil,
							},
						},
					},
					{
						Value: 4,
						Neighbors: []*types.GraphNode[int]{
							{
								Value:     30,
								Neighbors: nil,
							},
						},
					},
				},
			},
			target: 4,
			want: &types.GraphNode[int]{
				Value: 4,
				Neighbors: []*types.GraphNode[int]{
					{
						Value:     30,
						Neighbors: nil,
					},
				},
			},
		},
		{
			graph: &types.GraphNode[int]{
				Value: 1,
				Neighbors: []*types.GraphNode[int]{
					{
						Value:     2,
						Neighbors: nil,
					},
					{
						Value: 3,
						Neighbors: []*types.GraphNode[int]{
							{
								Value:     4,
								Neighbors: nil,
							},
							{
								Value:     11,
								Neighbors: nil,
							},
							{
								Value:     12,
								Neighbors: nil,
							},
						},
					},
					{
						Value:     4,
						Neighbors: nil,
					},
				},
			},
			target: 57,
			want:   nil,
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			node := BFS(tc.graph, tc.target)
			require.Equal(t, tc.want, node)
		})
	}
}
