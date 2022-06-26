package graph

import "math"

// FindPath search for shortest path uses Dijkstra algorythm
func (g *Graph) FindPath(target string) (string, int, bool) {
	if g == nil || len(g.Nodes) == 0 || target == "" {
		return "", 0, false
	}

	var (
		costs     = map[string]int{target: math.MaxInt}
		parents   = map[string]string{target: "unknown"}
		processed = map[string]bool{}
	)

	for name, cost := range g.Value {
		costs[name] = cost.(int)
		parents[name] = g.Key
	}

	node := findLowerCostNode(g, costs, processed)
	for node != nil {
		cost := costs[node.Key]
		for _, n := range node.Nodes {
			newCost := cost + node.Value[n.Key].(int)
			if _, ok := costs[n.Key]; !ok || costs[n.Key] > newCost {
				costs[n.Key] = newCost
				parents[n.Key] = node.Key
			}
		}
		processed[node.Key] = true
		node = findLowerCostNode(g, costs, processed)
	}

	return buildPath(parents, target), costs[target], true
}

func findLowerCostNode(g *Graph, costs map[string]int, processed map[string]bool) *Node {
	var (
		cheaperKey = ""
		minCost    = math.MaxInt
	)

	for name, cost := range costs {
		if minCost > cost && !processed[name] {
			cheaperKey = name
			minCost = cost
		}
	}

	node, _ := g.Search(func(n *Node) bool {
		return n.Key == cheaperKey
	})

	return node
}

func buildPath(parents map[string]string, target string) string {
	steps := []string{target}

	n := parents[target]
	for n != "" {
		steps = append(steps, n)
		n = parents[n]
	}

	path := ""
	for i := len(steps) - 1; i >= 0; i-- {
		path += steps[i]
		if i > 0 {
			path += " -> "
		}
	}
	return path
}
