package genomicintervaltree

import (
	"fmt"
)

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
func (i *IntervalTree) SearchOverlap(queryNode Node) {
	//parentNode := nil
	currentNode := i.root
	fmt.Println(currentNode)
}
