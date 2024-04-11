package main

import "fmt"

type Order struct {
	OrderID        int
	TotalAmount    int
	IsOderDiliverd bool
	Item           []string
}

type Costumer struct {
	Name  string
	Email string
	Order Order
}

var CostumerOder = []Costumer{
	{Name: "Pawan",
		Email: "pawansharma93520@gmail.com",
		Order: Order{OrderID: 5620,
			TotalAmount:    520,
			IsOderDiliverd: true,
			Item: []string{
				"Smart Watch", "Phone", "Fan", "Table",
			},
		},
	},
}

func main() {
	fmt.Println(CostumerOder)
}
