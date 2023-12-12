package main

import "fmt"

func main() {
	login()
}

func login() {
	var def_username = "azizozkrcdg"
	var def_password = 123456
	var username string
	var password int
	fmt.Print("username : ")
	fmt.Scan(&username)
	fmt.Print("Password : ")
	fmt.Scan(&password)

	if def_username == username && def_password == password {
		fmt.Print("Login Success")
	} else {
		fmt.Println("Login failed! Try again please.")
	}
}
