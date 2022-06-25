package graph

import "fmt"

type (
	// Node represents single graph node
	Node struct {
		// Key is unique node key
		Key string
		// Value is node value
		Value map[string]any
		// Nodes are list of neighbours nodes
		Nodes []*Node
		// Parent is ptr to parent node
		Parent *Node
	}

	// Graph is data structure that represent nodes and it's relations
	Graph Node
)

// AsGraph returns node as graph root
func (n *Node) AsGraph() *Graph {
	return (*Graph)(n)
}

func (g *Graph) AsNode() *Node {
	return (*Node)(g)
}

// Print display all nodes with tab inted
func (g *Graph) Print() {
	printNode(g.AsNode(), "")
}

// printNode recursivly prints nodes
func printNode(n *Node, intend string) {
	fmt.Println(intend, n.Key)
	for _, sn := range n.Nodes {
		printNode(sn, intend+"\t")
	}
}
