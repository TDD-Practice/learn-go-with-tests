package iteration

func Iterator(character string) (sequence string) {
	for i := 0; i < 5; i++ {
		sequence += character
	}
	return
}
