package piscine

func StrLen(str string) int {

	k := 0
	runchik := []rune(str)
	for index := range runchik {
		index = index + 1
		k = index
	}
	return k
}
