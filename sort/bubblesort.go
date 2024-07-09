package sort

import "Akmf_Algorithm/constraints"

func BubbleSort[T constraints.Ordered](arr []T) []T {
	for i := len(arr) - 1; i >= 0 ; i-- {
		flag := false
		for j := 0; j + 1 <= i; j++ {
			if arr[j] > arr[j+1] {
				flag = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !flag {
			break
		}
	}
	return arr
}

// form TheAlgorithms/Go
func Bubble[T constraints.Ordered](arr []T) []T {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr
}