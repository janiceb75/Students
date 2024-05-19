package main

import (
	"fmt"

	"github.com/janiceb75/students/options"
)

func main() {

	teacherMenu()
	options.Choice()

}

// Teacher menus asks for teacher name, greets teacher and presents menu choices.  Will likely never
// need to use this anywhere else, which is why I put in main package
func teacherMenu() {
	var name string
	fmt.Println("What is your name? Don't enter any spaces: ")
	fmt.Scanln(&name)
	fmt.Println("Welcome to Brown Elementary,", name, " Please choose an option below to get started:")
	fmt.Println("Option 1 - Register your students")
	fmt.Println("Option 2 - Submit final grades")
	fmt.Println("Option 3 - Order classroom supplies")

}
