// structures theme

package main

import "fmt"

type user struct {
	login    string
	password string
	id       int
}

type customerType struct {
	firstname  string
	secondname string
	age        int
}

func incrementAge(customer *customerType) {
	customer.age += 1
}

func main() {

	var user1 user
	fmt.Println(user1) // "", "", 0, because default values

	var user2 user = user{"tom", "tom11623", 2}
	fmt.Println(user2)

	fmt.Println("Login -", user2.login, ",", "password -", user2.password, ",", "id -", user2.id)

	admin := struct { // anononymous structure
		login    string
		password string
	}{
		"Bob",
		"bob2621",
	}

	fmt.Println(admin)

	type person struct {
		login  string
		string // anonymous structure field (only 1 for 1 variables type else error)
		id     int
	}

	person1 := person{"Mike", "mikeLOL", 2}
	fmt.Println(person1.login, person1.string, person1.id)

	person2 := person1 // structure copy
	fmt.Println(person2)

	customer1 := customerType{"Rob", "Moon", 19}
	fmt.Println(customer1)
	incrementAge(&customer1)
	fmt.Println(customer1)

	type perInf struct {
		name    string
		surname string
	}

	type adminStruct struct {
		login        string
		password     string
		personalInfo perInf // struct as structer field 
	}

	var admin2 adminStruct = adminStruct{"bobLOL", "bob5t21376", perInf{"bob", "horse"}}

	fmt.Println(admin2)

}
