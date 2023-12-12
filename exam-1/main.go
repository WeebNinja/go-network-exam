package main

import "fmt"

func main() {
	// Got Username and Password from keyboard
	var username string
	var password string

	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)
	fmt.Println("Add = ", Add(2, 3))
	fmt.Println("Minus = ", Minus(3, 2))
	// Call function LogIn

	LogIn(username, password)
}

func HelloFunction() string {
	return "Hello"
}

func Add(a int, b int) int {
	return a + b
}

func Minus(a int, b int) int {
	return a - b
}

func UserLogin(username string, password string) bool {
	if username == "123" && password == "123" {
		return true
	} else {
		return false
	}
}

// Use function UserLogin
func LogIn(username string, password string) {
	if UserLogin(username, password) {
		fmt.Println("Login successful")
	} else {
		fmt.Println("Login failed")
	}
}
