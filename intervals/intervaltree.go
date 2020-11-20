package intervals

import "fmt"

// IntervalTree holder
type IntervalTree struct {
	root *Node
}

// AddNode adds a node to an existing tree in place
func (i *IntervalTree) AddNode(newNode Node) {
	currentNode := i.root
	for currentNode != nil {
		if newNode.interval.Start <= currentNode.interval.Start {
			if (currentNode.left == &Node{}) {
				currentNode.left = &newNode
				return
			}
			currentNode = currentNode.left
		} else {
			if (currentNode.right == &Node{}) {
				currentNode.right = &newNode
				return
			}
			currentNode = currentNode.right
		}
	}

}

// SearchOverlap searches the interval tree for an overlap
func (i *IntervalTree) SearchOverlap(queryNode Node) {
	//parentNode := nil
	currentNode := i.root
	fmt.Println(currentNode)
}
