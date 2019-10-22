package piscine

import "fmt"

func AlphaCount(str string) int {
	k := 0
	for letter := range str{
	if (letter >= 'a' && letter <= 'z') || (letter>= 'A' && letter <= 'Z'){
	k++
	}
	return k
}
}

 