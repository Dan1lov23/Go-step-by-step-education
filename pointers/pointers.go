package main

import "fmt"

func changeVarValue(x *int) {
	*x = (*x) * (*x)
}

func createPointer(num int) *int { 

	p := new(int)
	*p = num

	return p

}

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

	valueForChange := 5 // <-- before
	fmt.Println(valueForChange)
	changeVarValue(&valueForChange) // <-- function chnage variable value by using pointer
	fmt.Println(valueForChange)     // <-- after

	fmt.Println(createPointer(12)) // <-- pointer
	fmt.Println(createPointer(18))
}
