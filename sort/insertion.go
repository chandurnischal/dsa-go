package sort

func InsertionSort(array []int) []int {
	n := len(array)

	for i := 1; i < n; i++ {
		key := array[i]
		j := i - 1

		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
	return array
}
