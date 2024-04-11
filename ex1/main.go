package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type AllCompanyName struct {
	CompanyNames []string
	Employye     Employye
}
type Employye struct {
	ID       int
	Name     string
	Age      int
	Position string
	Salary   int
}

var AllCompanyeDtails = []AllCompanyName{{
	CompanyNames: []string{
		"Nokia", "Lava", "Samsung", "Redmi",
	},
	Employye: Employye{
		ID:       1,
		Name:     "Pawan",
		Age:      21,
		Position: "Manager",
		Salary:   55000,
	},
}}

func main() {
	data, err := json.MarshalIndent(AllCompanyeDtails, " ", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data))
}
