package trees

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	return depthFirstSearch(n, array)
}

func depthFirstSearch(currentNode *Node, array []string) []string {
	if currentNode == nil {
		return array
	}
	array = append(array, currentNode.Name)
	for _, child := range currentNode.Children {
		array = depthFirstSearch(child, array)
	}
	return array
}
