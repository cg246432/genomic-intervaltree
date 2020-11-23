package genomicintervaltree

// Node holds tree node info
type Node struct {
	NodeInterval Interval
	Left         *Node
	Right        *Node
	Max          int
}

//HasChild determines whether this is a parent node or the end of a branch
func (n Node) HasChild() bool {
	if n.Left != nil || n.Right != nil {
		return true
	}
	return false
}

//MaxChild determines largest child interval
func (n Node) MaxChild() Interval {
	if !(n.HasChild()) {
		return Interval{}
	}
	var childIntervals []Interval
	if n.Left != nil {
		childIntervals = append(childIntervals, n.Left.NodeInterval)
	}
	if n.Right != nil {
		childIntervals = append(childIntervals, n.Right.NodeInterval)
	}
	if len(childIntervals) == 1 {
		return childIntervals[0]
	}
	if len(childIntervals) == 2 {
		if childIntervals[0].Sort(childIntervals[1]) == -1 ||
			childIntervals[0].Sort(childIntervals[1]) == 0 {
			return childIntervals[0]
		}
		if childIntervals[0].Sort(childIntervals[1]) == 1 {
			return childIntervals[1]
		}
	}
	return Interval{}
}
