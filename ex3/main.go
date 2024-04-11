package main

import "fmt"

type Student []struct {
	Name   string
	Age    int
	Class  int
	Grade  string
	Adress Address
}

type Address struct {
	HouseNo  int
	City     string
	District string
	State    string
}

var StudentDtails = []Student{
	{
		{
			Name:   "John Doe",
			Age:    18,
			Class:  10,
			Grade:  "A",
			Adress: Address{HouseNo: 123, City: "New York", District: "Manhattan", State: "NY"},
		},
		{
			Name:   "Jane Smith",
			Age:    19,
			Class:  11,
			Grade:  "B",
			Adress: Address{HouseNo: 456, City: "Los Angeles", District: "Downtown", State: "CA"},
		},
	},
}

func main() {
	for _, student := range StudentDtails {
		for _, detail := range student {
			fmt.Printf("Name: %s, Age: %d, Class: %d, Grade: %s, Address: %d %s, %s, %s\n",
				detail.Name, detail.Age, detail.Class, detail.Grade,
				detail.Adress.HouseNo, detail.Adress.City, detail.Adress.District, detail.Adress.State)
		}
	}
}
