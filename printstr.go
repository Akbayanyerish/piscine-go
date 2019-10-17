package piscine

import "fmt"

func PrintStr(str string) {
	for word := range str {
		fmt.Printf(word)
	}
}
