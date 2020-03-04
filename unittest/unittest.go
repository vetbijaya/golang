package main

import (
	"fmt"
	"os"
)

func main(){

fmt.Println("This is the unit case testing")

}
// addition is used to add numbers
func addition (a ...int)int {

	var total int

	for _, value := range a{
		total += value
	}
	return total
}
// multiply is a function used for multiplication
func multiply(a ...int)int{
	var total =1
	for _, value:= range a {
		total *= value
	}
	return total
}
// openFile will open the file that is presented in current package 
func openFile(filename string) error{
_, err:= os.Open(filename)
return err
}