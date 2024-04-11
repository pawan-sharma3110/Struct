package main

import "fmt"

type District struct {
	DistrictName string
	Schools      []School
}
type School struct {
	SchoolName    string
	PrincipalName string
	Address       any
	Students      []Student
}
type Student struct {
	StudentName string
	Class       int
	Age         int
	Subject     string
}

var DistrictDetails = []District{
	{DistrictName: "Rohtak",
		Schools: []School{
			{SchoolName: "Govt. Sr. Sec School",
				PrincipalName: "Bijander",
				Address:       "Kheri-Meham",
				Students: []Student{
					{StudentName: "Pawan",
						Class:   12,
						Age:     19,
						Subject: "Simple Arts",
					},
				},
			},
		},
	},
}

func main() {

	fmt.Println(DistrictDetails)
	// School Struct within District Struct:
	// Create a struct type School representing a school with fields like Name, Principal, and Students (a slice of Student structs representing the students in the school). Define another struct type District with fields such as Name and an embedded slice of School structs representing the schools in the district. Write methods to add new schools to the district, count the total number of students in the district, and print the list of schools along with their principals.
}
