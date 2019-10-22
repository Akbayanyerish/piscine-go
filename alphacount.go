package piscine

func AlphaCount(str string) int {
	k := 0
	strrune := []rune(str)
	for letter := range strrune {
		if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) {
			k++
		}
	}
	return k
}
