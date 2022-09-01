package interfaces

import (
	"tictactoe/pkg/domain"
)

type Interface interface {
	ShowGrid(grid *domain.Grid)
	AskCoordinatesTo(player domain.Player) (domain.Coordinates, error)
	Congratulate(player domain.Player)
	Slap(player domain.Player, err error)
	NoWinner()
}
