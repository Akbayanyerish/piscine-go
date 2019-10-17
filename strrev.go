package piscine

func StrRev(s string) string {
	j := 0
	var strrune = []rune(s)
	var s2 string
	for index := range strrune {
		j = index
	}
	j++
	for i := j - 1; i >= 0; i-- {
		s2 = s2 + string(strrune[i])
	}
	return s2
}
