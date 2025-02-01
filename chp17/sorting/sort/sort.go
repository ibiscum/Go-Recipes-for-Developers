package sorting

import (
	"sort"
	"time"
)

// Sort times in ascending or descending order
func SortTimes(input []time.Time, asc bool) []time.Time {
	output := make([]time.Time, len(input))
	copy(output, input)
	if asc {
		sort.Slice(output, func(i, j int) bool {
			return output[i].Before(output[j])
		})
		return output
	}
	sort.Slice(output, func(i, j int) bool {
		return output[j].Before(output[i])
	})
	return output
}
