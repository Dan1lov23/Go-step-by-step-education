// functions practic

package main

import "fmt"

func sum(a int, b int) (sum int) { // <-- you can set type fot return variables and just write return
	return a + b
}

func endlessVarsSum(nums ...int) (sum int) { // <-- functions with unlimitade params

	for _, nums := range nums {
		sum += nums
	}

	return

}

func add(a int, b int, sumFunction func(int, int) int) (sum int) { // <-- function as function parameter
	return sumFunction(a, b)
}

func add2(a int, b int, sumFunciton func(int, int) int) int {
	return sumFunciton(a, b)
}

func getFunction(functionName string) func(int, int) int {
	switch functionName {
	case "add":
		return func(a int, b int) int {return a + b}
	case "mult":
		return func(a int, b int) int {return a * b}
	}

	return func(a int, b int) int {return a - b}

}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(endlessVarsSum(1, 2, 3, 4, 5))
	fmt.Println(add(1, 2, sum))
	aninonimSumFunction := func(a int, b int) int { return a + b } // <-- anonim sum function
	fmt.Println(aninonimSumFunction(1, 2))
	fmt.Println(add2(1, 2, func(a int, b int) int { return a + b }))
	someFunction := getFunction("mult")
	fmt.Println(someFunction(5, 6)) // <-- function as function result
}
