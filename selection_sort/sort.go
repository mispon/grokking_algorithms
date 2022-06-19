package selection_sort

func SortInt(values []int) []int {
	for i := 0; i < len(values)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[minIdx] {
				minIdx = j
			}
		}
		values[i], values[minIdx] = values[minIdx], values[i]
	}

	return values
}
