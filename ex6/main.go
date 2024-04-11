package main

import "fmt"

type Course struct {
	CourseName string
	Instructor string
	Credits    int
}

type Student struct {
	StudentName string
	RollNo      int
	Age         int
	CourseName  []Course
}

var studentDetails = []Student{
	{
		StudentName: "Pawan",
		RollNo:      5,
		Age:         21,
		CourseName: []Course{
			{CourseName: "Hindi", Instructor: "Ashu", Credits: 3},
			{CourseName: "English", Instructor: "Ram", Credits: 4},
			{CourseName: "S.S", Instructor: "Raman", Credits: 3},
		},
	},
}

func main() {
	fmt.Println(studentDetails)
}
