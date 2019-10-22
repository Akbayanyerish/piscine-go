package main
import "fmt"

func main() {
var str string := "hello 4455 "
k:=0
for _, letter := range str{
if letter >= "a" letter <= "z" {
	k++
}
}
fmt.Println(k)
}
