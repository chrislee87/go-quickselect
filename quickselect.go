package quickselect

type CompareFn[T any] func(a, b T) bool

/*
 * After QuickSelect, array[:k] is top-k without order
 * time-complexity is O(N), N is array.length
 */
func QuickSelect[T any](array []T, k int, cmp CompareFn[T]) {

	if len(array) == 0 || k >= len(array) {
		return
	}

	left, right := 0, len(array)-1

	for {
		// insertion sort for small ranges
		if right-left <= 20 {
			for i := left + 1; i <= right; i++ {
				for j := i; j > 0 && cmp(array[j], array[j-1]); j-- {
					array[j], array[j-1] = array[j-1], array[j]
				}
			}
			return
		}

		// median-of-three to choose pivot
		pivotIndex := left + (right-left)/2
		if cmp(array[right], array[left]) {
			array[right], array[left] = array[left], array[right]
		}
		if cmp(array[pivotIndex], array[left]) {
			array[pivotIndex], array[left] = array[left], array[pivotIndex]
		}
		if cmp(array[right], array[pivotIndex]) {
			array[right], array[pivotIndex] = array[pivotIndex], array[right]
		}

		// partition
		array[left], array[pivotIndex] = array[pivotIndex], array[left]
		ll := left + 1
		rr := right
		for ll <= rr {
			for ll <= right && cmp(array[ll], array[left]) {
				ll++
			}
			for rr >= left && cmp(array[left], array[rr]) {
				rr--
			}
			if ll <= rr {
				array[ll], array[rr] = array[rr], array[ll]
				ll++
				rr--
			}
		}
		array[left], array[rr] = array[rr], array[left] // swap into right place
		pivotIndex = rr

		if k == pivotIndex {
			return
		}

		if k < pivotIndex {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}
}
