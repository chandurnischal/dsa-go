package sort

func SelectionSort(array []int) []int {
	n := len(array)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}
	return array
}
