package piscine

func PrintStr(str string) {
	for word := range str {
		Printf(word)
	}
}
