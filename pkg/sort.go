package pkg

func InsertionSort(input []float64) []float64 {
	for j := 1; j < len(input); j++ {
		key := input[j]
		i := j - 1
		for ; i >= 0 && input[i] > key; i-- {
			input[i+1] = input[i]
		}
		input[i+1] = key
	}
	return input
}
