package intervals

// Interval holds interval info
type Interval struct {
	Start int
	Stop  int
	Data  map[string]interface{}
}

func (i Interval) Overlap(begin int, end int) bool{
	return true
}