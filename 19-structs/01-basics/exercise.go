package basics

import (
	"fmt"
	"strconv"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR STRUCTS HERE
type item struct {
	id, price int
	name      string
}

type game struct {
	item
	genre string
}

// newGame returns a new game struct.
func newGame(id int, name string, price int, genre string) game {
	// INSERT YOUR CODE HERE
	return game{item{id, price, name}, genre}
}

// String stringifies an item.
func (i item) String() string {
	// INSERT YOUR CODE HERE
	return strconv.Itoa(i.id) + ": " + i.name + " costs " + strconv.Itoa(i.price)
}

// String stringifies a game.
func (g game) String() string {
	// INSERT YOUR CODE HERE
	return "Game " + strconv.Itoa(g.item.id) + ": " + g.item.name + " costs " + strconv.Itoa(g.item.price) + " of genre " + g.genre
}

// newGameList creates a game store.
func newGameList() []game {
	// INSERT YOUR CODE HERE
	var gameList []game
	gameList = append(gameList,
		game{
			item: item{
				id:    1,
				price: 50,
				name:  "god of war",
			},
			genre: "action adventure",
		},
	)
	gameList = append(gameList,
		game{
			item: item{
				id:    3,
				price: 20,
				name:  "minecraft",
			},
			genre: "sandbox",
		},
	)
	gameList = append(gameList,
		game{
			item: item{
				id:    4,
				price: 40,
				name:  "warcraft",
			},
			genre: "strategy",
		},
	)
	return gameList
}

// queryById returns the game in the specified store with the given id or returns a "no such game" error.
func queryById(games []game, id int) (game, error) {
	// INSERT YOUR CODE HERE
	for i := range games {
		if games[i].item.id == id {
			return games[i], nil
		}
	}
	return game{}, fmt.Errorf("no such game")
}

// listNameByPrice returns the name of the game(s) with price equal or smaller than a given price.
func listNameByPrice(games []game, price int) []string {
	// INSERT YOUR CODE HERE
	var filteredGames []string
	for i := range games {
		if games[i].item.price <= price {
			filteredGames = append(filteredGames, games[i].item.name)
		}
	}
	return filteredGames
}
