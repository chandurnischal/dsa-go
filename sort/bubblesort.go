package sort

func BubbleSort(list []int) {
	n := len(list)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}
