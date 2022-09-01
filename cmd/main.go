package main

import (
	"tictactoe/pkg/domain"
	"tictactoe/pkg/interfaces"
	"tictactoe/pkg/interfaces/console"
	"tictactoe/pkg/usecase"
)

var (
	grid     *domain.Grid
	player1  domain.Player
	player2  domain.Player
	game     *usecase.Game
	intrface interfaces.Interface
)

func init() {
	grid = domain.NewGrid()
	player1 = domain.NewPlayer('x')
	player2 = domain.NewPlayer('o')
	game = usecase.NewGame(grid, player1, player2)
	intrface = console.NewConsole()
}

func main() {
	game.Play(intrface)
}
