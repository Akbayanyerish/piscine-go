package main
import "fmt"
func main() {
str := "aaa B F G4455 "
k:=0
for _, letter := range str{
if (letter >= 'a' && letter <= 'z') || (letter>= 'A' && letter <= 'Z'){
	k++
}
}
//z01.PrintRune(rune())
fmt.Println(k)
}
