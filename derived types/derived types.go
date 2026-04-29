// derived types theme

package main

import "fmt"

type mile int

func addMile(distanceOne mile, distanceTwo mile) (distanceSum mile) {
	
	return distanceOne + distanceTwo

}

func main() {

	var mileCounter mile // default value 0, because mile type it is int type, and int default value - 0
	fmt.Println(mileCounter)

	var distanceOne mile = 12
	var distanceTwo mile = 12

	fmt.Println(addMile(distanceOne, distanceTwo)) // <-- only mile types params

}
