// variables practic

package main

import "fmt"

func main() {
	var num1 int // <-- will be zero default meaning
	var num2 int = 0
	num3 := 0
	var flag1 bool // <-- false it is a def value for bool variables
	flag2 := false
	var numSlice1 [5]int  // <-- will be zero default meaning
	numSlice2 := [5]int{} // <-- also as on the line above
	fmt.Println(num1, num2, num3, flag1, flag2, numSlice1, numSlice2)
}
