package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	grade float32
	Person
}

func computeDetails(p Person) string {
	return ("Person's name is:" + p.name + "\n and the age:" )
}

func main() {
	var user1 Person
	var user2 Person
	user1.name = "Denisa Dan"
	user1.age = 23
	user2.name = "John Doe"

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(Person{})

	new := Student{
		grade: 10,
		Person: Person{
			name: "Denisa",
			age: 50,
		},
	}
	fmt.Println(new)

	fmt.Println(Student{5, Person{name: "Unknown", age: 33}})
	pers := Person{"Young Bob", 33}
	fmt.Println(computeDetails(pers))
}
