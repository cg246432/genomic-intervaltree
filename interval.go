package genomicintervaltree

// Interval holds interval info
type Interval struct {
	Start int
	Stop  int
	Data  map[string]interface{}
}

//Intervals Holds a slice of Intervals
type Intervals []Interval

// It's required for genomic intervals to be sorted
func (i *Interval) checkorder() {
	if i.Start > i.Stop {
		i.Start, i.Stop = i.Stop, i.Start
	}
}

// NewInterval acts as a settter
func NewInterval(start int, stop int) Interval {
	newInterval := Interval{
		Start: start,
		Stop:  stop,
	}
	newInterval.checkorder()
	return newInterval
}

// DeleteInterval sets the interval back to empty
func (i *Interval) DeleteInterval() {
	i = nil
}

// Length is size of the interval
func (i Interval) Length() int {
	return i.Stop - i.Start
}

//Overlap determines whether a start and stop overlaps this interval
func (i Interval) Overlap(other Interval) bool {
	if i.Stop <= other.Start || i.Start >= other.Stop {
		return false
	}
	return true

}

// OverlapSize returns size of overlap
func (i Interval) OverlapSize(other Interval) int {
	if i.Overlap(other) {
		var minOverlap int
		minOverlap = i.Stop - i.Start
		if (i.Stop - other.Start) < minOverlap {
			minOverlap = i.Stop - other.Start
		}
		if (other.Stop - other.Start) < minOverlap {
			minOverlap = other.Stop - other.Start
		}
		if (other.Stop - i.Start) < minOverlap {
			minOverlap = other.Stop - i.Start
		}
		return minOverlap
	}
	return 0
}

// RangeMatch determine swhether the other.Starts equal and other.Stops equal
func (i Interval) RangeMatch(other Interval) bool {
	if i.Start == other.Start && i.Stop == other.Stop {
		return true
	}
	return false
}

//DistanceTo calculates the distance between two intervals
func (i Interval) DistanceTo(other Interval) int {
	if i.Overlap(other) {
		return 0
	}
	// i starts earlier
	if i.Start < other.Start {
		return other.Start - i.Stop
	}
	// other starts earlier
	return i.Start - other.Stop
}

// Sort returns whether a comparable interval would sort
// before or after the current one. We sort by other.Start, then
// other.Stop. If this interval sorts before other, we return -1, else
// we return 1. A tie returns 0
func (i Interval) Sort(other Interval) int {
	if i.RangeMatch(other) {
		return 0 // Exit quick if equal
	}

	startDiff := i.Start - other.Start
	if startDiff > 0 {
		return 1 // i start after other start
	}
	if startDiff < 0 {
		return -1
	}
	if startDiff == 0 {
		endDiff := i.Stop - other.Stop
		if endDiff > 0 {
			return 1 // i stops after other
		}
		if endDiff < 0 {
			return -1
		}
	}
	return 0
}
