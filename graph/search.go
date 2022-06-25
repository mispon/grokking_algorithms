package graph

type (
	searchPredicate func(n *Node) bool
)

// Search process bfs search by graph
// bfs is Breadth-First Search algorithm
func (g *Graph) Search(predicate searchPredicate) (*Node, bool) {
	if g == nil {
		return nil, false
	}

	root := (*Node)(g)
	if predicate(root) {
		return root, true
	}

	var (
		nodes    = g.Nodes
		searched = map[string]bool{}
	)

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]

		if searched[node.Key] {
			continue
		}

		if predicate(node) {
			return node, true
		}

		nodes = append(nodes, node.Nodes...)

		searched[node.Key] = true
	}

	return nil, false
}
