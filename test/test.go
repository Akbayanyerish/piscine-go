package main

import "fmt"

func main() {

	str := "Háloo Potato!"
	nb := StrLen(str)
	fmt.Println(nb)

}

func StrLen(str string) int {
	
	k := 0;
	runchik := []rune(str)
	for index := range runchik {

		index = index + 1
		k = index
	}
	return k
}

