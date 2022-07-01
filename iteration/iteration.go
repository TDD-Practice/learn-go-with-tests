package iteration

func Iterator(character string, times int) (sequence string) {
	for i := 0; i < 5; i++ {
		sequence += character
	}
	return
}
