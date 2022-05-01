package iteration

func Repeat(characterToRepeat string, numberOfRepetitions int) string {
	var repeated string

	for i := 0; i < numberOfRepetitions; i++ {
		repeated += characterToRepeat
	}

	return repeated
}
