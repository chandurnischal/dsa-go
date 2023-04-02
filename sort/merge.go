package sort

func merge(left, right []int) []int {
	res := []int{}

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(res, right...)
		}
		if len(right) == 0 {
			return append(res, left...)
		}
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	return res
}

func MergeSort(array []int) []int {
	n := len(array)
	if n <= 1 {
		return array
	}
	mid := int(n / 2)
	left := array[:mid]
	right := array[mid:]
	return merge(MergeSort(left), MergeSort(right))
}
