package genomicintervaltree

// IntervalTree holder
type IntervalTree struct {
	root *Node
}

// AddNode adds a node to an existing tree in place
func (i *IntervalTree) AddNode(newNode Node) {
	currentNode := i.root
	for currentNode != nil {
		if newNode.NodeInterval.Start <= currentNode.NodeInterval.Start {
			if (currentNode.Left == &Node{}) {
				currentNode.Left = &newNode
				return
			}
			currentNode = currentNode.Left
		} else {
			if (currentNode.Right == &Node{}) {
				currentNode.Right = &newNode
				return
			}
			currentNode = currentNode.Right
		}
	}

}

// SearchOverlap searches the interval tree for an overlap
func (i *IntervalTree) SearchOverlap(queryNode Node) (Node, Node, bool) {
	currentNode := i.root
	for !(currentNode != nil) {
		if queryNode.NodeInterval.Overlap(currentNode.NodeInterval) {
			print("Overlapping with ", currentNode.NodeInterval)
			return queryNode, *currentNode, true
		}
		if currentNode.Max >= queryNode.NodeInterval.Start {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
	return queryNode, Node{}, false
}
