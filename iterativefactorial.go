package piscine
import "fmt"
func IterativeFactorial(nb int) int {
	result:=1
	if nb==0 {
		result:=1
	} 
	for i:=1; i <= nb; i++ {
		result:=result*i
	} 
fmt.Println(result)
}