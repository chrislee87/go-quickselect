package quickselect

func ArrayWithSameElements[T comparable](array1, array2 []T) bool {
	elements := make(map[T]int)

	for _, elem1 := range array1 {
		elements[elem1]++
	}

	for _, elem2 := range array2 {
		elements[elem2]--
	}

	for _, count := range elements {
		if count != 0 {
			return false
		}
	}
	return true
}
