package intervals

// Node holds tree node info
type Node struct {
	interval Interval
	left     *Node
	right    *Node
	max      int
}

//HasChild determines whether this is a parent node or the end of a branch
func (n Node) HasChild() bool {
	if n.left != nil || n.right != nil {
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
	if n.left != nil {
		childIntervals = append(childIntervals, n.left.interval)
	}
	if n.right != nil {
		childIntervals = append(childIntervals, n.right.interval)
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
