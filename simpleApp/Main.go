package main

import "fmt"

type User struct {
	name      string
	password  string
	accountNo int
}

func login() {
	var password string
	var username string

	fmt.Print("Enter your name: ")
	fmt.Scanln(&username)
	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	if (username == "admin" && password == "admin") {
		fmt.Print("Login successfully")
	} else {
		fmt.Print("Login failed! Try again.")

	}

}
func main() {
	login()



}
