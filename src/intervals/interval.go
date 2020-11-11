package intervals

// Interval holds interval info
type Interval struct {
	start int
	stop  int
	data  map[string]interface{}
}

func (i Interval) overlap(begin int, end int) bool{
	return true
}