package intervals

// Interval holds interval info
type Interval struct {
	Start int
	Stop  int
	Data  map[string]interface{}
}

// It's required for genomic intervals to be sorted
func (i *Interval) checkorder() {
	if i.Start > i.Stop{
		i.Start, i.Stop = i.Stop, i.Start
	}
}

// NewInterval acts as a settter
func NewInterval(start int, stop int) Interval{
	newInterval := Interval{
		Start: start,
		Stop: stop,
	}
	newInterval.checkorder()
	return newInterval
}

// DeleteInterval sets the interval back to empty
func (i *Interval) DeleteInterval() {
	i = &Interval{}
}

//Overlap determines whether a start and stop overlaps this interval
func (i Interval) Overlap(begin int, end int) bool {
	if i.Stop <= begin || i.Start >= end {
		return false
	}
	return true

}

// OverlapSize returns size of overlap
func (i Interval) OverlapSize(begin int, end int) int {
	if i.Overlap(begin, end){
		var minOverlap int
		minOverlap = i.Stop - i.Start
		if (i.Stop - begin) < minOverlap{
			minOverlap = i.Stop - begin
		}
		if (end - begin) < minOverlap{
			minOverlap = end - begin
		}
		if (end - i.Start) < minOverlap{
			minOverlap = end - i.Start
		}
		return minOverlap
	}
	return 0
}