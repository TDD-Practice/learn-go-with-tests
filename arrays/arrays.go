package arrays

func SliceSum(arr []int) (sum int) {

	for _, number := range arr {
		sum += number
	}

	return
}

func SliceSumAll(slices ...[]int) (sums []int) {

	for _, slice := range slices {
		sums = append(sums, SliceSum(slice))
	}

	return sums
}
