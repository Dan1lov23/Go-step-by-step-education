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

type person3 struct{
	name string
	age int
}

func (p *person3) update_age(new_age int){
	p.age = new_age
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

	person3 := person3{"bob", 18}
	person3P := &person3
	fmt.Println("Before", person3)
	person3P.update_age(33)
	fmt.Println("After", person3)	
	
}
