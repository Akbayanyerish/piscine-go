package piscine

func StrLen(str string) int {

	k := 1
	for index := range str {

		index = index + 1
		k = index
	}
	return k
}
