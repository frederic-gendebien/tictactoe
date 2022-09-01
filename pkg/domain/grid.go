package domain

import "fmt"

type Coordinates struct {
	Row int
	Col int
}

func NewGrid() *Grid {
	var matrix [3][3]rune
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = ' '
		}
	}

	return &Grid{
		Matrix:    matrix,
		available: 3 * 3,
	}
}

type Grid struct {
	Matrix    [3][3]rune
	available int
}

func (g *Grid) Set(coordinates Coordinates, symbol rune) error {
	if coordinates.Row < 0 || coordinates.Row > 3 || coordinates.Col < 0 || coordinates.Col > 3 {
		return fmt.Errorf("you entered wrong coordinates, dumbass")
	}

	if g.Matrix[coordinates.Row][coordinates.Col] != ' ' {
		return fmt.Errorf("you already played there, dumbass")
	}

	g.Matrix[coordinates.Row][coordinates.Col] = symbol
	g.available--

	return nil
}

func (g *Grid) IsFull() bool {
	return g.available == 0
}

func (g *Grid) Has3InARow(interesting rune) bool {
	verticals := []int{0, 0, 0}
	dDiagonals := 0
	aDiagonals := 0

	for row, line := range g.Matrix {
		horizontal := 0
		for col, symbol := range line {
			if symbol == interesting {
				horizontal++
				verticals[col] = verticals[col] + 1
				if row == col {
					dDiagonals++
				}
				if row+col == 2 {
					aDiagonals++
				}
			}
		}
		if horizontal == 3 {
			return true
		}
	}

	for _, sum := range verticals {
		if sum == 3 {
			return true
		}
	}

	if dDiagonals == 3 || aDiagonals == 3 {
		return true
	}

	return false
}
