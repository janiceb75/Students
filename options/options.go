package options

import (
	"fmt"
)

type Student struct { //first I will create a struct which tells what data I will collect from studen
	FirstName string
	LastName  string
	Age       int
}

func newStudent(FirstName, LastName string, Age int) Student {
	return Student{FirstName: FirstName, LastName: LastName, Age: Age}

}

func (s Student) FullName() string {
	return s.FirstName + s.LastName

}

// Choice func asks for choice number and presents code for each option
func Choice() {
	var choice int
	students := make(map[int]Student)
	for {
		fmt.Println("Enter your option number here: ")
		fmt.Scanln(&choice) //takes input from user(scans) and assigns to variable choice
		switch choice {     //after switch keyword, enter the value you wants to evaluate
		case 1:

			studentID := 1

			for {
				var studentFirstName, studentLastName, continueInput, continueProgram string
				var studentAge int

				fmt.Println("Please enter student's first name: ")
				fmt.Scanln(&studentFirstName)
				fmt.Println("Please enter student's last name: ")
				fmt.Scanln(&studentLastName)
				fmt.Println("Please enter student's age: ")
				fmt.Scanln(&studentAge)
				s := newStudent(studentFirstName, studentLastName, studentAge)
				fmt.Println(s.FullName())
				students[studentID] = s
				studentID += 1

				fmt.Println("Do you want to add another student? Enter 'yes' or 'no': ")
				fmt.Scanln(&continueInput)
				if continueInput != "yes" {
					fmt.Println(students)
					fmt.Println("Do you want to go back to the main menu? Enter 'yes' or 'no': ")
					fmt.Scanln(&continueProgram)
					// if continueProgram == "yes" {
					// 	Choice()

					// } else {
					// 	break
					// }
					break

				}

			}

		case 2:
			var studentID, age int

			fmt.Println("Please enter student's ID: ")
			fmt.Scanln(&studentID)
			fmt.Println("Please enter student's age: ")
			fmt.Scanln(&age)
			fmt.Println(students)
			student, ok := students[studentID]
			fmt.Println(ok)
			student.Age = age
			fmt.Printf("student: %v\n", student)
			//Choice()

		}

	}
}
