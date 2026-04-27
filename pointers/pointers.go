package main

import "fmt"

func main() {

	var x int = 4 // <-- number variable
	p := &x       // <-- pointer on x variable
	fmt.Println("Num:", x, ", num pointer:", p)

	*p = 12 // <-- chnage x variable value
	fmt.Println("New number value:", x)

	var strVarPointer *string  // <-- empty pointer
	fmt.Println(strVarPointer) // <-- vil be nil

	pointerN := new(int) // <-- make pointer bu using new command
	*pointerN = 148
	fmt.Println(pointerN)

	var a, b, c = 1, 2, 3
	var pointersArr [4]*int
	pointersArr[0] = &a
	pointersArr[1] = &b
	pointersArr[2] = &c
	fmt.Println(pointersArr)

}
