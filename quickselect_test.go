package quickselect

import (
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"testing"
)

// go test -v -run="IntSliceQuickSelect"
func TestIntSliceQuickSelect(t *testing.T) {
	lesser, greater := func(a, b int) bool { return a < b }, func(a, b int) bool { return a > b }

	N := 1000
	largeA, largeB := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		largeA[i], largeB[i] = i, i
	}
	lo.Shuffle(largeA)
	lo.Shuffle(largeB)

	testCases := []struct {
		Array     []int
		ExpectedK []int
		K         int
		cmp       CompareFn[int]
	}{
		{[]int{0, 14, 16, 29, 12, 2, 4, 4, 7, 29}, []int{0, 2, 4, 4}, 4, lesser},
		{[]int{0, 14, 16, 29, 12, 2, 4, 4, 7, 29}, []int{14, 16, 29, 29}, 4, greater},
		{[]int{9, 3, 2, 18}, []int{9, 3, 2, 18}, 4, lesser},
		{[]int{9, 3, 2, 18}, []int{9, 3, 2, 18}, 4, greater},
		{[]int{16, 29, -11, 25, 28, -14, 10, 4, 7, -27}, []int{-27, -11, -14}, 3, lesser},
		{[]int{16, 29, -11, 25, 28, -14, 10, 4, 7, -27}, []int{29, 25, 28}, 3, greater},
		{largeA, []int{0, 1, 2, 3, 4}, 5, lesser},
		{largeB, []int{999, 998, 997, 996, 995}, 5, greater},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, []int{4, 5, 5, 6}, 4, greater},
	}

	for i, tc := range testCases {
		QuickSelect(tc.Array, tc.K, tc.cmp)

		if !arrayWithSameElements(tc.Array[0:tc.K], tc.ExpectedK) {
			t.Errorf("[%d], Expected smallest K elements to be '%v', but got '%v'", i, tc.ExpectedK, tc.Array[0:tc.K])
		}
	}
}

/*
 *  go test -bench="BenchmarkQuick"
 *  Benchmark test, quickselect 10x faster than quicksort
 */
func BenchmarkQuickSelectSize1e6K1e1(b *testing.B) { cmpQuickselectAndSort(b, 1e6, 1e1, true) }
func BenchmarkQuickSelectSize1e6K1e2(b *testing.B) { cmpQuickselectAndSort(b, 1e6, 1e2, true) }
func BenchmarkQuickSelectSize1e6K1e3(b *testing.B) { cmpQuickselectAndSort(b, 1e6, 1e3, true) }
func BenchmarkQuickSelectSize1e6K1e4(b *testing.B) { cmpQuickselectAndSort(b, 1e6, 1e4, true) }
func BenchmarkQuickSelectSize1e6K1e5(b *testing.B) { cmpQuickselectAndSort(b, 1e6, 1e5, true) }

func BenchmarkQuickSortSize1e6K1e1(b *testing.B) { cmpQuickselectAndSort(b, 1e6, 1e1, false) }
func BenchmarkQuickSortSize1e6K1e2(b *testing.B) { cmpQuickselectAndSort(b, 1e6, 1e2, false) }
func BenchmarkQuickSortSize1e6K1e3(b *testing.B) { cmpQuickselectAndSort(b, 1e6, 1e3, false) }
func BenchmarkQuickSortSize1e6K1e4(b *testing.B) { cmpQuickselectAndSort(b, 1e6, 1e4, false) }
func BenchmarkQuickSortSize1e6K1e5(b *testing.B) { cmpQuickselectAndSort(b, 1e6, 1e5, false) }

func cmpQuickselectAndSort(b *testing.B, size, k int, quickselect bool) {
	lesser := func(a, b int) bool { return a < b }
	b.StopTimer()
	data := make([]int, size)
	for i := 0; i < b.N; i++ {
		for i := 0; i < size; i++ {
			data[i] = i
		}
		lo.Shuffle(data)
		if quickselect {
			b.StartTimer()
			QuickSelect(data, k, lesser)
			b.StopTimer()
		} else {
			b.StartTimer()
			slices.Sort(data)
			b.StopTimer()
		}
	}
}
