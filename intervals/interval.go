package intervals

// Interval holds interval info
type Interval struct {
	Start int
	Stop  int
	Data  map[string]interface{}
}

func (i *Interval) checkorder() {
	if i.Start > i.Stop{
		i.Start, i.Stop = i.Stop, i.Start
	}
}

//Overlap determines whether a start and stop overlaps this interval
func (i Interval) Overlap(begin int, end int) bool {
	i.checkorder()
	if i.Stop <= begin || i.Start >= end {
		return false
	}
	return true

}

// OverlapSize returns size of overlap
func (i Interval) OverlapSize(begin int, end int) int {
	i.checkorder()
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