package arrays

func ArrayIterator(arr []int) (sum int) {

	for _, number := range arr {
		sum += number		
	}
	
	return
}
