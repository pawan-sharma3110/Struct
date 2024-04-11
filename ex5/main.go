package main

import "fmt"

type Address struct {
	HouseNo  int
	City     string
	District string
	State    string
	ZipCode  int
}
type Employee struct {
	Name   string
	ID     int
	Salary int
	Address
}

var Details = []Employee{
	{Name: "Pawan",
		ID:     1,
		Salary: 50000,
		Address: Address{
			HouseNo:  123,
			City:     "Rohtak",
			District: "Rohtak",
			State:    "Haryana",
			ZipCode:  124112,
		},
	}, {Name: "Rahul",
		ID:     1,
		Salary: 20000,
		Address: Address{
			HouseNo:  1553,
			City:     "Meham",
			District: "Rohtak",
			State:    "Haryana",
			ZipCode:  124112,
		},
	},
}

func main() {
	fmt.Println(Details)
}
