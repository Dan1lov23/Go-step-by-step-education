// closures practic

package main

import "fmt"

func closureOne() func() {

	num := 0

	increment := func() {
		num += 1
		fmt.Println(num)
	}

	return increment

}

func main() {
		
	fn := closureOne()

	fn() // 1
	fn() // 2
	fn() // 3
	fn() // 4
}
