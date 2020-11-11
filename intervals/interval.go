package intervals

// Interval holds interval info
type Interval struct {
	Start int
	Stop  int
	Data  map[string]interface{}
}

//Overlap determines whether a start and stop overlaps this interval
func (i Interval) Overlap(begin int, end int) bool{
	if !(i.Stop <= begin) || !(i.Start >= end){
		return false
	}
	return true

}