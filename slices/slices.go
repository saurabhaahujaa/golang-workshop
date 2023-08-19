package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func printSlice(s string, x []int) {

	fmt.Printf("%s len=%d cap=%d %v\n",

		s, len(x), cap(x), x)

}

func main() {
	games := []string{"kokemon", "sims"}
	newGames := []string{"pacman", "doom", "pong"}

	newGames = games
	games = nil

	games = []string{}
	var ok string
	for i, game := range games {
		if game != newGames[i] {
			ok = "not "
			break
		}
	}

	// if games == nil {
	// 	ok = "not"
	// }

	if len(games) != len(newGames) {
		ok = "not"
	}

	fmt.Printf("games and newGames are %s equal \n\n", ok)
	fmt.Printf("games : %#v\n", games)
	fmt.Printf("new games :%#v\n", newGames)
	var todo []string

	todo = append(todo, "sing")

	todo = append(todo, "run")

	todo = append(todo, "code", "play")

	tomorrow := []string{"see mom", "learn go"}

	todo = append(todo, tomorrow...)

	s.Show("todo", todo)

	items := []string{

		"pacman", "mario", "tetris", "doom",

		"galaga", "frogger", "asteroids", "simcity",

		"metroid", "defender", "rayman", "tempest",

		"ultima",
	}

	s.MaxPerLine = 4

	s.Show("All items", items)

	top3 := items[:3]

	s.Show("Top 3 items", top3)

	l := len(items)

	last4 := items[l-4:]

	s.Show("Last 4 items", last4)

	mid := last4[1:3]

	s.Show("Last4[1:3]", mid)

	fmt.Printf("slicing : %T %[1]q\n", items[2:3])

	fmt.Printf("indexing: %T %[1]q\n", items[2])
	numbers := [5]int{1, 2, 3, 4, 5}

	//Both start and end

	num1 := numbers[2:4]

	fmt.Println("Both start and end")

	fmt.Printf("num1=%v\n", num1)

	fmt.Printf("length=%d\n", len(num1))

	fmt.Printf("capacity=%d\n", cap(num1))
	a := make([]int, 5)

	printSlice("a", a)
}
