package piscine

func StrLen(str string) int {

	k := 1
	for index := range str {
		k = index
		index = index + 1

	}
	return k
}
