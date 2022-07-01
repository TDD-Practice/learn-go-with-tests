package iteration

func Iterator(character string, times int) (sequence string) {
	for i := 0; i < times; i++ {
		sequence += character
	}
	return
}
