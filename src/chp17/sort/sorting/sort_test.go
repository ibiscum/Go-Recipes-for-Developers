package sorting

import (
	"math/rand"
	"testing"
	"time"
)

func TestSortTimesAscending(t *testing.T) {
	input := []time.Time{
		time.Date(2023, 2, 1, 12, 8, 37, 0, time.Local),
		time.Date(2021, 5, 6, 9, 48, 11, 0, time.Local),
		time.Date(2022, 11, 13, 17, 13, 54, 0, time.Local),
		time.Date(2022, 6, 23, 22, 29, 28, 0, time.Local),
		time.Date(2023, 3, 17, 4, 5, 9, 0, time.Local),
	}
	t.Logf("Input: %v", input)
	output := SortTimes(input, true)
	t.Logf("Output: %v", output)
	for i := 1; i < len(output); i++ {
		if !output[i-1].Before(output[i]) {
			t.Error("Wrong order")
		}
	}
}

func TestSortTimesDescending(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping descending test")
	}
	input := []time.Time{
		time.Date(2023, 2, 1, 12, 8, 37, 0, time.Local),
		time.Date(2021, 5, 6, 9, 48, 11, 0, time.Local),
		time.Date(2022, 11, 13, 17, 13, 54, 0, time.Local),
		time.Date(2022, 6, 23, 22, 29, 28, 0, time.Local),
		time.Date(2023, 3, 17, 4, 5, 9, 0, time.Local),
	}
	output := SortTimes(input, false)
	for i := 1; i < len(output); i++ {
		if !output[i-1].After(output[i]) {
			t.Error("Wrong order")
		}
	}
}

func TestSortParallel(t *testing.T) {
	input := []time.Time{
		time.Date(2023, 2, 1, 12, 8, 37, 0, time.Local),
		time.Date(2021, 5, 6, 9, 48, 11, 0, time.Local),
		time.Date(2022, 11, 13, 17, 13, 54, 0, time.Local),
		time.Date(2022, 6, 23, 22, 29, 28, 0, time.Local),
		time.Date(2023, 3, 17, 4, 5, 9, 0, time.Local),
	}
	t.Run("SortAscending", func(t *testing.T) {
		t.Parallel()
		output := SortTimes(input, false)
		for i := 1; i < len(output); i++ {
			if !output[i-1].After(output[i]) {
				t.Error("Wrong order")
			}
		}
	})
	t.Run("SortDescending", func(t *testing.T) {
		t.Parallel()
		output := SortTimes(input, false)
		for i := 1; i < len(output); i++ {
			if !output[i-1].After(output[i]) {
				t.Error("Wrong order")
			}
		}
	})
}

func BenchmarkSortAscending(b *testing.B) {
	input := []time.Time{
		time.Date(2023, 2, 1, 12, 8, 37, 0, time.Local),
		time.Date(2021, 5, 6, 9, 48, 11, 0, time.Local),
		time.Date(2022, 11, 13, 17, 13, 54, 0, time.Local),
		time.Date(2022, 6, 23, 22, 29, 28, 0, time.Local),
		time.Date(2023, 3, 17, 4, 5, 9, 0, time.Local),
	}
	for i := 0; i < b.N; i++ {
		SortTimes(input, true)
	}
}

func benchmarkSort(b *testing.B, nItems int, asc bool) {
	input := make([]time.Time, nItems)
	t := time.Now().UnixNano()
	for i := 0; i < nItems; i++ {
		input[i] = time.Unix(0, t-int64(i))
	}
	rand.Shuffle(len(input), func(i, j int) { input[i], input[j] = input[j], input[i] })
	for i := 0; i < b.N; i++ {
		SortTimes(input, asc)
	}
}

func BenchmarkSort1000Ascending(b *testing.B)  { benchmarkSort(b, 1000, true) }
func BenchmarkSort100Ascending(b *testing.B)   { benchmarkSort(b, 100, true) }
func BenchmarkSort10Ascending(b *testing.B)    { benchmarkSort(b, 10, true) }
func BenchmarkSort1000Descending(b *testing.B) { benchmarkSort(b, 1000, false) }
func BenchmarkSort100Descending(b *testing.B)  { benchmarkSort(b, 100, false) }
func BenchmarkSort10Descending(b *testing.B)   { benchmarkSort(b, 10, false) }
