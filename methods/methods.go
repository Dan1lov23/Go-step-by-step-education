package main

import "fmt"

type mile int

func (maileValue mile) mileLog() { // for current data type
	fmt.Println(maileValue)
}

func (maileValue mile) incrementMile(incrementValue mile) (newValue mile) { // like default function but with relation with current data type
	return maileValue + incrementValue

}

type userType struct {
	username string
	password string
}

func (user userType) chagePassword(newPassword string) (newUserProfile userType) {

	return userType{user.username, newPassword}

} // also for struct

func main() {
	var mileValue1 mile = 12
	mileValue1.mileLog()
	valueAfterIncrement := mileValue1.incrementMile(12)
	fmt.Println(valueAfterIncrement)

	user1 := userType{"Bob", "Bob234!"}
	newUser1Profile := user1.chagePassword("Bob345!")
	fmt.Println(newUser1Profile)

}
