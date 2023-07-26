package pkg

func IsArraySorted(array []int, n uint) bool {
	if n == 1 || n == 0 {
		return true
	}

	if array[n-1] < array[n-2] {
		return false
	} else {
		return IsArraySorted(array, n-1)
	}
}
