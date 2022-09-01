package usecase

import (
	"tictactoe/pkg/domain"
	"tictactoe/pkg/interfaces"
)

func NewGame(grid *domain.Grid, player1 domain.Player, player2 domain.Player) *Game {
	return &Game{
		grid:    grid,
		players: []domain.Player{player1, player2},
	}
}

type Game struct {
	grid               *domain.Grid
	players            []domain.Player
	currentPlayerIndex int
	currentPlayer      domain.Player
}

func (g *Game) Play(intrface interfaces.Interface) {
	for {
		currentPlayer := g.nextPlayer()
		correctPlay := false

		intrface.ShowGrid(g.grid)
		for !correctPlay {
			if err := currentPlayer.Play(intrface.AskCoordinatesTo, g.grid); err != nil {
				intrface.Slap(currentPlayer, err)
				intrface.ShowGrid(g.grid)
			} else {
				correctPlay = true
			}
		}
		if currentPlayer.Wins(g.grid) {
			intrface.ShowGrid(g.grid)
			intrface.Congratulate(currentPlayer)
			return
		}
		if g.grid.IsFull() {
			intrface.ShowGrid(g.grid)
			intrface.NoWinner()
			return
		}
	}
}

func (g *Game) nextPlayer() domain.Player {
	player := g.players[g.currentPlayerIndex]
	g.currentPlayerIndex = (g.currentPlayerIndex + 1) % 2

	return player
}

func (g *Game) finished(player domain.Player) bool {
	return g.grid.IsFull() || g.grid.Has3InARow(player.Symbol)
}
