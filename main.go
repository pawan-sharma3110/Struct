package main

import "fmt"

type Employee struct {
	Name string
	Age  int
}

func main() {
	var data []Employee
	var choice string
	e1 := Employee{
		Name: "Pawan",
		Age:  18,
	}
	e2 := Employee{
		Name: "Vikash",
		Age:  55,
	}
	fmt.Println("Do want Update Extis employes fields chose Y for Yes or N for No")
	fmt.Scanf("%s", &choice)
	if choice == "Y" || choice == "y" {

	}
	fmt.Println(data, e1, e2)
}
