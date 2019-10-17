package piscine
import "fmt"
func PointOne(n *int){
*n=1
}  

func Piscine() {
    n := 0
  piscine.PointOne(&n)
    fmt.Println(n)
}