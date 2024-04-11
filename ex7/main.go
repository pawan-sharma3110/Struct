package main

import "fmt"

type Department struct {
	DepartmentName string
	ManagerName    string
	Employees      []Employee
}
type Employee struct {
	Name         string
	EmployeId    any
	IsManager    bool
	Age          int
	MobileNumber string
}

var DepartmentDetails = []Department{
	{
		DepartmentName: "Forest Department",
		ManagerName:    "Rahul Tiwari",
		Employees: []Employee{
			{Name: "Pawan",
				EmployeId:    1234,
				Age:          21,
				MobileNumber: "7988323110",
				IsManager:    false,
			},
			{Name: "Ashu",
				Age:          20,
				EmployeId:    21321,
				MobileNumber: "8295068010",
				IsManager:    false,
			},
			{Name: "Ashish",
				Age:          28,
				EmployeId:    859746,
				MobileNumber: "8952639515",
				IsManager:    true,
			},
		},
	},
}

func main() {
	fmt.Println(DepartmentDetails)
}
