package arrays

func SliceSum(arr []int) (sum int) {

	for _, number := range arr {
		sum += number
	}

	return
}

func SliceSumAll(slices ...[]int) []int {

	numSlices := len(slices)
	sums := make([]int, numSlices)

	for i, slice := range slices {
		sums[i] = SliceSum(slice)
	}

	return sums
}
