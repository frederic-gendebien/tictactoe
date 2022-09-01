package domain

import (
	"fmt"
)

func NewPlayer(symbol rune) Player {
	return Player{
		Symbol: symbol,
	}
}

type Player struct {
	Symbol rune
}

func (p Player) Play(ask func(Player) (Coordinates, error), grid *Grid) error {
	coordinates, err := ask(p)
	if err != nil {
		return fmt.Errorf("what did you do? %w", err)
	}

	return grid.Set(coordinates, p.Symbol)
}

func (p Player) Wins(grid *Grid) bool {
	return grid.Has3InARow(p.Symbol)
}
