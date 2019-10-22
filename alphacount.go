package piscine

func AlphaCount(str string) int {
	k := 0
	strrune := []rune(str)
	for letter := range strrune {
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			k++
		}
	}
	return k
}
