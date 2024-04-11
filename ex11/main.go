package main

import "fmt"

type Neighborhood struct {
	Name       string
	Households []Household
}
type Household struct {
	Owner       string
	MuchMembers int
	Address     any
}

var NeighborhoodDetails = []Neighborhood{
	{Name: "Subash Colony",
		Households: []Household{
			{Owner: "Pawan",
				MuchMembers: 8,
				Address:     "Kheri-Meham",
			},
		},
	},
}

func main() {
	fmt.Println(NeighborhoodDetails)
	// Household Struct within Neighborhood Struct:
	// Create a struct type Household representing a household with fields like Address, Owner, and Residents (a slice of Resident structs representing the people living in the household). Define another struct type Neighborhood with fields such as Name and an embedded slice of Household structs representing the households in the neighborhood. Write methods to add new households to the neighborhood, find the household with the most residents, and print the list of residents in a specific household
}
