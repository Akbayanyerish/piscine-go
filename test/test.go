package main

import "fmt"

func main() {

	str := "Heloo Potato!"
	nb := StrLen(str)
	fmt.Println(nb)

}

func StrLen(str string) int {
	
	k := 0;
	for index := range str {

		index = index + 1
		k = index
	}
	return k
}

