package main

import "fmt"

type Tournament struct {
	Name string
	Team []Team
}
type Team struct {
	TotalPlayer int
	CoachName   string
	Players     []Player
}
type Player struct {
	PlayerName    string
	Age           int
	Gander        string
	ContactNumber string
}

var TournamentDetails = Tournament{
	Name: "Word tournament",
	Team: []Team{
		{TotalPlayer: 12,
			CoachName: "Arvind Kajriwal",
			Players: []Player{
				{PlayerName: "Pawan",
					Age:           19,
					ContactNumber: "7988323110",
					Gander:        "Male",
				},
				{PlayerName: "Khushi",
					Age:           19,
					ContactNumber: "7988254896",
					Gander:        "Female",
				},
				{PlayerName: "Kunal",
					Age:           19,
					ContactNumber: "7456241369",
					Gander:        "Male",
				},
				{PlayerName: "Kamal",
					Age:           19,
					ContactNumber: "9919205869",
					Gander:        "Male",
				},
				{PlayerName: "Rahul",
					Age:           19,
					ContactNumber: "78552233654",
					Gander:        "Male",
				},
				{PlayerName: "Vishal",
					Age:           19,
					ContactNumber: "74101245689",
					Gander:        "Male",
				},
				{PlayerName: "Arun",
					Age:           19,
					ContactNumber: "7845127589",
					Gander:        "Male",
				},
			},
		},
	},
}

func main() {
	fmt.Println(TournamentDetails)
}
