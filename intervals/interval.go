package intervals

import (
	"fmt"
)

// Interval holds interval info
type Interval struct {
	Start int
	Stop  int
	Data  map[string]interface{}
}

//Overlap determines whether a start and stop overlaps this interval
func (i Interval) Overlap(begin int, end int) bool {
	//
	if !(i.Stop <= begin || i.Start >= end) {
		fmt.Println(i.Stop <= begin)
		fmt.Println(i.Start >= end)
		fmt.Println(!(false || false))
		fmt.Println(i.Start, i.Stop, begin, end)
		return false
	}
	return true

}
